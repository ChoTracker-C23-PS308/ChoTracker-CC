package admin

import (
	"context"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
)

func NewFirebaseAuth(app *firebase.App) (*auth.Client, error) {
	auth, err := app.Auth(context.Background())
	if err != nil {
		return nil, err
	}

	return auth, nil
}
