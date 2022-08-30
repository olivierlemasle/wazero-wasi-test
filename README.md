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
   GO_ENABLED=0 GOOS=linux go build .
   ```

1. Build the Docker image:

   ```
   docker build --platform linux/amd64 . -t wazero-wasi-test
   ```

2. Run the Docker image

   ```
   docker run --rm wazero-wasi-test
   ```
