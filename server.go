package main

import (
	"functions/backend/config"
	"functions/backend/handler/auth/signin"
	"functions/backend/handler/auth/signup"
	"log"
	"net/http"
)

// HTTP server for the handler
func main() {
	// check server configuration
	configErrors := config.CheckServerConfig()
	if configErrors != nil {
		for _, err := range configErrors {
			log.Printf(err.Error())
		}
		log.Fatal("killing the server")
	}

	mux := http.NewServeMux()

	// register handlers
	mux.HandleFunc("/signin", signin.Handler)
	mux.HandleFunc("/signup", signup.Handler)

	err := http.ListenAndServe(":3000", mux)
	log.Fatal(err)
}
