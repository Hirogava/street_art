package handlers

import (
	"html/template"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("static/html/register.html"))
	tmpl.Execute(w, nil)
}

func Login(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("static/html/login.html"))
	tmpl.Execute(w, nil)
}

func Profile(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("static/html/profile.html"))
	tmpl.Execute(w, nil)
}

func EditProfile(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("static/html/index.html"))
	tmpl.Execute(w, nil)
}

func Orders(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("static/html/index.html"))
	tmpl.Execute(w, nil)
}

func OrderDetails(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("static/html/index.html"))
	tmpl.Execute(w, nil)
}