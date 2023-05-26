package bucket

import "github.com/gin-gonic/gin"

type Repository interface {
	UploadBucketImage(c *gin.Context, bucketFolder string, name string) (string, error)
}
