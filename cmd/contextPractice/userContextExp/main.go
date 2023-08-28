package main

import (
	stdctx "context"
	"fmt"

	"github.com/psanti93/galleryValleyv1/context"
	"github.com/psanti93/galleryValleyv1/models"
)

func main() {
	ctx := stdctx.Background()

	user := models.User{
		Email: "paul.santiago@mail.com",
	}

	ctx = context.WithUser(ctx, &user)

	value := context.User(ctx)

	fmt.Println(value.Email)
}
