package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := 8080
	http.HandleFunc("/", handler)
	fmt.Printf("Starting service on: http://localhost:%d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

func handler(w http.ResponseWriter, req *http.Request) {
	log.Printf("Received request from: %s\n", req.RemoteAddr)
	fmt.Fprintf(w, "Hello World")
}
