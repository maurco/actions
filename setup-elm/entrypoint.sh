#!/bin/bash

set -Eeuo pipefail

wget -qO elm.gz https://github.com/elm/compiler/releases/download/${INPUT_ELM_VERSION}/binary-for-linux-64-bit.gz
gunzip elm.gz
chmod +x elm
mv elm /usr/local/bin

echo "Elm ${INPUT_ELM_VERSION} installed!"
