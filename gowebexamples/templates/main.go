package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

type Todo struct {
	Title string
	Done  bool
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := TodoPageData{
			PageTitle: "All todos",
			Todos: []Todo{
				{Title: "Task 1", Done: false},
				{Title: "Task 2", Done: true},
				{Title: "Task 3", Done: true},
			},
		}

		tmpl, err := template.ParseFiles("layout.html")
		if err != nil {
			fmt.Fprint(w, "An error has occured")
		}
		tmpl.Execute(w, data)
	})
	
	http.ListenAndServe("localhost:8080", nil)
}
