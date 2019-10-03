#!/bin/bash

set -e

BUCKET_NAME=$1
CLOUDFRONT_ID=$2

FILES=()

if [ -z $BUCKET_NAME ]; then
	echo "Missing argument: \`bucket_name\`"
	exit 1;
fi

if [ -n $GITHUB_WORKSPACE ]; then
	cd $GITHUB_WORKSPACE
fi

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

		aws cloudfront create-invalidation \
			--distribution-id $CLOUDFRONT_ID \
			--paths $(echo "${FILES[@]}") \
			--output text
	fi
}

