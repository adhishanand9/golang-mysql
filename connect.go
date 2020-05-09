package main

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
)

func main() {
	_ = mysql.MySQLDriver{}

	db, err := sql.Open("mysql",
		"admin:BWIewRdNAf6Tn8z81ly6@tcp(test.chwye9de94ch.eu-west-2.rds.amazonaws.com:3306)/project")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	// Open doesn't open a connection. Validate DSN data:
	err = db.Ping()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
}
