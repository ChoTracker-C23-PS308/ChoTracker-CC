package http

import (
	"firebase.google.com/go/v4/auth"
	httpCommon "github.com/ChoTracker-C23-PS308/ChoTracker-CC/common/http"
	aRepo "github.com/ChoTracker-C23-PS308/ChoTracker-CC/internal/repository/article"
	"github.com/gin-gonic/gin"
)

type HTTPArtikelDelivery struct {
	articleRepo aRepo.Repository
}

func NewHTTPArticleDelivery(g *gin.RouterGroup, articleRepo aRepo.Repository, fauth *auth.Client) HTTPArtikelDelivery {
	h := HTTPArtikelDelivery{articleRepo: articleRepo}

	g.GET("/articles/:id", httpCommon.Auth(fauth), h.getArticle)
	g.POST("/articles", httpCommon.Auth(fauth), h.addArticle)
	g.PUT("/articles/:id", httpCommon.Auth(fauth), h.updateArticle)
	g.DELETE("/articles/:id", httpCommon.Auth(fauth), h.deleteArticle)

	return h
}
