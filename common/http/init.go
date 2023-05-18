package http

import "github.com/gin-gonic/gin"

type HTTPServer struct {
	Router *gin.Engine
}

func NewHTTPServer() HTTPServer {
	e := gin.New()
	h := HTTPServer{
		Router: e,
	}

	e.Use(gin.Recovery())
	e.Use(h.errorHandler())

	return h
}
