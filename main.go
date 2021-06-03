package main

import (
	"go-site/handlers"
	"go-site/redis_api"
	"go-site/storage"
	"log"
	"net/http"
	"os"
)

func main() {
	storage.Connect(os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	redis_api.Connect(os.Getenv("REDIS_ADDRESS"), os.Getenv("REDIS_PASSWORD"), 0)

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/favicon.ico", handlers.FaviconHandler)

	http.HandleFunc("/registration", handlers.RegistrationHandler)

	http.HandleFunc("/login", handlers.LoginHandler)

	http.HandleFunc("/do/registration", handlers.DoRegistration)

	http.HandleFunc("/do/login", handlers.DoLogin)

	http.HandleFunc("/id/", handlers.PageHandler)

	http.HandleFunc("/", handlers.MainPageHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
