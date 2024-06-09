package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/sessions"
)

var (
	key = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
)

func login(w http.ResponseWriter, r * http.Request) {
	sessions, _ := store.Get(r, "cookie-name")

	sessions.Values["Authenticated"] = true
	sessions.Save(r, w)
}

func logout(w http.ResponseWriter, r * http.Request) {
	sessions, _ := store.Get(r, "cookie-name")

	sessions.Values["Authenticated"] = false
	sessions.Save(r, w)
}

func secret(w http.ResponseWriter, r * http.Request) {
	sessions, _ := store.Get(r, "cookie-name")

	if auth, ok := sessions.Values["Authenticated"].(bool); !auth || !ok {
		http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
		return
	}

	fmt.Fprint(w, "The cake is a lie!")
}

func main() {
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/secret", secret)	

	http.ListenAndServe("localhost:8080", nil)
}