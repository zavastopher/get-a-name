package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

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
	log.Println("DB Connection started successfully")

	rows, err := db.Query("SELECT name FROM names WHERE Culture = 'Celtic';")
	if err != nil {
		log.Println("Select query failed")
		log.Fatal(err)
	}
	fmt.Printf("%t", rows.Next())
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			log.Fatal(err)
		}
		fmt.Println(name)
	}
}
