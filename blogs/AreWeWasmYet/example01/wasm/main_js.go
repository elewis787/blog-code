package main

import (
	"crypto/sha256"
	"encoding/hex"
	"syscall/js"
)

func sha256Hash(this js.Value, args []js.Value) interface{} {
	h := sha256.New()
	h.Write([]byte(args[0].String()))
	return hex.EncodeToString(h.Sum(nil))
}

func main() {
	c := make(chan struct{})
	js.Global().Set("Sha256Hash", js.FuncOf(sha256Hash))
	<-c
}
