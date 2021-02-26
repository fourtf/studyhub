package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	log.Println("Listening at http://localhost:3001")

	srv := http.Server{
		Addr:         ":3001",
		ReadTimeout:  1 * time.Minute,
		WriteTimeout: 1 * time.Minute,
		Handler:      r,
	}

	srv.ListenAndServe()
}
