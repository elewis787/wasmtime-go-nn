package main

import (
	"fmt"
	"log"
	"os"
	"unsafe"

	"github.com/bytecodealliance/wasmtime-go/v18"
	"github.com/elewis787/wasmtime-go-nn/ml"
)

func load_by_name(a, b, c int32) int32 {
	fmt.Println("load_by_name")
	return 0
}

func main() {
	engine := wasmtime.NewEngine()

	//read the wasmfile
	wasm, err := os.ReadFile("/Users/elewis/Projects/go/src/github.com/elewis787/wasmtime-go-nn/example/rust/wasmedge-ggml-chatml.wasm")
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
	linker.FuncNew("wasi_ephemeral_nn", "load_by_name", wasmtime.NewFuncType([]*wasmtime.ValType{
		wasmtime.NewValType(wasmtime.KindI32),
		wasmtime.NewValType(wasmtime.KindI32),
		wasmtime.NewValType(wasmtime.KindI32)},
		[]*wasmtime.ValType{wasmtime.NewValType(wasmtime.KindI32)}),
		func(caller *wasmtime.Caller, args []wasmtime.Val) ([]wasmtime.Val, *wasmtime.Trap) {
			mem := caller.GetExport("memory").Memory()
			ptr := args[0].I32()
			len := args[1].I32()
			data := mem.UnsafeData(caller)[ptr : ptr+len]
			name := string(data)
			fmt.Println(name)

			// This is where LLAMA CPP lives

			res := &ml.Result[ml.WasiNnGraphGraph, ml.WasiNnGraphError]{}
			return []wasmtime.Val{wasmtime.ValI32(int32(uintptr(unsafe.Pointer(res))))}, nil
		},
	)
	linker.FuncWrap("wasi_ephemeral_nn", "load_by_name", func(name string) ml.Result[ml.WasiNnGraphGraph, ml.WasiNnGraphError] {

		return ml.Result[ml.WasiNnGraphGraph, ml.WasiNnGraphError]{}
	})

	linker.FuncNew("wasi_ephemeral_nn", "load_by_name_with_config", wasmtime.NewFuncType([]*wasmtime.ValType{
		wasmtime.NewValType(wasmtime.KindI32),
		wasmtime.NewValType(wasmtime.KindI32),
		wasmtime.NewValType(wasmtime.KindI32),
		wasmtime.NewValType(wasmtime.KindI32),
		wasmtime.NewValType(wasmtime.KindI32)},
		[]*wasmtime.ValType{wasmtime.NewValType(wasmtime.KindI32)}),
		func(caller *wasmtime.Caller, args []wasmtime.Val) ([]wasmtime.Val, *wasmtime.Trap) {
			mem := caller.GetExport("memory").Memory()
			ptr := args[0].I32()
			len := args[1].I32()
			data := mem.UnsafeData(caller)[ptr : ptr+len]
			name := string(data)

			fmt.Println(name)

			// This is where LLAMA CPP lives
			// if err != nil {
			//	return ml.not-found
			//}else {
			// return
			//}

			res := &ml.Result[ml.WasiNnGraphGraph, ml.WasiNnGraphError]{}
			res.Set(0)
			// Assume that args[2] is a pointer to a buffer in the Wasm memory
			bufPtr := args[2].I32()
			buf := mem.UnsafeData(caller)[bufPtr : bufPtr+int32(unsafe.Sizeof(*res))]
			copy(buf, (*[unsafe.Sizeof(*res)]byte)(unsafe.Pointer(res))[:])

			return []wasmtime.Val{wasmtime.ValI32(0)}, nil
		},
	)

	linker.FuncNew("wasi_ephemeral_nn", "init_execution_context", wasmtime.NewFuncType([]*wasmtime.ValType{
		wasmtime.NewValType(wasmtime.KindI32),
		wasmtime.NewValType(wasmtime.KindI32)},
		[]*wasmtime.ValType{wasmtime.NewValType(wasmtime.KindI32)}),
		func(*wasmtime.Caller, []wasmtime.Val) ([]wasmtime.Val, *wasmtime.Trap) {
			fmt.Println("load-3")
			res := load_by_name(0, 0, 0)
			return []wasmtime.Val{wasmtime.ValI32(res)}, nil
		},
	)

	linker.FuncNew("wasi_ephemeral_nn", "set_input", wasmtime.NewFuncType([]*wasmtime.ValType{
		wasmtime.NewValType(wasmtime.KindI32),
		wasmtime.NewValType(wasmtime.KindI32),
		wasmtime.NewValType(wasmtime.KindI32)},
		[]*wasmtime.ValType{wasmtime.NewValType(wasmtime.KindI32)}),
		func(*wasmtime.Caller, []wasmtime.Val) ([]wasmtime.Val, *wasmtime.Trap) {
			fmt.Println("load-4")
			res := load_by_name(0, 0, 0)
			return []wasmtime.Val{wasmtime.ValI32(res)}, nil
		},
	)

	linker.FuncNew("wasi_ephemeral_nn", "compute", wasmtime.NewFuncType([]*wasmtime.ValType{
		wasmtime.NewValType(wasmtime.KindI32)},
		[]*wasmtime.ValType{wasmtime.NewValType(wasmtime.KindI32)}),
		func(*wasmtime.Caller, []wasmtime.Val) ([]wasmtime.Val, *wasmtime.Trap) {
			fmt.Println("load-5")
			res := load_by_name(0, 0, 0)
			return []wasmtime.Val{wasmtime.ValI32(res)}, nil
		},
	)

	linker.FuncNew("wasi_ephemeral_nn", "get_output", wasmtime.NewFuncType([]*wasmtime.ValType{
		wasmtime.NewValType(wasmtime.KindI32),
		wasmtime.NewValType(wasmtime.KindI32),
		wasmtime.NewValType(wasmtime.KindI32),
		wasmtime.NewValType(wasmtime.KindI32),
		wasmtime.NewValType(wasmtime.KindI32)},
		[]*wasmtime.ValType{wasmtime.NewValType(wasmtime.KindI32)}),
		func(*wasmtime.Caller, []wasmtime.Val) ([]wasmtime.Val, *wasmtime.Trap) {
			fmt.Println("load-6")
			res := load_by_name(0, 0, 0)
			return []wasmtime.Val{wasmtime.ValI32(res)}, nil
		},
	)

	// Configure WASI imports to write stdout into a file, and then create
	// a `Store` using this wasi configuration.
	wasiConfig := wasmtime.NewWasiConfig()
	wasiConfig.InheritArgv()
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
