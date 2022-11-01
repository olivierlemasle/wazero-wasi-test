package main

import (
	"bytes"
	"context"
	_ "embed"
	"fmt"
	"log"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/imports/wasi_snapshot_preview1"
)

//go:embed hello-world-rs.wasm
var helloWorldRs []byte

//go:embed hello-world-as.wasm
var helloWorldAs []byte

func main() {
	ctx := context.Background()

	compileAndRun(ctx, "hello-world-rust", helloWorldRs)
	compileAndRun(ctx, "hello-world-assemblyscript", helloWorldAs)
}

func compileAndRun(ctx context.Context, name string, wasm []byte) {
	runtime := wazero.NewRuntime(ctx)
	defer runtime.Close(ctx)

	if _, err := wasi_snapshot_preview1.Instantiate(ctx, runtime); err != nil {
		log.Panicln(err)
	}

	fmt.Printf("Compilation of %s... ", name)
	compiled, err := runtime.CompileModule(ctx, wasm)
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println("[ok]")

	fmt.Println("Running...")
	var out bytes.Buffer
	c := wazero.NewModuleConfig().
		WithStdout(&out).
		WithName(name)
	if _, err = runtime.InstantiateModule(ctx, compiled, c); err != nil {
		log.Panicln(err)
	}

	fmt.Println(out.String())
	fmt.Println()
}
