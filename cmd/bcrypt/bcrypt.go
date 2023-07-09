package main

import (
	"fmt"
	"os"
)

func main() {

	/**
		loops through all the args passed via the command line

		Result:
		what I run:
		$ ./bcrypt compare "some password" "some hash value"

		Args:
		0 C:\Users\Paul\git\galleryValleyv1\bcrypt.exe
		1 compare
		2 some password
		3 some hash value

	**/

	for i, arg := range os.Args {
		fmt.Println(i, arg)
	}

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
	//TODO hash password

	fmt.Printf("TODO hash password %q\n", password)
}

func compare(password, hash string) {
	//TODO compare password with the hash

	fmt.Printf("TODO compare the password %q with the hash %q\n", password, hash)
}
