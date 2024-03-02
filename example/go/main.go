package main

import (
	"fmt"

	"github.com/elewis787/wasmtime-go-nn/example/go/ml"
)

func main() {
	res := ml.WasiNnGraphLoadByName("/Users/elewis/Projects/go/src/github.com/elewis787/wasmtime-go-nn/example/models/Mixtral-8x7B-Instruct-v0.1-Q2_K.gguf")
	if res.IsOk() {
		fmt.Println("value")
	} else {
		fmt.Println("error")
	}
}
