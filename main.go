package main

import (
	"fmt"
	"log"
	"os"

	"github.com/bytecodealliance/wasmtime-go/v18"
	"github.com/elewis787/wasmtime-go-nn/ml"
	llama "github.com/go-skynet/go-llama.cpp"
)

func main() {
	config := wasmtime.NewConfig()
	config.SetWasmSIMD(true)
	engine := wasmtime.NewEngineWithConfig(config)

	//read the wasmfile
	wasm, err := os.ReadFile("/Users/elewis/Projects/go/src/github.com/elewis787/wasmtime-go-nn/example/go/main.wasm")
	if err != nil {
		log.Fatal(err)
	}
	module, err := wasmtime.NewModule(engine, wasm)
	if err != nil {
		log.Fatal(err)
	}

	// Create a linker with WASI functions defined within it
	linker := wasmtime.NewLinker(engine)
	err = linker.DefineWasi()
	if err != nil {
		log.Fatal(err)
	}
	linker.FuncNew("wasi:nn/graph", "load-by-name", wasmtime.NewFuncType([]*wasmtime.ValType{
		wasmtime.NewValType(wasmtime.KindI32),
		wasmtime.NewValType(wasmtime.KindI32),
		wasmtime.NewValType(wasmtime.KindI32)},
		[]*wasmtime.ValType{}),
		func(caller *wasmtime.Caller, args []wasmtime.Val) ([]wasmtime.Val, *wasmtime.Trap) {
			mem := caller.GetExport("memory").Memory()
			ptr := args[0].I32()
			len := args[1].I32()
			data := mem.UnsafeData(caller)[ptr : ptr+len]
			name := string(data)
			var (
				//threads   = 4
				//tokens    = 128
				gpulayers = 0
				//seed      = -1
			)
			// This is where LLAMA CPP lives
			_, err := llama.New(name, llama.EnableF16Memory, llama.SetContext(128), llama.EnableEmbeddings, llama.SetGPULayers(gpulayers))
			if err != nil {
				fmt.Println("Loading the model failed:", err.Error())
				os.Exit(1)
			}

			res := &ml.Result[ml.WasiNnGraphGraph, ml.WasiNnGraphError]{}
			res.Set(0)
			fmt.Println(res)
			return []wasmtime.Val{}, nil
		},
	)

	// Configure WASI imports to write stdout into a file, and then create
	// a `Store` using this wasi configuration.
	wasiConfig := wasmtime.NewWasiConfig()
	wasiConfig.InheritArgv()
	wasiConfig.InheritStdout()
	store := wasmtime.NewStore(engine)
	store.SetWasi(wasiConfig)

	instance, err := linker.Instantiate(store, module)
	if err != nil {
		log.Fatal(err)
	}

	// Run the function
	nom := instance.GetFunc(store, "_start")
	_, err = nom.Call(store)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("done")
}
