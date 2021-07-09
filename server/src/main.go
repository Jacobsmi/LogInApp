package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func signup(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]interface{}{"success": true})
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/signup", signup)

	originsOk := handlers.AllowedOrigins([]string{"http://localhost:3000"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST"})
	headersOk := handlers.AllowedHeaders([]string{"Content-Type"})
	allowCreds := handlers.AllowCredentials()

	fmt.Println("API running at http://localhost:5000")
	// Runs the http server on port 5000 wrapped with a CORS handler
	log.Fatal(http.ListenAndServe(":5000", handlers.CORS(originsOk, methodsOk, allowCreds, headersOk)(r)))
}
