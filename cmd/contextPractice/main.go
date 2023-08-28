package main

import (
	"context"
	"fmt"
)

func main() {
	// creates a new context
	ctx := context.Background()
	ctx = context.WithValue(ctx, "Favorite-Color", "blue")
	value := ctx.Value("Favorite-Color")

	fmt.Printf("Value : %v", value)
}
