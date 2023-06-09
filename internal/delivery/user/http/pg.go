package http

import (
	httpCommon "github.com/ChoTracker-C23-PS308/ChoTracker-CC/common/http"
	bucket "github.com/ChoTracker-C23-PS308/ChoTracker-CC/internal/delivery/bucket/http"
	uModel "github.com/ChoTracker-C23-PS308/ChoTracker-CC/internal/model/user"
	"github.com/gin-gonic/gin"
	http "net/http"
)

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
		Data:    nid,
		Message: "Create User Succesfuly",
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
	c.JSON(http.StatusOK, httpCommon.Response{Data: u, Message: "Get user Succesfuly"})
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
		Data:    nid,
		Message: "Update User Succesfuly",
	})
}

func (d HTTPUserDelivery) uploadProfilePict(c *gin.Context) {
	ctx := c.Request.Context()
	au := c.MustGet(httpCommon.AUTH_USER).(uModel.AuthUser)

	id := c.Param("id")

	imageUrl, err := bucket.UploadBucketImage(c, "users-pict", id)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, httpCommon.Error{
			Message: "No Such File",
		})
		return
	}

	updateUserPic, err := d.userRepo.UpdateUserImage(ctx, uModel.UpdateUser{
		ID:       id,
		ImageUrl: imageUrl,
	}, au)

	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, httpCommon.Response{
		Data:    updateUserPic,
		Message: "Update User picture Succesfuly",
	})

}
