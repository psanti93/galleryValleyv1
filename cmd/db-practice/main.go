package main

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {

	db, err := sql.Open("pgx", "host=localhost port=3356 user=mike password=ditka dbname=galleyvalley sslmode=disable")

	if err != nil {
		panic(err)
	}
	defer db.Close()
	if err := db.Ping(); err != nil {
		panic(err)
	}

	fmt.Println("Connected")
}
