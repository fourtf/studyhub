package routing

import (
	"log"
	"net/http"

	"github.com/fourtf/studyhub/controllers"
	"github.com/fourtf/studyhub/middlewares"
	"github.com/gorilla/mux"
	"github.com/justinas/nosurf"
)

var router *mux.Router = mux.NewRouter().StrictSlash(true)

//SetupRouter initializes the router by assigning paths and middlewares
func SetupRouter() *mux.Router {
	setupMiddlewares()
	setupPublicPaths()
	setupAuthedPaths()
	return router
}

func setupMiddlewares() {
	router.Use(commonMiddleware)
}

func setupPublicPaths() {
	router.Methods("OPTIONS").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	router.HandleFunc("/register", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/login", controllers.Login).Methods("POST")
}

func setupAuthedPaths() {
	authRouter := router.PathPrefix("/auth").Subrouter()
	authRouter.Use(middlewares.JWTVerify)
}

//CSRF-Protection
func nosurfing(h http.Handler) http.Handler {
	surfing := nosurf.New(h)
	surfing.SetFailureHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Failed to validate CSRF token:", nosurf.Reason(r))
		w.WriteHeader(http.StatusBadRequest)
	}))
	return surfing
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		next.ServeHTTP(w, r)
	})
}
