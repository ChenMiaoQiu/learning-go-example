package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/ChenMiaoQiu/learning-go-example/ch4/4.14/issue"
)

func main() {
	http.HandleFunc("/", Handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func Handler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Println(err)
	}
	q := r.FormValue("key")
	result, err := issue.SearchIssues(q)
	if err != nil {
		log.Println(err)
	}
	tmpl := template.Must(template.ParseFiles("index.html"))
	if err := tmpl.Execute(w, result); err != nil {
		log.Println(err)
	}
}
