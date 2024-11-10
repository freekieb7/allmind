package router

import (
	"encoding/gob"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"viavia.io/pkg/auth"
	"viavia.io/pkg/controllers"
)

// New registers the routes and returns the router.
func New() *mux.Router {
	// env := os.Getenv("APP_ENV")

	router := mux.NewRouter()

	sessionKey := os.Getenv("SESSION_KEY")
	store := sessions.NewCookieStore([]byte(sessionKey))
	store.Options.SameSite = http.SameSiteLaxMode

	// Create authenticator
	oauthProvider, err := auth.NewOAuthProvider()
	if err != nil {
		log.Fatalf("Failed to initialize the authenticator: %v", err)
	}

	landingController := controllers.LandingController{
		CookieStore: store,
	}

	authenticationController := controllers.AuthenticationController{
		CookieStore:   store,
		OAuthProvider: oauthProvider,
	}

	homeController := controllers.HomeController{
		CookieStore: store,
	}

	// To store custom types in our cookies,
	// we must first register them using gob.Register
	gob.Register(auth.Profile{})

	router.Use(otelhttp.NewMiddleware("test"))

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))

	// Public pages
	router.HandleFunc("/", landingController.ShowLanding).Methods("GET")

	router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	}).Methods("GET")

	// // Authentication pages
	router.HandleFunc("/login", authenticationController.Login)
	router.HandleFunc("/callback", authenticationController.Callback)
	router.HandleFunc("/logout", authenticationController.Logout)

	router.HandleFunc("/home", homeController.ShowHome).Methods("GET")

	return router
}
