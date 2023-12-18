package rand

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func Bytes(n int) ([]byte, error) {
	b := make([]byte, n) // an array of 0's based one the size ex: length of 5 [0,0,0,0,0]

	nRead, err := rand.Read(b) // creates an array of random bytes based on the size ex: [158 15 94 223 155}
	if err != nil {
		return nil, fmt.Errorf("bytes: %w", err)
	}

	if nRead < n {
		return nil, fmt.Errorf("bytes: didn't read enough random bytes")
	}

	return b, nil
}

// SessionToken returns a random string using crypto/rand.
// n is the number of bytes being used to generate the random string

func GenerateSessionToken(n int) (string, error) {
	b, err := Bytes(n)
	if err != nil {
		return "", fmt.Errorf("string: %w", err)
	}
	return base64.URLEncoding.EncodeToString(b), nil
}
