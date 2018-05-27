package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, web")
}

func byeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Bye, web")
}

// Declare a new type and implement http.Handler interface
type newHelloHandler struct{}

func (newHelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, this is another hello page!")
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/bye", byeHandler)
	http.Handle("/another_hello", newHelloHandler{})

	err := http.ListenAndServe("127.0.0.1:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
