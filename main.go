package main


import (
   "fmt"
   "log"
   "os"

   "github.com/bytecodealliance/wasmtime-go/v18"
)


func load_by_name(a, b, c int32) int32 {
   fmt.Println("load_by_name")
   return 0
}


func main() {
   engine := wasmtime.NewEngine()

   //read the wasmfile
   wasm, err := os.ReadFile("wasm")
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
       func(*wasmtime.Caller, []wasmtime.Val) ([]wasmtime.Val, *wasmtime.Trap) {
           fmt.Println("load-1")
           res := load_by_name(0, 0, 0)
           return []wasmtime.Val{wasmtime.ValI32(res)}, nil
       },
   )

   linker.FuncNew("wasi_ephemeral_nn", "load_by_name_with_config", wasmtime.NewFuncType([]*wasmtime.ValType{
       wasmtime.NewValType(wasmtime.KindI32),
       wasmtime.NewValType(wasmtime.KindI32),
       wasmtime.NewValType(wasmtime.KindI32),
       wasmtime.NewValType(wasmtime.KindI32),
       wasmtime.NewValType(wasmtime.KindI32)},
       []*wasmtime.ValType{wasmtime.NewValType(wasmtime.KindI32)}),
       func(*wasmtime.Caller, []wasmtime.Val) ([]wasmtime.Val, *wasmtime.Trap) {
           fmt.Println("load-2")
           res := load_by_name(0, 0, 0)
           return []wasmtime.Val{wasmtime.ValI32(res)}, nil
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






