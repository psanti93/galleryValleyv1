package models

import (
	"database/sql"
	"fmt"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int
	Email    string
	Password string
}

type UserService struct {
	DB *sql.DB
}

func (us *UserService) Create(email, password string) (*User, error) {
	//lower case email
	email = strings.ToLower(email)

	//encrypt password
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("Creating user: %v", err)
	}

	passwordHash := string(hashedBytes)

	row := us.DB.QueryRow(`
		INSERT INTO users (email,password_hash)
		VALUES($1,$2) RETURNING id`, email, passwordHash)

	user := User{
		Email:    email,
		Password: passwordHash,
	}

	err = row.Scan(&user.ID)
	if err != nil {
		return nil, fmt.Errorf("Creating user:%v", err)
	}

	return &user, nil
}

func (us *UserService) Authenticate(email, password string) (*User, error) {

	email = strings.ToLower(email)
	user := User{
		Email: email,
	}

	row := us.DB.QueryRow(`	
		SELECT id, password_hash 
		FROM users WHERE email=$1`, email)
	err := row.Scan(&user.ID, &user.Password)

	if err != nil {
		return nil, fmt.Errorf("Authenticate:%w", err)
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return nil, fmt.Errorf("Authenticate: %w", err)
	}

	return &user, nil
}
