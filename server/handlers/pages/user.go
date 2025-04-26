package handlers

import (
	"html/template"
	"net/http"
	"street-art/services/cookies"
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
	user := cookies.GetUserCookie(r, w)

	tmpl := template.Must(template.ParseFiles("static/html/profile.html"))
	tmpl.Execute(w, user)
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

func Cart(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("static/html/index.html"))
	tmpl.Execute(w, nil)
}