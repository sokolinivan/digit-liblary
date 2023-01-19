package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hi")

	http.HandleFunc("/", handler)

	log.Println("Start HTTP server on port 8081")
	log.Fatal(http.ListenAndServe("localhost:8081", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write(([]byte("Hi")))
}