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
	query := `INSERT INTO users (first_name, last_name) VALUES ($1, $2) RETURNING id`
	_, err = conn.Exec(query, "John", "Brown")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Inserted a row")

	// get rows from table again
	err = getAllRows(conn)
	if err != nil {
		log.Fatal(fmt.Errorf("error getting rows from table: %v", err))
	}

	// update a row
	stmt := `UPDATE users SET last_name = $1 WHERE id = $2`
	_, err = conn.Exec(stmt, "Brown", 5)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Updated one or more rows")

	// get rows from table again
	err = getAllRows(conn)
	if err != nil {
		log.Fatal(fmt.Errorf("error getting rows from table: %v", err))
	}

	// get one row by id
	query = `SELECT id, first_name, last_name FROM users WHERE id = $1`

	var firstName, lastName string
	var id int

	row := conn.QueryRow(query, 4)
	err = row.Scan(&id, &firstName, &lastName)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("QueryRow returns: ", id, firstName, lastName)

	// delete a row
	query = `DELETE FROM users WHERE id = $1`
	_, err = conn.Exec(query, 10)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Deleted id 10")

	// get rows again
	err = getAllRows(conn)
	if err != nil {
		log.Fatal(fmt.Errorf("error getting rows from table: %v", err))
	}
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
