package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/fourtf/studyhub/routing"
	"github.com/fourtf/studyhub/utils"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	loadDotEnv()
	db, err := utils.ConnectToDB()
	if err != nil {
		os.Exit(1)
	}
	router := routing.SetupRouter(db)
	startServer(router)
}

func loadDotEnv() {
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
