package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested for %s\n", r.URL.String())
	})

	http.ListenAndServe("localhost:8080", nil)
}
