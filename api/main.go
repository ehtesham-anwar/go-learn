package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("live port :8080")
	mux := http.NewServeMux()
	mux.HandleFunc("/track", handler)
	http.ListenAndServe(":8080", mux)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}
