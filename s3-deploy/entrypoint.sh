#!/bin/bash

set -Eeuo pipefail

if [ -n "${GITHUB_WORKSPACE-}" ]; then
	cd $GITHUB_WORKSPACE
fi

if [ -n "${INPUT_BASE_DIR-}" ]; then
	cd $INPUT_BASE_DIR
fi

aws s3 sync --delete . s3://$INPUT_BUCKET_NAME | {
	# List of files that are being uploaded or deleted
	FILES=()

	# Extract filename from s3 sync output
	while read -r line;
 	do
 		PARSED=$(echo $line |\
 			# Get full operation path
			sed -E "s/.*(upload|delete): (.*)/\2/" |\
			# Remove current directory slash
			sed -E "s/^.\///" |\
			# Get destination filename
			sed -E "s/.*s3:\/\/$INPUT_BUCKET_NAME\/(.*)/\1/")
		FILES+=("/$PARSED")
	done

	for file in "${FILES[@]}"; do
		# Add an invalidation path for index routes without "index.html"
		# e.g. invalidate both /about/index.html and /about/
		if [[ "$file" =~ \/index\.html$ ]]; then
			PARSED=$(echo $file | sed -E "s/index.html//")
			FILES+=("$PARSED")
		fi
	done

	# Run invalidation if any files were uploaded or deleted
	if [ -n "${INPUT_CLOUDFRONT_ID-}" ] && [ ${#FILES[@]} -gt 0 ]; then
		# Create string from filelist that can be passed as argument
		if [ "${INPUT_INVALIDATE_WILDCARD-}" != "true" ]; then
			PATHS=$(printf "%q " "${FILES[@]}")
		fi

		aws cloudfront create-invalidation \
			--distribution-id $INPUT_CLOUDFRONT_ID \
			--paths ${PATHS:-"/*"} \
			--output text
	fi
}

