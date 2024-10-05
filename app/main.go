package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hi"))
	})

	log.Fatal(http.ListenAndServe("0.0.0.0:3000", nil))

	// db, err := database.New()
	// if err != nil {
	// 	log.Fatalf("Failed to initialize the database: %v", err)
	// }

	// if err := db.Migrate(); err != nil {
	// 	log.Fatalf("Failed to migrate the database: %v", err)
	// }

	// auth, err := authenticator.New()
	// if err != nil {
	// 	log.Fatalf("Failed to initialize the authenticator: %v", err)
	// }

	// rtr := router.New(auth)

	// if err := http.ListenAndServe("0.0.0.0:3000", rtr); err != nil {
	// 	log.Fatalf("There was an error with the http server: %v", err)
	// }
}
