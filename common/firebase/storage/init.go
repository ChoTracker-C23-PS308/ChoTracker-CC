package storage

import (
	gcs "cloud.google.com/go/storage"
	"context"

	firebase "firebase.google.com/go/v4"
)

func NewFirebaseStorage(app *firebase.App, bucket string) (*gcs.BucketHandle, error) {
	str, err := app.Storage(context.Background())
	if err != nil {
		return nil, err
	}
	bkt, err := str.Bucket(bucket)
	if err != nil {
		return nil, err
	}

	return bkt, nil
}
