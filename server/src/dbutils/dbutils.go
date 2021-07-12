package dbutils

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var DbConn *sql.DB

// Create a connection to the database when imported
func init() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	var err error
	DbConn, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = DbConn.Ping()
	if err != nil {
		panic(err)
	}
}
