package http

import (
	"firebase.google.com/go/v4/auth"
	"fmt"
	errorCommon "github.com/ChoTracker-C23-PS308/ChoTracker-CC/common/error"
	uModel "github.com/ChoTracker-C23-PS308/ChoTracker-CC/internal/model/user"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Auth(fAuth *auth.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		a := c.Request.Header.Get("Authorization")
		if len(a) <= BEARER {
			c.Error(errorCommon.NewInvariantError("authorization header not valid"))
			c.Abort()
			return
		}
		idToken := a[BEARER:]

		// verify token is real from user
		token, err := fAuth.VerifyIDToken(ctx, idToken)
		if err != nil {
			c.Error(errorCommon.NewUnauthorizedError(fmt.Sprintf("cannot authorize token: %s", err.Error())))
			c.Abort()
			return
		}

		user := uModel.AuthUser{
			ID:   token.UID,
			Role: uModel.DEFAULT,
		}
		// set role if claim exists
		if r, ok := token.Claims["role"]; ok {
			if rs, ok := r.(string); ok {
				user.SetRoleString(rs)
			}
		}

		c.Set(AUTH_USER, user)
		c.Next()
	}
}

func CORS() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	})
}
