package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/xws117/bluebell/controllder"
)

func SetupRouter() *gin.Engine {
	server := gin.Default()

	v1 := server.Group("/v1")

	v1.GET("/login", controllder.Loginfunc())
	v1.GET("/signup", controllder.Signupfunc())
	v1.GET("/refresh_token", controllder.Refreshfunc())

	v1.Use(controllder.JWTAUTH())
	{
		v1.GET("/post", func(context *gin.Context) {

		})
	}

	return server
}
