# LogInApp

- "Boilerplate" for a web app that has a NextJS client and a Go Server that can authenticate users with JWT

## Prerequistes

- Node.js
- npm
- Golang

## Setting Up the client

- Change directory into the client directory
- Run `npm i` to install npm dependencies

## Setting up the server
- Create a PostgreSQL database to hold the info for the system
- Create a `dbVars.go` file in `server/src/dbutils` with the following content
  ```go
  package dbutils

  const (
    host     = "localhost"
    port     = 5432
    user     = "user"
    password = "your-pass"
    dbname   = "db_name"
  )
  ```
- cd into the migrations folder and `go run main.go` to create the tables for the database
- API is now ready for use, cd into the src directory and `go run main.go` to run the API

## Other Notes

- `nodemon --exec go run main.go --signal SIGTERM` can be used to run the server with live updates (requires nodemon installed)