package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"street-art/db"
	"street-art/routes"
	"street-art/services"
	"street-art/services/cookies"

	"github.com/gorilla/mux"
)

func main() {
	services.LoadEnvFile(".env")
	manager := db.NewDBManager("postgres", os.Getenv("DB_CONNECTION_STRING"))
	db.Migrate(manager)

	log.Println("База данных успешно инициализирована и мигрирована.")
	defer manager.Close()

	cookies.Init(os.Getenv("SESSION_KEY"))
	r := mux.NewRouter()

	routes.Init(r, manager)

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	serverPort := os.Getenv("SERVER_PORT")
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", serverPort),
		Handler: r,
	}

	log.Println("Сервер запущен на порту " + serverPort)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}