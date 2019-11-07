#!/bin/bash

set -Eeuo pipefail

BUILD_COMMAND=$1
BASE_DIR=$2

if [ -n "${GITHUB_WORKSPACE-}" ]; then
	cd $GITHUB_WORKSPACE
fi

if [ -n "${BASE_DIR-}" ]; then
	cd $BASE_DIR
fi

$BUILD_COMMAND
