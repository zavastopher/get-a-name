package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	fmt.Println("Name Gen")
	connstr := "postgres://oracle:supersecurepassword@localhost:3210/random-tables?sslmode=verify-full"
	db, err := sql.Open("postgres", connstr)
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT name FROM names WHERE category = Male")
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			log.Fatal(err)
		}
		log.Printf("%s", name)
	}
}
