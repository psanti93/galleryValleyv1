package main

import (
	"fmt"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func main() {

	switch os.Args[1] {
	case "hash":
		//hash the password
		hash(os.Args[2])
	case "compare":
		compare(os.Args[2], os.Args[3])
	default:
		fmt.Printf("Invalid command: %v\n", os.Args[1])
	}

}

func hash(password string) {
	// cost increases the the amount of time of a server to encrypt the password
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Printf("error hashing: %v\n", password)
		return
	}
	fmt.Println(string(hashedBytes)) //prints out the hashed password. hash contains the salt and the cost, verifies password by comparing the salt password
}

func compare(password, hash string) {
	//TODO compare password with the hash

	fmt.Printf("TODO compare the password %q with the hash %q\n", password, hash)
}
