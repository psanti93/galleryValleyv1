package models

import (
	"database/sql"
	"fmt"

	"github.com/psanti93/galleryValleyv1/rand"
)

type Session struct {
	ID     int
	UserID int
	// Token is only set when creating a new session. when looking up a session
	// this will be left empty, as we only store the hash of
	// a session token in our DB and we cannot reverse it into a raw token
	Token     string
	TokenHash string
}

type SessionService struct {
	DB *sql.DB
}

func (ss *SessionService) Create(userID int) (*Session, error) {
	//TODO: Create the session token
	token, err := rand.SessionToken()

	if err != nil {
		return nil, fmt.Errorf("Creating Session: %w", err)
	}

	// TODO: hash the session token

	session := Session{
		UserID: userID,
		Token:  token,
		//TODO set th the token hash
	}

	//TODO store session in DB

	//TODO: Implement SessionService.Create

	return &session, nil

}

func (ss *SessionService) User(token string) (*User, error) {
	// TODO: Implement SessionService.User
	return nil, nil
}
