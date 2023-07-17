package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func main() {
	n := 32
	b := make([]byte, n)
	// before rand
	fmt.Println(b)

	nRead, err := rand.Read(b)

	if nRead < n {
		panic("didnt read enough random bytes")
	}
	if err != nil {
		panic(err)
	}
	// after rand
	fmt.Println(base64.URLEncoding.EncodeToString(b)) // encodes the bytes into a string

}
