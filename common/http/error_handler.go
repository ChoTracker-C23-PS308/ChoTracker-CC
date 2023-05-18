package http

import (
	"errors"
	errorCommon "github.com/ChoTracker-C23-PS308/ChoTracker-CC/common/error"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func (h HTTPServer) errorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		errs := c.Errors.ByType(gin.ErrorTypeAny)
		if len(errs) > 0 {
			err := errs[0]
			if !err.IsType(gin.ErrorTypePrivate) {
				var ves validator.ValidationErrors
				errors.As(err, &ves)
				keys := make(map[string]string)
				for _, ve := range ves {
					keys[ve.Field()] = ve.Tag()
				}
				//c.Writer.Header().Set("Content-Type", "application/json")
				render.JSON{Data: Error{
					Message: err.Error(),
					Errors:  keys,
				}}.Render(c.Writer)

				return
			}
			switch err := err.Err.(type) {
			case *errorCommon.ClientError:
				c.JSON(err.StatusCode, Error{
					Message: err.Message,
				})
			default:
				if he, ok := errorCommon.DomainErrorTranslatorDirectories[err.Error()]; ok {
					c.JSON(he.StatusCode, Error{
						Message: he.Message,
					})
				} else {
					c.JSON(http.StatusInternalServerError, Error{
						Message: "Internal server error",
					})
				}
			}
		}
	}
}
