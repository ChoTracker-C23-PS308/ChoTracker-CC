package http

import (
	"cloud.google.com/go/storage"
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
	"io"
	"time"
)

const (
	BUCKET_NAME = "dev-chotracker-image"
	PATHKEY     = "configs/var/chotracker-c23-ps308-9eb68c17cf5f.json"
)

func UploadBucketImage(c *gin.Context, bucketFolder string, image string) string {
	file, err := c.FormFile(image)
	if err != nil {
		return fmt.Sprintf("%s", c.Error(err))
	}
	gcsname := fmt.Sprintf("%s/%s", bucketFolder, time.Now().Format("20060102-150405"))
	ctx := c.Request.Context()

	client, err := storage.NewClient(ctx, option.WithCredentialsFile(PATHKEY))
	if err != nil {
		return fmt.Sprintf("%s", c.Error(err))
	}
	defer client.Close()

	bucket := client.Bucket(BUCKET_NAME)
	obj := bucket.Object(gcsname)

	wc := obj.NewWriter(ctx)
	defer wc.Close()

	uploadedFile, err := file.Open()
	if err != nil {
		return fmt.Sprintf("%s", c.Error(err))
	}

	if _, err := io.Copy(wc, uploadedFile); err != nil {
		return fmt.Sprintf("%s", c.Error(err))
	}

	if err := wc.Close(); err != nil {
		return fmt.Sprintf("%s", c.Error(err))
	}

	url := fmt.Sprintf("https://storage.googleapis.com/%s/%s", BUCKET_NAME, gcsname)
	return url
}
