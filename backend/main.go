package main

import (
	"log"
	"net/http"
	"time"

	"github.com/fourtf/studyhub/routing"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	loadEnvironmentVariables()
	router := routing.SetupRouter()
	startServer(router)
}

func loadEnvironmentVariables() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func startServer(router *mux.Router) {
	log.Println("Listening at http://localhost:3001")

	srv := http.Server{
		Addr:         ":3001",
		ReadTimeout:  1 * time.Minute,
		WriteTimeout: 1 * time.Minute,
		Handler:      router,
	}

	srv.ListenAndServe()
}
