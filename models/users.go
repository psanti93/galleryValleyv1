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
