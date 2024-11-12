package main

import (
	"database/sql"
	"fmt"
	"log"
)

func main() {
	fmt.Println("Name Gen")
	connstr := "postgres://oracle:supersecurepassword@localhost:3210/randomtables?sslmode=verify-full"
	db, err := sql.Open("postgres", connstr)
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT name FROM names WHERE Category = Male")
	fmt.Println(rows)

}
