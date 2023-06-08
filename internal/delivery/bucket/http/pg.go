package http

import (
	"cloud.google.com/go/storage"
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
	"io"
)

const (
	BUCKET_NAME = "dev-chotracker-image"
	PATHKEY     = "configs/var/chotracker-c23-ps308-9eb68c17cf5f.json"
)

func UploadBucketImage(c *gin.Context, bucketFolder string, image string) (string, error) {
	//config := cfg.LoadConfig()
	file, err := c.FormFile("file")
	if err != nil {
		return "", err
	}
	gcsname := fmt.Sprintf("%s/%s", bucketFolder, image)
	ctx := c.Request.Context()

	client, err := storage.NewClient(ctx, option.WithCredentialsFile(PATHKEY))
	//client, err := storage.NewClient(ctx, option.WithCredentialsFile(config.Bucket.CredentialValue))
	if err != nil {
		return "", err
	}
	defer client.Close()

	bucket := client.Bucket(BUCKET_NAME)
	//bucket := client.Bucket(config.Bucket.Name)
	obj := bucket.Object(gcsname)

	wc := obj.NewWriter(ctx)
	defer wc.Close()

	uploadedFile, err := file.Open()
	if err != nil {
		return "", err
	}

	if _, err := io.Copy(wc, uploadedFile); err != nil {
		return fmt.Sprintf("%s", c.Error(err)), err
	}

	if err := wc.Close(); err != nil {
		return "", err
	}

	url := fmt.Sprintf("https://storage.googleapis.com/%s/%s", BUCKET_NAME, gcsname)
	return url, nil
}
