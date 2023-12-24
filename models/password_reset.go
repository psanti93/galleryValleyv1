package models

import (
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"fmt"
	"strings"
	"time"

	"github.com/psanti93/galleryValleyv1/rand"
)

const (
	DefaultResetDuration = 1 * time.Hour
)

type PasswordReset struct {
	ID     int
	UserID int
	// Only set when a password reset is being created
	Token     string
	TokenHash string
	ExpiresAt time.Time
}
type PasswordResetService struct {
	DB *sql.DB
	// BytesPerToken is used to determine how many bytes to use when generating
	// each session token. If this value is not set or is less than the
	// MinBytesPerToken const it will be ignored and MinBytesPerToken will be
	// used.
	BytesPerToken int
	// Duration is the amount of time that a password reset is valid for.
	// Default to DefaultResetDuration
	Duration time.Duration
}

func (passwordResetService *PasswordResetService) Create(email string) (*PasswordReset, error) {
	var userId int
	email = strings.ToLower(email)

	//Verify we have a valid email address for a user and get that user's ID
	row := passwordResetService.DB.QueryRow(`SELECT id FROM users WHERE email=$1;`, email)
	err := row.Scan(&userId)

	if err != nil {
		return nil, fmt.Errorf("Creating Password Reset: %v", err)
	}

	// building the password reset
	bytesPerToken := passwordResetService.BytesPerToken
	if bytesPerToken < MinBytesPerToken {
		bytesPerToken = MinBytesPerToken
	}
	token, err := rand.GenerateSessionToken(bytesPerToken)
	if err != nil {
		return nil, fmt.Errorf("Token generated: %v", err)
	}

	duration := passwordResetService.Duration
	if duration == 0 {
		duration = DefaultResetDuration
	}
	passwordReset := PasswordReset{
		UserID:    userId,
		Token:     token,
		TokenHash: passwordResetService.hashToken(token),
		ExpiresAt: time.Now().Add(duration),
	}

	// Insert password reset into the db
	row = passwordResetService.DB.QueryRow(`
	INSERT INTO password_resets (user_id,token_hash,expires_at) 
	VALUES($1,$2,$3) ON CONFLICT (user_id) 
	DO UPDATE SET token_hash=$2, expires_at=$3 RETURNING id;`,
		passwordReset.UserID, passwordReset.TokenHash, passwordReset.ExpiresAt)

	err = row.Scan(&passwordReset.ID)

	if err != nil {
		return nil, fmt.Errorf("Password Reset failed: %v", err)
	}

	return &passwordReset, nil
}

func (passwordResetService *PasswordResetService) Consume(token string) (*User, error) {

	//1. Validate the token and make sure it's not expired
	tokenHash := passwordResetService.hashToken(token)
	var user User
	var pwReset PasswordReset
	row := passwordResetService.DB.QueryRow(`
		SELECT password_resets.id,
			password_resets.expires_at,
			users.id,
			users.email,
			users.password_hash
		FROM password_resets 
		  JOIN users ON users.id = password_resets.user_id
		WHERE password_resets.token_hash = $1;`, tokenHash)
	err := row.Scan(&pwReset.ID, &pwReset.ExpiresAt, &user.ID, &user.Email, &user.Password)

	if err != nil {
		return nil, fmt.Errorf("Consume:  %v", err)
	}

	// in the event the time is after the password is expired
	if time.Now().After(pwReset.ExpiresAt) {
		return nil, fmt.Errorf("Token Expired %v", token)
	}
	// Deleting the reset token after it's been  used
	err = passwordResetService.delete(pwReset.ID)
	if err != nil {
		return nil, fmt.Errorf("Consume: %w", err)
	}

	return &user, nil
}

func (passwordResetService *PasswordResetService) hashToken(token string) string {
	tokenHash := sha256.Sum256([]byte(token))
	return base64.URLEncoding.EncodeToString(tokenHash[:]) // [:] take the start and end of a byte array and use all the bytes of w/in an array
}

func (passwordResetService *PasswordResetService) delete(id int) error {
	_, err := passwordResetService.DB.Exec(`
		DELETE FROM password_resets
		WHERE id = $1;`, id)
	if err != nil {
		return fmt.Errorf("delete: %w", err)
	}

	return nil
}
