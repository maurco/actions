#!/bin/bash

set -Eeuo pipefail

BIN_FILE=/bin/elm

wget -qO- https://github.com/elm/compiler/releases/download/${INPUT_ELM_VERSION}/binary-for-linux-64-bit.gz |\
	gunzip -c > $BIN_FILE

chmod +x $BIN_FILE

echo "::add-path::$BIN_FILE"
echo "Elm ${INPUT_ELM_VERSION} installed!"
