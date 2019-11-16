#!/bin/bash

set -Eeuo pipefail

if [ -n "${GITHUB_WORKSPACE-}" ]; then
	cd $GITHUB_WORKSPACE
fi

if [ -n "${INPUT_BASE_DIR-}" ]; then
	cd $INPUT_BASE_DIR
fi

if [ -n "${INPUT_PREBUILD-}" ]; then
	$INPUT_PREBUILD
fi

zola check \
	--config $INPUT_CONFIG_FILE

zola build \
	--config $INPUT_CONFIG_FILE \
	--output-dir $INPUT_OUTPUT_DIR

if [ -n "${INPUT_POSTBUILD-}" ]; then
	$INPUT_POSTBUILD
fi
