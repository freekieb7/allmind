package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"viavia.io/platform/authenticator"
	"viavia.io/platform/database"
	"viavia.io/platform/router"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Failed to load the env vars: %v", err)
	}

	db, err := database.New()
	if err != nil {
		log.Fatalf("Failed to initialize the database: %v", err)
	}

	if err := db.Migrate(); err != nil {
		log.Fatalf("Failed to migrate the database: %v", err)
	}

	auth, err := authenticator.New()
	if err != nil {
		log.Fatalf("Failed to initialize the authenticator: %v", err)
	}

	rtr := router.New(auth)

	log.Print("Server listening on http://localhost:3000/")
	if err := http.ListenAndServe("0.0.0.0:3000", rtr); err != nil {
		log.Fatalf("There was an error with the http server: %v", err)
	}
}
