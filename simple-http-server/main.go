package main

import (
	"fmt"
	"net/http"
)

func hello(hwritter http.ResponseWriter, hreq *http.Request) {
	fmt.Fprintf(hwritter, "hello %v", *hreq)
}

func main() {
	http.HandleFunc("/hello", hello)

	fmt.Println("Starting the server...")
	http.ListenAndServe(":8080", nil)
}
