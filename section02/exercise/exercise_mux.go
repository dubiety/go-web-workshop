package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func nameHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	fmt.Fprintf(w, "Hello %v!\n", name)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/hello/{name}", nameHandler)

	err := http.ListenAndServe("127.0.0.1:8080", r)
	if err != nil {
		log.Fatal(err)
	}
}
