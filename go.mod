module github.com/elewis787/wasmtime-go-nn

go 1.21.5

replace github.com/go-skynet/go-llama.cpp => ../../go-skynet/go-llama.cpp

require (
	github.com/bytecodealliance/wasmtime-go/v18 v18.0.0
	github.com/go-skynet/go-llama.cpp v0.0.0-20231009155254-aeba71ee8428
)
