package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Age       int    `json:"age"`
}

func encode(w http.ResponseWriter, _r *http.Request) {
	peter := User{
		Firstname: "Peter",
		Lastname:  "Griffin",
		Age:       40,
	}

	json.NewEncoder(w).Encode(peter)
}

func decode(w http.ResponseWriter, r *http.Request) {
	var user User

	json.NewDecoder(r.Body).Decode(&user)

	fmt.Fprintf(w, "%s %s is %d years old!", user.Firstname, user.Lastname, user.Age)
}

func main() {
	http.HandleFunc("/encode", encode)
	http.HandleFunc("/decode", decode)

	http.ListenAndServe("localhost:8080", nil)
}
