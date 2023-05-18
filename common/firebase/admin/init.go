package admin

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

func NewFirebaseAdmin(credentialType string, credentialValue string) (*firebase.App, error) {
	var opt option.ClientOption

	if credentialType == "file" {
		opt = option.WithCredentialsFile(credentialValue)
	} else if credentialType == "json" {
		opt = option.WithCredentialsJSON([]byte(credentialValue))
	} else {
		return nil, fmt.Errorf("unsupported FIREBASE_CREDENTIAL_TYPE")
	}

	// Firebase Admin SDK
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, err
	}

	return app, nil
}
