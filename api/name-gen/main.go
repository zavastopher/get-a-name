package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"math/rand/v2"
	"strconv"
)

func getAnyName(db *sql.DB) {
	category := "Female"
	if rand.IntN(2) != 0 {
		category = "Male"
	}

	count, err := db.Query("SELECT COUNT(name) FROM names WHERE category = $1", category)
	if err != nil {
		log.Println("failed to get count from names table")
		log.Fatal(err)
	}

	var total string
	if count.Next() {
		if err := count.Scan(&total); err != nil {
			log.Println("Error in scanning the total")
			log.Fatal(err)
		}
	}

	totalint, err := strconv.Atoi(total)
	if err != nil {
		log.Fatal(err)
	}

	random := rand.IntN(totalint)

	rows, err := db.Query("SELECT name FROM names;")
	if err != nil {
		log.Println("failed to get names from table")
		log.Fatal(err)
	}

	for random > 0 {
		random -= 1
		rows.Next()
	}
	var name string
	if err := rows.Scan(&name); err != nil {
		log.Println("error in generating name")
		log.Fatal(err)
	}
	fmt.Printf("\nYour name is now %s", name)
}

func main() {
	fmt.Println("connecting to random-tables db")
	connstr := "postgres://oracle:supersecurepassword@localhost:3210/randomtables?sslmode=disable"
	db, err := sql.Open("postgres", connstr)
	defer db.Close()
	if err != nil {
		log.Println("Connection to database failed")
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Println("DB Ping Failed")
		log.Fatal(err)
	}

	fmt.Println("Connected Successfully")

	getAnyName(db)
}
