package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/lliicchh/apiserver/handler"
	"github.com/lliicchh/apiserver/pkg/errno"
	"github.com/lliicchh/apiserver/pkg/token"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, err := token.ParseRequest(c); err != nil {
			handler.SendResponse(c, errno.ErrTokenInvalid, nil)
			c.Abort()
			return
		}

		c.Next()
	}
}
