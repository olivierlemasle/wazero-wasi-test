#!/bin/bash

VERSION=0.0.1
for file in hello-world-as hello-world-rs
do
	curl -sL -o ${file}.wasm https://github.com/olivierlemasle/wasi-examples/releases/download/${VERSION}/${file}.wasm
	echo "Downloaded ${file}.wasm"
done
