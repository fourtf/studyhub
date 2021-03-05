package main

import (
	"log"
	"net/http"
	"time"

	"github.com/fourtf/studyhub/routing"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

const (
	connectionString string = "user=postgres dbname=postgres password=studyhub_dev"
)

func main() {
	router := routing.SetupRouter()
	startServer(router)
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
