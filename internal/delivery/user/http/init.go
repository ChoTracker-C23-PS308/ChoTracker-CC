package http

import (
	"firebase.google.com/go/v4/auth"
	httpCommon "github.com/ChoTracker-C23-PS308/ChoTracker-CC/common/http"

	uRepo "github.com/ChoTracker-C23-PS308/ChoTracker-CC/internal/repository/user"
	"github.com/gin-gonic/gin"
)

type HTTPUserDelivery struct {
	userRepo uRepo.Repository
}

func NewHTTPUserDelivery(g *gin.RouterGroup, userRepo uRepo.Repository, fauth *auth.Client) HTTPUserDelivery {
	h := HTTPUserDelivery{userRepo: userRepo}

	g.GET("/users/:id", httpCommon.Auth(fauth), h.getUser)
	g.POST("/users", httpCommon.Auth(fauth), h.addUser)
	g.PUT("/users/:id", httpCommon.Auth(fauth), h.updateUser)

	return h
}