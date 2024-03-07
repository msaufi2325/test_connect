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
	err = getAllRows(conn)
	if err != nil {
		log.Fatal(fmt.Errorf("error getting rows from table: %v", err))
	}

	// insert a row

	// get rows from table again

	// update a row

	// get rows from table again

	// get one row by id

	// delete a row

	// get rows again
}

func getAllRows(conn *sql.DB) error {
	rows, err := conn.Query("SELECT id, first_name, last_name FROM users")
	if err != nil {
		log.Println(err)
		return err
	}
	defer rows.Close()

	var firstName, lastName string
	var id int

	for rows.Next() {
		err := rows.Scan(&id, &firstName, &lastName)
		if err != nil {
			log.Println(err)
			return err
		}
		fmt.Println("Record is: ", id, firstName, lastName)
	}

	if err = rows.Err(); err != nil {
		log.Fatal("Error scanning rows: ", err)
	}

	fmt.Println("------------------------------")

	return nil
}
