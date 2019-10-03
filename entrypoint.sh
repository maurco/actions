#!/bin/bash

set -e

BUCKET_NAME=$1
CLOUDFRONT_ID=$2

FILES=()

if [ -z $BUCKET_NAME ] || [ -z $CLOUDFRONT_ID ]; then
	echo "Missing arguments: \`entrypoint.sh BUCKET_NAME CLOUDFRONT_ID\`"
	exit 1;
fi

aws s3 sync --delete . s3://$BUCKET_NAME | {
	while read -r i;
 	do
		FILES+=("/$(echo $i |\
 			sed -E "s/.*upload: (.*) to.*/\1/" |\
 			sed -E "s/^.\///" |\
 			sed -E "s/\/index.html/\//")")
	done

	if [ ${#FILES[@]} -gt 0 ]; then
		IFS=" "; shift

		aws cloudfront create-invalidation \
			--distribution-id $CLOUDFRONT_ID \
			--paths $(echo "${FILES[@]}") \
			--output text
	else
		echo "Nothing was uploaded to S3."
	fi
}

