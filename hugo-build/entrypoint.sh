#!/bin/bash

set -Eeuo pipefail

if [ -n "${GITHUB_WORKSPACE-}" ]; then
	cd $GITHUB_WORKSPACE
fi

if [ -n "${INPUT_BASE_DIR-}" ]; then
	cd $INPUT_BASE_DIR
fi

if [ -n "${INPUT_COMMAND-}" ]; then
	$INPUT_COMMAND
else
	hugo \
		--gc \
		--minify \
		--path-warnings \
		--cleanDestinationDir \
		--environment $INPUT_ENVIRONMENT
fi
