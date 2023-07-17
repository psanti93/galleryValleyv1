package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	s := make([]int, 100)
	// will print out the same random number 3 different times b/c rand package has a seed function rand.Seed(1) - default seed is 1
	// seed has number that you pass in that starts off the random number generator, thus a pseudo random generator
	// numbers you pass on will always be the same as a result
	rand.Seed(time.Now().UnixNano()) // to mimic randomness not a good fit b/c it allows an attacker to figure out our seed the can figure out our session tokens easily
	fmt.Println(rand.Intn(len(s)))
	fmt.Println(rand.Intn(len(s)))
	fmt.Println(rand.Intn(len(s)))

}
