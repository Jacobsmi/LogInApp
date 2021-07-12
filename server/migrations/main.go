package main

import (
	"fmt"
	"server/src/dbutils"
)

func main() {
	defer dbutils.DbConn.Close()

	sqlStatement := `CREATE TABLE users(
		id SERIAL,
		fname VARCHAR NOT NULL,
		lname VARCHAR NOT NULL,
		email VARCHAR NOT NULL UNIQUE,
		pass VARCHAR NOT NULL
	)`
	_, err := dbutils.DbConn.Exec(sqlStatement)
	if err != nil {
		panic(err)
	}
	fmt.Println("Created tables")
}
