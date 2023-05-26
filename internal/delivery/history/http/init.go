package http

import (
	"firebase.google.com/go/v4/auth"
	httpCommon "github.com/ChoTracker-C23-PS308/ChoTracker-CC/common/http"
	hRepo "github.com/ChoTracker-C23-PS308/ChoTracker-CC/internal/repository/history"
	"github.com/gin-gonic/gin"
)

type HTTPHistoryDelivery struct {
	historyRepo hRepo.Repository
}

func NewHTTPHistoryDelivery(g *gin.RouterGroup, historyRepo hRepo.Repository, fauth *auth.Client) HTTPHistoryDelivery {
	h := HTTPHistoryDelivery{historyRepo: historyRepo}

	g.POST("/history/:id", httpCommon.Auth(fauth), h.addHistory)
	g.DELETE("/history/:uid/:id", httpCommon.Auth(fauth), h.deleteHistory)
	g.GET("/history/:uid", httpCommon.Auth(fauth), h.getHistories)
	return h

}
