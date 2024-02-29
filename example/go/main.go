package main

import (
	"fmt"
	"log"

	"github.com/elewis787/wasmtime-go-nn/ml"
)

func main() {
	res := ml.WasiNnGraphLoadByName("value")
	if res.IsOk() {
		graph := res.Unwrap()
		fmt.Println(graph)
	} else {
		log.Fatal(res.UnwrapErr())
	}
}
