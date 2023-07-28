package models

import (
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"fmt"

	"github.com/psanti93/galleryValleyv1/rand"
)

const (
	// Minimum number of bytes to be used for each session token
	MinBytesPerToken = 32
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
	// BytesPerToken is used to determine how many bytes to use when generating
	// each session token. If this value is not set or is less than the
	// MinBytesPerToken const it will be ignored and MinBytesPerToken will be
	// used.
	BytesPerToken int
}

func (ss *SessionService) Create(userID int) (*Session, error) {
	//TODO: Create the session token
	// give option for someone to put value of bytes per token, if empty default to min bytes
	bytesPerToken := ss.BytesPerToken
	if bytesPerToken < MinBytesPerToken {
		bytesPerToken = MinBytesPerToken
	}
	token, err := rand.GenerateSessionToken(bytesPerToken)

	if err != nil {
		return nil, fmt.Errorf("Creating Session: %w", err)
	}

	session := Session{
		UserID: userID,
		Token:  token,
		//TODO set the token hash

		TokenHash: ss.hash(token),
	}

	//(a user is currently limited to one session upon signing up or signing in, will cause an error)
	// HOW TO FIX:
	// Option 1:
	//1. Query for a user
	//2. If found, update user's session
	//3. if not found create new session for user

	// Option 2 - PSQL specfic:
	//1. try to update the user's session
	//2. if err, create a new session

	row := ss.DB.QueryRow(`
	UPDATE sessions
	SET token_hash = $2
	WHERE user_id=$1
	RETURNING id;`, session.UserID, session.TokenHash)
	err = row.Scan(&session.ID)

	if err == sql.ErrNoRows {
		row = ss.DB.QueryRow(`
		INSERT INTO sessions (user_id,token_hash) 
		VALUES ($1,$2)
		RETURNING id;
	`, session.UserID, session.TokenHash)
		err = row.Scan(&session.ID)

	}

	if err != nil {
		return nil, fmt.Errorf("Inserting session token: %w", err)
	}

	return &session, nil

}

func (ss *SessionService) User(token string) (*User, error) {
	// TODO: Optimize sql query using join
	var user User

	// 1. hash the session token
	tokenHash := ss.hash(token)
	// 2. query the session with that hash and populate user

	row := ss.DB.QueryRow(`
	    SELECT users.id, users.email, users.password_hash FROM users
		JOIN sessions ON users.id=sessions.user_id WHERE sessions.token_hash=$1;
	`, tokenHash)

	err := row.Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		return nil, fmt.Errorf("user: %w", err)
	}

	// 3. return the user
	return &user, nil
}

func (ss *SessionService) Delete(token string) error {
	tokenHash := ss.hash(token)
	_, err := ss.DB.Exec(`
		DELETE FROM sessions
		WHERE token_hash = $1;
	`, tokenHash)

	if err != nil {
		return fmt.Errorf("delete: %w", err)
	}
	return nil
}

func (ss *SessionService) hash(token string) string {
	tokenHash := sha256.Sum256([]byte(token))

	return base64.URLEncoding.EncodeToString(tokenHash[:]) // [:] take the start and end of a byte array and use all the bytes of w/in an array
}

// TODO Token Manger
