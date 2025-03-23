package model

import (
	"database/sql"
	"fmt"
	"log"
)

var con *sql.DB

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:1234@(tcp:localhost:6379)")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Conneted to the database")
	return db
}
