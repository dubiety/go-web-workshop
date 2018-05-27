package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// Exercise Hello, parameter
func helloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	if name == "" {
		name = "friend"
	}
	fmt.Fprintf(w, "Hello %s!", name)
}

// Exercise Hello, parameter
func parseFormHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal("Fail to parse form data")
	}

	names, ok := r.Form["name"]
	if !ok {
		fmt.Fprintf(w, "Hello friend!")
		return
	}
	for _, name := range names {
		fmt.Fprintf(w, "Hello, %s!\n", name)
	}
}

// Exercise Hello, body
func bodyHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Fail to read body: %v", err)
		return
	}
	name := string(body)
	fmt.Fprintf(w, "Hello, %s!", strings.Trim(name, " "))
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/hello_all", parseFormHandler)
	http.HandleFunc("/hello_body", bodyHandler)

	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}
