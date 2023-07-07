package main

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v4/stdlib"
)

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
	SSLMode  string
}

func (cfg PostgresConfig) String() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DbName, cfg.SSLMode)
}

func main() {

	cfg := PostgresConfig{
		Host:     "localhost",
		Port:     "3356",
		User:     "mike",
		Password: "ditka",
		DbName:   "galleyvalley",
		SSLMode:  "disable",
	}

	db, err := sql.Open("pgx", cfg.String())

	if err != nil {
		panic(err)
	}
	defer db.Close()
	if err := db.Ping(); err != nil {
		panic(err)
	}

	fmt.Println("Connected")

	//NOTE always do docker compose down before running stuff

	// Create a table....
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			age INT,
			first_name TEXT,
			last_name TEXT,
			email TEXT UNIQUE NOT NULL
		);

		CREATE TABLE IF NOT EXISTS orders(
			id SERIAL PRIMARY KEY,
			user_id INT NOT NULL,
			amount INT,
			description TEXT
		);

	`)

	if err != nil {
		panic(err)
	}

	fmt.Println("Tables Created!")

	// Inserting a user...
	first_name := "Paul"
	last_name := "Santiago"
	age := 29
	email := "paul@santiago.com"
	_, err = db.Exec(`
		INSERT INTO users(first_name, last_name, age, email)
		VALUES($1,$2,$3,$4);
	
	`, first_name, last_name, age, email)

	if err != nil {
		panic(err)
	}

	fmt.Println("User Created")

}
