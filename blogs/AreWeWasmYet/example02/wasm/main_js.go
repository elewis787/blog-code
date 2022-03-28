package main

import (
	"syscall/js"

	"github.com/elewis787/blog-code/blogs/AreWeWasmYet/example02/client"
)

type jsWrapperCounterClient struct {
	c *client.CounterClient
}

func (j *jsWrapperCounterClient) IncrementCounter() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		handler := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			resolve := args[0]
			reject := args[1]
			go func() {
				if err := j.c.IncrementCounter(); err != nil {
					errorConstructor := js.Global().Get("Error")
					errorObject := errorConstructor.New(err.Error())
					reject.Invoke(errorObject)
				}
				resolve.Invoke("")
			}()
			return nil
		})
		promise := js.Global().Get("Promise")
		return promise.New(handler)
	})
}

func (j *jsWrapperCounterClient) Count() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		handler := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			resolve := args[0]
			reject := args[1]
			go func() {
				v, err := j.c.GetCount()
				if err != nil {
					errorConstructor := js.Global().Get("Error")
					errorObject := errorConstructor.New(err.Error())
					reject.Invoke(errorObject)
				}
				resolve.Invoke(v)
			}()
			return nil
		})
		promise := js.Global().Get("Promise")
		return promise.New(handler)
	})
}

func newCounter(this js.Value, args []js.Value) interface{} {
	jsWrapper := &jsWrapperCounterClient{
		c: client.New(),
	}
	return js.ValueOf(map[string]interface{}{
		"IncrementCounter": jsWrapper.IncrementCounter(),
		"Count":            jsWrapper.Count(),
	})
}

func main() {
	c := make(chan struct{})
	js.Global().Set("NewCounter", js.FuncOf(newCounter))
	<-c
}
