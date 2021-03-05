package main

import (
	"database/sql"
	"encoding/base64"
	"log"
	"net/http"
	"regexp"
	"time"

	"github.com/fourtf/studyhub/authorization"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/gorilla/schema"
	"github.com/gorilla/sessions"
	"github.com/justinas/nosurf"
	_ "github.com/lib/pq"
	abclientstate "github.com/volatiletech/authboss-clientstate"
	"github.com/volatiletech/authboss/v3"
	"github.com/volatiletech/authboss/v3/defaults"
	"github.com/volatiletech/authboss/v3/remember"
)

var (
	ab          = authboss.New()
	database    *sql.DB
	userStorage = authorization.CreatingUserStorer{Database: database}
	schemaDec   = schema.NewDecoder()

	sessionStore abclientstate.SessionStorer
	cookieStore  abclientstate.CookieStorer
)

const (
	sessionCookieName        = "studyhub_session"
	connectionString  string = "user=postgres dbname=postgres password=studyhub_dev"
)

func main() {
	initializeSessionsAndCookies()
	connectToDB()
	setupAuthboss()
	setupRouter()
}
func initializeSessionsAndCookies() {
	//TODO: Replace keys
	cookieStoreKey, _ := base64.StdEncoding.DecodeString(`NpEPi8pEjKVjLGJ6kYCS+VTCzi6BUuDzU0wrwXyf5uDPArtlofn2AG6aTMiPmN3C909rsEWMNqJqhIVPGP3Exg==`)
	sessionStoreKey, _ := base64.StdEncoding.DecodeString(`AbfYwmmt8UCwUuhd9qvfNA9UCuN1cVcKJN1ofbiky6xCyyBj20whe40rJa3Su0WOWLWcPpO1taqJdsEI/65+JA==`)
	cookieStore = abclientstate.NewCookieStorer(cookieStoreKey, nil)
	cookieStore.HTTPOnly = false
	cookieStore.Secure = false
	sessionStore = abclientstate.NewSessionStorer(sessionCookieName, sessionStoreKey, nil)
	cstore := sessionStore.Store.(*sessions.CookieStore)
	cstore.Options.HttpOnly = false
	cstore.Options.Secure = false
	cstore.MaxAge(int((30 * 24 * time.Hour) / time.Second))
}

func connectToDB() {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	database = db
}

func setupAuthboss() {
	ab.Config.Paths.RootURL = "http://localhost:3001"

	ab.Config.Storage.Server = userStorage
	ab.Config.Storage.SessionState = sessionStore
	ab.Config.Storage.CookieState = cookieStore

	defaults.SetCore(&ab.Config, false, false)

	emailRule := defaults.Rules{
		FieldName:  "email",
		Required:   true,
		MatchError: "Must be a valid e-mail address",
		MustMatch:  regexp.MustCompile(`.*@.*\..+`),
	}
	passwordRule := defaults.Rules{
		FieldName: "password",
		Required:  true,
		MinLength: 4,
	}
	nameRule := defaults.Rules{
		FieldName: "name",
		Required:  true,
		MinLength: 2,
	}

	ab.Config.Core.BodyReader = defaults.HTTPBodyReader{ReadJSON: true,
		Rulesets: map[string][]defaults.Rules{
			"register": {emailRule, passwordRule, nameRule},
		}}

	if err := ab.Init(); err != nil {
		panic(err)
	}
}

func setupRouter() {
	schemaDec.IgnoreUnknownKeys(true)

	router := chi.NewRouter()

	//Remove in production
	router.Use(corsMiddleware)

	router.Use(middleware.Logger, nosurfing, ab.LoadClientStateMiddleware, remember.Middleware(ab))

	router.Group(func(router chi.Router) {
		router.Use(authboss.ModuleListMiddleware(ab))
		router.Mount("/auth", http.StripPrefix("/auth", ab.Core.Router))
	})

	// In order to have a "proper" API with csrf protection we allow
	// the options request to return the csrf token that's required to complete the request
	// when using post
	optionsHandler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-CSRF-TOKEN", nosurf.Token(r))
		w.WriteHeader(http.StatusOK)
	}
	// We have to add each of the authboss get/post routes specifically because
	// chi sees the 'Mount' above as overriding the '/*' pattern.
	routes := []string{"register"}
	router.MethodFunc("OPTIONS", "/*", optionsHandler)
	for _, r := range routes {
		router.MethodFunc("OPTIONS", "/auth/"+r, optionsHandler)
	}

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	log.Println("Listening at http://localhost:3001")

	srv := http.Server{
		Addr:         ":3001",
		ReadTimeout:  1 * time.Minute,
		WriteTimeout: 1 * time.Minute,
		Handler:      router,
	}

	srv.ListenAndServe()
}

func nosurfing(h http.Handler) http.Handler {
	surfing := nosurf.New(h)
	surfing.SetFailureHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Failed to validate CSRF token:", nosurf.Reason(r))
		w.WriteHeader(http.StatusBadRequest)
	}))
	return surfing
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", "GET,HEAD,OPTIONS,POST,PUT")
		w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")
		next.ServeHTTP(w, r)
	})
}
