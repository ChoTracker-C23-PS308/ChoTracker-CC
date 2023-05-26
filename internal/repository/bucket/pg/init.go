package pg

import (
	"cloud.google.com/go/storage"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
	"io"
)

type pgBucketRepository struct {
}

const (
	BUCKET_NAME = "dev-chotracker-image"
	PATHKEY     = "configs/var/chotracker-c23-ps308-9eb68c17cf5f.json"
)

var (
	ErrFileNotFound = errors.New("IMAGE_FILE_NOT_FOUND")
	ErrCredential   = errors.New("CREDENTIAL_NOT_AUTHORIZED")
)

func (r pgBucketRepository) UploadBucketImage(c *gin.Context, bucketFolder string, name string) (string, error) {
	file, err := c.FormFile("file")
	if err != nil {
		return " ", ErrFileNotFound
	}
	//gcsname := fmt.Sprintf("%s/%s", bucketFolder, time.Now().Format("20060102-150405"))
	gcsname := fmt.Sprintf("%s/%s", bucketFolder, name)
	ctx := c.Request.Context()

	client, err := storage.NewClient(ctx, option.WithCredentialsFile(PATHKEY))
	if err != nil {
		return "", ErrCredential
	}
	defer client.Close()

	bucket := client.Bucket(BUCKET_NAME)
	obj := bucket.Object(gcsname)

	wc := obj.NewWriter(ctx)
	defer wc.Close()

	uploadedFile, err := file.Open()
	if err != nil {
		return " ", errors.New("FILE_NOT_OPENED")
	}

	if _, err := io.Copy(wc, uploadedFile); err != nil {
		return " ", errors.New("COPY_FILE_FAIL")
	}

	if err := wc.Close(); err != nil {
		return "", errors.New("CLOSE ERR")
	}

	url := fmt.Sprintf("https://storage.googleapis.com/%s/%s", BUCKET_NAME, gcsname)
	return url, nil
}
