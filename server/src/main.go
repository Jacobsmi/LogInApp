package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"server/src/dbutils"
	"server/src/models"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

func signup(w http.ResponseWriter, r *http.Request) {
	// Create the new user object and read the JSON from the request into new user
	var newUser models.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	// Handle an error in reading the JSON by printing the error in the console and returning an error message
	if err != nil {
		fmt.Println("Error reading JSON in request")
		fmt.Println(err)
		json.NewEncoder(w).Encode(map[string]interface{}{"success": false, "msg": "json_read_err"})
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Println(newUser)
	// Hash password
	passBytes, err := bcrypt.GenerateFromPassword([]byte(newUser.Pass), 8)
	if err != nil {
		fmt.Println("Error reading JSON in request")
		fmt.Println(err)
		json.NewEncoder(w).Encode(map[string]interface{}{"success": false, "msg": "json_read_err"})
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Try to store in the database
	sqlStatement := `INSERT INTO users(fname, lname, email, pass) VALUES($1, $2, $3, $4) RETURNING id`

	_, err = dbutils.DbConn.Exec(sqlStatement, newUser.Fname, newUser.Lname, newUser.Email, string(passBytes))
	if err != nil {
		fmt.Println("Error writing to the DB")
		fmt.Println(err)
		json.NewEncoder(w).Encode(map[string]interface{}{"success": false, "msg": "db_insert_err"})
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// Get the new user's ID from the result of the last query

	// Create a JWT based on ID

	// Return the JWT

	json.NewEncoder(w).Encode(map[string]interface{}{"success": true, "msg": "Testing"})
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/signup", signup)

	defer dbutils.DbConn.Close()

	originsOk := handlers.AllowedOrigins([]string{"http://localhost:3000"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST"})
	headersOk := handlers.AllowedHeaders([]string{"Content-Type"})
	allowCreds := handlers.AllowCredentials()

	fmt.Println("API running at http://localhost:5000")
	// Runs the http server on port 5000 wrapped with a CORS handler
	log.Fatal(http.ListenAndServe(":5000", handlers.CORS(originsOk, methodsOk, allowCreds, headersOk)(r)))
}
