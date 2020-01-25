package main

import (
	"database/sql"
	"log"
)

func connect() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:8889)/school")

	if err != nil {
		log.Fatal(err)
	}

	return db
}

