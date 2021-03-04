package main

import (
	"database/sql"
	"encoding/base64"
	"flag"
	"log"
	"net/http"
	"regexp"
	"time"

	"github.com/fourtf/studyhub/authorization"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/gorilla/schema"
	"github.com/gorilla/sessions"
	abclientstate "github.com/volatiletech/authboss-clientstate"
	"github.com/volatiletech/authboss/v3"
	"github.com/volatiletech/authboss/v3/defaults"
	"github.com/volatiletech/authboss/v3/remember"
)

var (
	flagDebug    = flag.Bool("debug", false, "output debugging information")
	flagDebugDB  = flag.Bool("debugdb", false, "output database on each request")
	flagDebugCTX = flag.Bool("debugctx", false, "output specific authboss related context keys on each request")
	flagAPI      = flag.Bool("api", false, "configure the app to be an api instead of an html app")
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

	router.Use(middleware.Logger, ab.LoadClientStateMiddleware, remember.Middleware(ab))

	router.Mount("/auth", http.StripPrefix("/auth", ab.Core.Router))

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
