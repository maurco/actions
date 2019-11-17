#!/bin/bash

set -Eeuo pipefail

BIN_FILE=/usr/local/bin/elm

wget -qO- https://github.com/elm/compiler/releases/download/${INPUT_ELM_VERSION}/binary-for-linux-64-bit.gz |\
	gunzip -c > $BIN_FILE

chmod +x $BIN_FILE
