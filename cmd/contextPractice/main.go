package main

import (
	"context"
	"fmt"
	"strings"
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
	intValue, ok := value.(int) // asserts that it is of type string
	if !ok {
		fmt.Println("it isn't an int")
	} else {
		fmt.Println(intValue + 4)
	}

	strValue, ok := value.(string)
	if !ok {
		fmt.Println("Not a string value")
	} else {
		fmt.Println(value)
		fmt.Println(strings.HasPrefix(strValue, "b"))
	}

}
