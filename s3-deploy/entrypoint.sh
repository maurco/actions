#!/bin/bash

set -e

DEPLOY_DIR=$1
BUCKET_NAME=$2
CLOUDFRONT_ID=$3
INVALIDATE_ALL=$4

FILES=()

if [ -z $BUCKET_NAME ]; then
	echo "Missing argument: \`bucket_name\`"
	exit 1;
fi

if [ -n $GITHUB_WORKSPACE ]; then
	cd $GITHUB_WORKSPACE
fi

cd $DEPLOY_DIR

aws s3 sync --delete . s3://$BUCKET_NAME | {
	while read -r i;
 	do
		FILES+=("/$(echo $i |\
 			sed -E "s/.*upload: (.*) to.*/\1/" |\
 			sed -E "s/^.\///" |\
 			sed -E "s/\/index.html/\//")")
	done

	if [ -n $CLOUDFRONT_ID ] && [ ${#FILES[@]} -gt 0 ]; then
		IFS=" "; shift

		if [ "$INVALIDATE_ALL" = "true" ]; then
			PATHS="'/*'"
		else
			PATHS=$(echo "${FILES[@]}")
		fi

		aws cloudfront create-invalidation \
			--distribution-id $CLOUDFRONT_ID \
			--paths $PATHS \
			--output text
	fi
}

