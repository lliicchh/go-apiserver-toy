package middleware

import "github.com/gin-gonic/gin"

func Nocache(c *gin.Context)  {
	c.Header("Cache-Control","no-cache, nostore, max-age=0, must-revalidate, value")
	//...
	c.Next()
}
