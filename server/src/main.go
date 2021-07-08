package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Runs the http server on port 5000 wrapped with a CORS handler
	log.Fatal(http.ListenAndServe(":5000", handlers.CORS(handlers.AllowCredentials(), handlers.AllowedOrigins([]string{"http://localhost:3000"}))(r)))
}
