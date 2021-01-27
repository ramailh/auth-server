package router

import (
	"github.com/ramailh/auth-server/http/controller"

	"github.com/gin-gonic/gin"
)

func NewHTTPRouter() *gin.Engine {
	rtr := gin.Default()

	rtr.LoadHTMLGlob("template/*")

	auth := rtr.Group("/auth")
	{
		oauth := auth.Group("/oauth")
		{
			oauth.GET("/google", controller.Google)
			oauth.GET("/google-callback", controller.GoogleCallback)
		}

		auth.POST("/login", controller.LoginPost)
		auth.GET("/login", controller.Login)

	}

	user := rtr.Group("/user")
	{
		user.GET("/", controller.GetUser)
		user.GET("/:id", controller.GetUserByID)
		user.POST("/insert", controller.InsertUser)
		user.PUT("/:id", controller.UpdateUser)
		user.DELETE("/:id", controller.DeleteUser)
	}

	return rtr
}
