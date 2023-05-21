package http

import (
	"cloud.google.com/go/storage"
	"fmt"
	httpCommon "github.com/ChoTracker-C23-PS308/ChoTracker-CC/common/http"
	uModel "github.com/ChoTracker-C23-PS308/ChoTracker-CC/internal/model/user"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
	"io"
	http "net/http"
	"time"
)

func (d HTTPUserDelivery) uploadProfilePict(c *gin.Context) {
	const bucketName = "dev-bucket-1605"
	const pathKey = "configs/gcloud/fadhil123-60d257749e8f.json"

	file, err := c.FormFile("file")
	if err != nil {
		c.Error(err)
		return
	}

	gcsname := time.Now().Format("20060102-150405")
	ctx := c.Request.Context()

	client, err := storage.NewClient(ctx, option.WithCredentialsFile(pathKey))
	if err != nil {
		c.Error(err)
		return
	}
	defer client.Close()

	bucket := client.Bucket(bucketName)
	obj := bucket.Object(gcsname)

	wc := obj.NewWriter(ctx)
	defer wc.Close()

	uploadedFile, err := file.Open()
	if err != nil {
		c.Error(err)
		return
	}

	if _, err := io.Copy(wc, uploadedFile); err != nil {
		c.Error(err)
		return
	}

	if err := wc.Close(); err != nil {
		c.Error(err)
		return
	}

	url := fmt.Sprintf("https://storage.googleapis.com/%s/%s", bucketName, gcsname)
	c.JSON(http.StatusCreated, httpCommon.Response{
		Message: url,
	})
}

func (d HTTPUserDelivery) addUser(c *gin.Context) {
	ctx := c.Request.Context()
	au := c.MustGet(httpCommon.AUTH_USER).(uModel.AuthUser)

	var user httpCommon.AddUser
	if err := c.ShouldBindJSON(&user); err != nil {
		c.Error(err).SetType(gin.ErrorTypeBind)
		return
	}

	nid, err := d.userRepo.CreateUser(ctx, uModel.AddUser{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		BirthDate: user.BirthDate,
		Gender:    user.Gender,
		ImageUrl:  user.ImageUrl,
	}, au)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, httpCommon.Response{
		Data: nid,
	})
}

func (d HTTPUserDelivery) getUser(c *gin.Context) {
	ctx := c.Request.Context()
	au := c.MustGet(httpCommon.AUTH_USER).(uModel.AuthUser)

	id := c.Param("id")

	u, err := d.userRepo.GetUser(ctx, id, au)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, httpCommon.Response{Data: u})
}

func (d HTTPUserDelivery) updateUser(c *gin.Context) {
	ctx := c.Request.Context()
	au := c.MustGet(httpCommon.AUTH_USER).(uModel.AuthUser)

	id := c.Param("id")

	var user httpCommon.UpdateUser
	if err := c.ShouldBindJSON(&user); err != nil {
		c.Error(err).SetType(gin.ErrorTypeBind)
		return
	}

	nid, err := d.userRepo.UpdateUser(ctx, uModel.UpdateUser{
		ID:        id,
		Name:      user.Name,
		Email:     user.Email,
		BirthDate: user.BirthDate,
		Gender:    user.Gender,
	}, au)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, httpCommon.Response{
		Data: nid,
	})
}
