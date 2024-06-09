package main

import (
	"fmt"
	"net/http"
	"log"
)

func logging(f http.HandlerFunc) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		f(w, r)
	}
}

func foo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "foo")
}

func bar(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "bar")
}

func main() {
	http.HandleFunc("/foo", logging(foo))
	http.HandleFunc("/bar", logging(bar))

	http.ListenAndServe("localhost:8080", nil)
}