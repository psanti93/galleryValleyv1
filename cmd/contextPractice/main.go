package main

import (
	"context"
	"fmt"
)

// type should never be exported
type ctxKey string

const (
	// keys should never be exported
	favoriteColorKey ctxKey = "favorite-color"
)

func main() {
	// creates a new context
	ctx := context.Background()
	ctx = context.WithValue(ctx, favoriteColorKey, "blue")
	//value := ctx.Value("favorite-color") // would return nil because "favorite-color" is of type string and not type ctx key
	value := ctx.Value(favoriteColorKey)
	fmt.Printf("Value : %v", value)
}
