package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	// connect to the database
	conn, err := sql.Open("pgx", "host=localhost port=5433 dbname=test_connect user=postgres password=password")
	if err != nil {
		log.Fatal(fmt.Errorf("error opening database connection: %v", err))
	}
	defer conn.Close()

	log.Println("Connected to the database")

	// test my connection
	err = conn.Ping()
	if err != nil {
		log.Fatal(fmt.Errorf("error pinging database: %v", err))
	}

	log.Println("Pinged the database")

	// get rows from table

	// insert a row

	// get rows from table again

	// update a row

	// get rows from table again

	// get one row by id

	// delete a row

	// get rows again
}
