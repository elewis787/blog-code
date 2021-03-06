package main

import (
	"log"
	"net/http"

	"github.com/elewis787/blog-code/blogs/AreWeWasmYet/example02/server"
)

func main() {
	server := &server.Server{}
	mux := http.NewServeMux()
	mux.HandleFunc("/add", server.HandleIncrementCount)
	mux.HandleFunc("/count", server.HandleGetCount)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
