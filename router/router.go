package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lliicchh/apiserver/handler/sd"
	"github.com/lliicchh/apiserver/handler/user"
	"github.com/lliicchh/apiserver/router/middleware"
	"net/http"
)

func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	var (
		u, svcd *gin.RouterGroup
	)

	g.Use(gin.Recovery())
	g.Use(mw...)

	// api for authentication
	g.POST("/login", user.Login)


	// 404
	g.NoRoute(func(context *gin.Context) {
		context.String(http.StatusNotFound, "the incorrect api route")
	})

	u = g.Group("/v1/user")
	u.Use(middleware.AuthMiddleware())
	{
		u.POST("", user.Create)
		u.DELETE("/:id", user.Delete)
		u.PUT("/:id", user.Update)
		u.GET("", user.List)
		u.GET("/:username", user.Get)

	}

	svcd = g.Group("/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
		svcd.GET("/disk", sd.DiskCheck)
		svcd.GET("/cpu", sd.CPUCheck)
	}

	return g
}
