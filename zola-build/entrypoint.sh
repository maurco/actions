#!/bin/bash

set -Eeuo pipefail

if [ -n "${GITHUB_WORKSPACE-}" ]; then
	cd $GITHUB_WORKSPACE
fi

if [ -n "${INPUT_BASE_DIR-}" ]; then
	cd $INPUT_BASE_DIR
fi

ARGS_PRE=()
ARGS_POST=()

if [ -n "${INPUT_CONFIG_FILE-}" ]; then
	ARGS_PRE+=(--config "$INPUT_CONFIG_FILE")
fi

if [ -n "${INPUT_BASE_URL-}" ]; then
	ARGS_POST+=(--base-url "$INPUT_BASE_URL")
fi

if [ -n "${INPUT_OUTPUT_DIR-}" ]; then
	ARGS_POST+=(--output-dir "$INPUT_OUTPUT_DIR")
fi

zola "${ARGS_PRE[@]}" check
zola "${ARGS_PRE[@]}" build "${ARGS_POST[@]}"

chmod 777 ${INPUT_OUTPUT_DIR:-public}
