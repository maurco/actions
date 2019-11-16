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

zola --config $INPUT_CONFIG_FILE check
zola --config $INPUT_CONFIG_FILE build --output-dir $INPUT_OUTPUT_DIR

if [ -n "${INPUT_POSTBUILD-}" ]; then
	$INPUT_POSTBUILD
fi
