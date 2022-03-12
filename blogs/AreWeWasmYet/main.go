package main

import (
	"log"
	"net/http"

	"github.com/elewis787/blog-code/blogs/AreWeWasm/server"
)

func main() {
	server := &server.Server{}
	mux := http.NewServeMux()
	mux.HandleFunc("/add", server.HandleAddtoCount)
	mux.HandleFunc("/count", server.HandleGetCount)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
