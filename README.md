# wazero-wasi-test

The objective of this repository is to reproduce an issue with
[wazero](https://wazero.io/).

Steps to reproduce the issue:

1. Download the WebAssembly WASI binaries:

   ```
   ./download-wasm.sh
   ```

1. Build the Go binary:

   ```
   GO_ENABLED=0 go build .
   ```

1. Build the Docker image:

   ```
   docker build . -t wazero-wasi-test
   ```
