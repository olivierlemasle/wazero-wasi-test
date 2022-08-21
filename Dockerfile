FROM gcr.io/distroless/static:nonroot

COPY ./wazero-wasi-test /
CMD ["/wazero-wasi-test"]
