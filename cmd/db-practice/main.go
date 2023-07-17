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

	// acquiring a record id example
	// first_name_2 := "John"
	// last_name_2 := "Santiago"
	// age_2 := 24
	// email_2 := "john@santiago.com"

	// row := db.QueryRow(`
	// 	INSERT INTO users(first_name, last_name, age, email)
	// 	VALUES($1,$2,$3,$4) RETURNING id;

	// `, first_name_2, last_name_2, age_2, email_2)
	// var id int
	// err = row.Scan(&id)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("User Created. Id =", id)

	id := 1

	row := db.QueryRow(`
	  SELECT first_name, email
	  FROM users
	  WHERE id=$1;`, id)

	var firstName, emailAddress string
	err = row.Scan(&firstName, &emailAddress)
	if err != nil {
		panic(err)
	}

	fmt.Printf("User information: name==%s, email=%s\n", firstName, emailAddress)

	userID := 1
	for i := 1; i <= 5; i++ {
		amount := i * 100
		desc := fmt.Sprintf("Fake order #%d", i)
		_, err := db.Exec(`
			INSERT INTO orders(user_id,amount,description)
			VALUES($1,$2,$3);`, userID, amount, desc)
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("Created fake orders")

	type Order struct {
		ID          int
		UserID      int
		Amount      int
		Description string
	}

	var orders []Order
	userId := 1

	rows, err := db.Query(`
		SELECT id, amount, description
		FROM orders
		WHERE user_id=$1`, userId)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	// rows.Next() returns true if there is an object in the rows object, and if there's another record to look at
	for rows.Next() {
		var order Order
		order.UserID = userId

		// looks at the id, amount, description fields from the the DB
		err := rows.Scan(&order.ID, &order.Amount, &order.Description)

		if err != nil {
			panic(err)
		}

		orders = append(orders, order)
	}

	//check for an err if there isn't initially anything on the rows.next

	if rows.Err() != nil {
		panic(err)
	}

	fmt.Println("Orders: ", orders)
}
