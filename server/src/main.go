package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"server/src/dbutils"
	"server/src/models"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

type claims struct {
	id int
	jwt.StandardClaims
}

var jwtKey = []byte("my-sercret-key")

func signup(w http.ResponseWriter, r *http.Request) {
	// Create the new user object and read the JSON from the request into new user
	var newUser models.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	// Handle an error in reading the JSON by printing the error in the console and returning an error message
	if err != nil {
		fmt.Println("Error reading JSON in request")
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"success": false, "msg": "json_read_err"})
		return
	}

	// Hash password
	passBytes, err := bcrypt.GenerateFromPassword([]byte(newUser.Pass), 8)
	if err != nil {
		fmt.Println("Error reading JSON in request")
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{"success": false, "msg": "hash_gen_error"})
		return
	}

	// Try to store in the database
	// Get the new user's ID from the result of the query
	sqlStatement := `INSERT INTO users(fname, lname, email, pass) VALUES($1, $2, $3, $4) RETURNING id`
	// Execute the query and get the return back from the database
	row := dbutils.DbConn.QueryRow(sqlStatement, newUser.Fname, newUser.Lname, newUser.Email, string(passBytes))
	err = row.Scan(&newUser.Id)
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			if string(pgErr.Code) == "23505" {
				fmt.Println("Duplicate error")
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(map[string]interface{}{"success": false, "msg": "duplicate_error"})
				return
			}
		}
		// Catch Duplicate error here
		fmt.Println("Error writing to the DB")
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{"success": false, "msg": "db_insert_err"})
		return
	}
	// Create an expiration time for the JWT
	expTime := time.Now().Add(60 * time.Minute)
	// Create a list of claims for the JWT specifically the id and an expiration time
	jwtClaims := claims{
		id: newUser.Id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime.Unix(),
		},
	}
	// Generates the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims)
	// Returns the complete signed token
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		fmt.Println("Error creating the token")
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{"success": false, "msg": "token_gen_error"})
		return
	}
	// Set the cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    tokenString,
		Expires:  expTime,
		HttpOnly: true,
		Secure:   true,
	})

	json.NewEncoder(w).Encode(map[string]interface{}{"success": true, "msg": nil})
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
