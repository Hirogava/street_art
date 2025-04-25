package handlers

import (
	"html/template"
	"net/http"
)

func Products(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("static/html/index.html"))
	tmpl.Execute(w, nil)
}

func Product(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("static/html/index.html"))
	tmpl.Execute(w, nil)
}

func ProductsByCategory(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("static/html/index.html"))
	tmpl.Execute(w, nil)
}

func Cart(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("static/html/index.html"))
	tmpl.Execute(w, nil)
}
