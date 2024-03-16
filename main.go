package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Todo struct {
	Title    string
	Done     bool
	Deadline string
}

func main() {
	fmt.Println("starting server...")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		args := map[string][]Todo{
			"todos": {
				{Title: "Buy milk", Done: true, Deadline: "2024-01-21"},
				{Title: "Buy eggs", Done: false, Deadline: "2024-04-02"},
				{Title: "Buy bread", Done: false, Deadline: "2024-04-03"},
			},
		}

		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.Execute(w, args)
	})

	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		log.Print("adding todo...")

		title := r.PostFormValue("title")
		deadline := r.PostFormValue("deadline")

		log.Printf("title: %s, deadline: %s", title, deadline)

		tmpl := template.Must(template.ParseFiles("index.html"))

		tmpl.ExecuteTemplate(w, "todo-list-element", Todo{
			Title:    title,
			Done:     false,
			Deadline: deadline,
		})

	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
