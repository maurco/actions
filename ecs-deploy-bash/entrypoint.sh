#!/bin/bash

set -Eeuo pipefail

if [ -n "${GITHUB_WORKSPACE-}" ]; then
	cd $GITHUB_WORKSPACE
fi

if [ -n "${INPUT_BASE_DIR-}" ]; then
	cd $INPUT_BASE_DIR
fi

echo "=> Starting deployment"
export AWS_DEFAULT_REGION=$INPUT_AWS_REGION

JQ="jq -Merc"
COMMIT=$(git rev-parse --short HEAD~0)
LOCAL_IMAGE="$INPUT_APP_NAME:$COMMIT"

AWS_ACCOUNT_ID=$(aws sts get-caller-identity | $JQ ".Account")
REMOTE_IMAGE=$AWS_ACCOUNT_ID.dkr.ecr.$AWS_DEFAULT_REGION.amazonaws.com/$INPUT_APP_NAME
BUILD_ARGS=$(echo "${INPUT_DOCKER_BUILD_ARGS:-[]}" | $JQ ".[] |= \"--build-arg \" + . | join(\" \")")

echo "=> Logging into ECR"
$(aws ecr get-login --no-include-email)

echo "=> Building $LOCAL_IMAGE"
docker build $BUILD_ARGS -t $LOCAL_IMAGE .

echo "=> Pushing $LOCAL_IMAGE"
docker tag $LOCAL_IMAGE $REMOTE_IMAGE:latest
docker push $REMOTE_IMAGE:latest
docker rmi $REMOTE_IMAGE:latest

echo "=> Pushing $LOCAL_IMAGE"
docker tag $LOCAL_IMAGE $REMOTE_IMAGE:$COMMIT
docker push $REMOTE_IMAGE:$COMMIT
docker rmi $REMOTE_IMAGE:$COMMIT

echo "=> Fetching existing task definition"
TASK_DEF_OLD=$(
	aws ecs describe-task-definition --task-definition $INPUT_APP_NAME
)

echo "=> Registering new task definition"
TASK_DEF_PENDING=$(
	echo "$TASK_DEF_OLD" |\
		# $JQ ".taskDefinition.containerDefinitions[0].image = \"$REMOTE_IMAGE:$COMMIT\"" |\
		$JQ ".taskDefinition.containerDefinitions | map(.image = \"$REMOTE_IMAGE:$COMMIT\")" |\
		$JQ ".taskDefinition | del(.taskDefinitionArn) | del(.revision) | del(.status) | del(.requiresAttributes) | del(.compatibilities)"
)

TASK_DEF_NEW=$(
	aws ecs register-task-definition \
		--family $INPUT_APP_NAME \
		--cli-input-json $TASK_DEF_PENDING
)

# echo "=> Deregistering old task definition"
# aws ecs deregister-task-definition \
# 	--task-definition $(echo "$TASK_DEF_OLD" | $JQ ".taskDefinition.taskDefinitionArn") \
# 	> /dev/null

echo "=> Updating $INPUT_APP_NAME with new task definition"
aws ecs update-service \
	--service $INPUT_APP_NAME \
	--cluster $INPUT_CLUSTER_NAME \
	--force-new-deployment \
	--task-definition $(echo "$TASK_DEF_NEW" | $JQ ".taskDefinition.taskDefinitionArn") \
	> /dev/null

# if [ "${INPUT_S3_DEPLOYMENT_FILE-}" ]; then
# 	echo "=> Updating S3 deployment file with new version"
# 	echo $(aws s3 cp s3://$INPUT_S3_DEPLOYMENT_FILE -) |\
# 		$JQ ".image_tag = \"$COMMIT\"" |\
# 		aws s3 cp - s3://$INPUT_S3_DEPLOYMENT_FILE \
# 			--content-type application/json \
# 			--metadata-directive COPY
# fi

if [ "${INPUT_WAIT_FOR_DEPLOYMENT-}" = true ]; then
	echo "=> Waiting for deployment to roll out"
	aws ecs wait services-stable \
	    --cluster $INPUT_CLUSTER_NAME \
	    --services $INPUT_APP_NAME
fi

echo "=> Deployed"
