package main

import (
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/bytecodealliance/wasmtime-go"
)

func GetCount() {
	req, err := http.NewRequest("get", "http://localhost:8080/count", nil)
	if err != nil {
		log.Println(err)
	}
	httpClient := http.DefaultClient
	resp, err := httpClient.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	v, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	count, _ := strconv.Atoi(string(v))
	log.Println(count)
}

func IncrementCount() {
	req, err := http.NewRequest("put", "http://localhost:8080/add", nil)
	if err != nil {
		log.Println(err)
	}
	httpClient := http.DefaultClient
	resp, err := httpClient.Do(req)
	if err != nil {
		log.Println(err)
	}
	if resp.StatusCode != http.StatusOK {
		log.Println(errors.New(resp.Status))
	}
}

func main() {
	// Almost all operations in wasmtime require a contextual `store`
	// argument to share, so create that first
	engine := wasmtime.NewEngine()
	store := wasmtime.NewStore(engine)
	linker := wasmtime.NewLinker(engine)
	linker.DefineWasi()
	linker.FuncNew("env", "main.getCount", wasmtime.NewFuncType([]*wasmtime.ValType{wasmtime.NewValType(wasmtime.KindI32)}, []*wasmtime.ValType{}), func(caller *wasmtime.Caller, args []wasmtime.Val) ([]wasmtime.Val, *wasmtime.Trap) {
		GetCount()
		return []wasmtime.Val{}, nil
	})
	linker.FuncNew("env", "main.IncrementCount", wasmtime.NewFuncType([]*wasmtime.ValType{wasmtime.NewValType(wasmtime.KindI32)}, []*wasmtime.ValType{}), func(caller *wasmtime.Caller, args []wasmtime.Val) ([]wasmtime.Val, *wasmtime.Trap) {
		IncrementCount()
		return []wasmtime.Val{}, nil
	})

	wasm, err := os.ReadFile("./main.wasm")
	check(err)
	// Once we have our binary `wasm` we can compile that into a `*Module`
	// which represents compiled JIT code.
	module, err := wasmtime.NewModule(store.Engine, wasm)
	check(err)

	// Next up we instantiate a module which is where we link in all our
	// imports. We've got one import so we pass that in here.
	instance, err := linker.Instantiate(store, module)
	check(err)

	// After we've instantiated we can lookup our `run` function and call
	// it.
	run := instance.GetFunc(store, "_start")
	if run == nil {
		panic("not a function")
	}
	_, err = run.Call(store)
	check(err)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
