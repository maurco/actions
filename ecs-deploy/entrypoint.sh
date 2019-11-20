#!/bin/bash

set -Eeuo pipefail

if [ -n "${GITHUB_WORKSPACE-}" ]; then
	cd $GITHUB_WORKSPACE
fi

if [ -n "${INPUT_BASE_DIR-}" ]; then
	cd $INPUT_BASE_DIR
fi

if [ -n "${INPUT_AWS_REGION-}" ]; then
	export AWS_DEFAULT_REGION=$INPUT_AWS_REGION
fi

echo "=> Starting deployment"

JQ="jq -Merc"
COMMIT=$(git rev-parse --short HEAD~0)
LOCAL_IMAGE="ecs-deploy-$COMMIT"

AWS_ACCOUNT_ID=$(aws sts get-caller-identity | $JQ ".Account")
REMOTE_IMAGE=$AWS_ACCOUNT_ID.dkr.ecr.$AWS_DEFAULT_REGION.amazonaws.com/$INPUT_APP_NAME

echo "=> Logging into ECR"
$(aws ecr get-login --no-include-email)

echo "=> Building $LOCAL_IMAGE:latest"
docker build ${BUILD_ARGS-} -t $LOCAL_IMAGE:latest . \
	$(echo ${INPUT_BUILD_ARGS:-[]} | $JQ ".[] |= \"--build-arg \" + . | join(\" \")")

echo "=> Pushing $LOCAL_IMAGE:latest"
docker tag $LOCAL_IMAGE:latest $REMOTE_IMAGE:latest
docker push $REMOTE_IMAGE:latest
docker rmi $REMOTE_IMAGE:latest

echo "=> Pushing $LOCAL_IMAGE:$COMMIT"
docker tag $LOCAL_IMAGE:latest $REMOTE_IMAGE:$COMMIT
docker push $REMOTE_IMAGE:$COMMIT
docker rmi $REMOTE_IMAGE:$COMMIT

echo "=> Fetching existing task definition"
TASK_DEF_OLD=$(
	aws ecs describe-task-definition --task-definition $INPUT_APP_NAME
)

echo "=> Registering new task definition"
TASK_DEF_PENDING=$(
	echo "$TASK_DEF_OLD" |\
		$JQ ".taskDefinition.containerDefinitions[0].image = \"$REMOTE_IMAGE:$COMMIT\"" |\
		$JQ ".taskDefinition | del(.taskDefinitionArn) | del(.revision) | del(.status) | del(.requiresAttributes) | del(.compatibilities)"
)

TASK_DEF_NEW=$(
	aws ecs register-task-definition \
		--family $INPUT_APP_NAME \
		--cli-input-json $TASK_DEF_PENDING
)

echo "=> Deregistering old task definition"
aws ecs deregister-task-definition \
	--task-definition $(echo "$TASK_DEF_OLD" | $JQ ".taskDefinition.taskDefinitionArn") \
	> /dev/null

echo "=> Cleaning up untagged images"
UNTAGGED_IMAGES=$(
	aws ecr describe-images \
		--repository-name $INPUT_APP_NAME \
		--filter tagStatus=UNTAGGED |\
		$JQ "[.imageDetails[] | \"imageDigest=\" + .imageDigest]"
)

if [ $(echo "$UNTAGGED_IMAGES" | $JQ length) -gt 0 ]; then
	aws ecr batch-delete-image \
		--repository-name $INPUT_APP_NAME \
		--image-ids $(echo "$UNTAGGED_IMAGES" | $JQ "join(\",\")") \
		> /dev/null
fi

echo "=> Updating $INPUT_APP_NAME with new task definition"
aws ecs update-service \
	--service $INPUT_APP_NAME \
	--cluster $INPUT_CLUSTER_NAME \
	--force-new-deployment \
	--task-definition $(echo "$TASK_DEF_NEW" | $JQ ".taskDefinition.taskDefinitionArn") \
	> /dev/null

echo "=> Waiting for deployment to roll out"
aws ecs wait services-stable \
    --cluster $INPUT_CLUSTER_NAME \
    --services $INPUT_APP_NAME

echo "=> Deployed"
