package main

import (
	"fmt"
	"log"
	"net/http"
)

type respHandler struct {
	h http.HandlerFunc
}

func (t respHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	t.h(w, r)
}

func statusCodeHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintln(w, "Fail to parse body")
		return
	}

	names, ok := r.Form["name"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "No name sent!")
		return
	}

	for _, name := range names {
		fmt.Fprintln(w, "Hello")
		fmt.Fprintln(w, name)
	}
}

func errorHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Error!!\n500", http.StatusInternalServerError)
}

func main() {
	http.Handle("/hello", respHandler{statusCodeHandler})
	http.Handle("/error", respHandler{errorHandler})

	err := http.ListenAndServe("127.0.0.1:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
