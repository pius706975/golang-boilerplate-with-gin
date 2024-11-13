package user

import (
	"go-gin/middlewares"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine, controller *userController, prefix string) {
	userGroup := router.Group(prefix + "/user")
	{
		userGroup.GET("/", func(ctx *gin.Context) {
			controller.GetUsers(ctx)
		})

		userGroup.GET("/:id", middlewares.AuthMiddleware(), func(ctx *gin.Context) {
			controller.GetUserById(ctx)
		})

		userGroup.GET("/profile", middlewares.AuthMiddleware(), func(ctx *gin.Context) {
			controller.GetProfile(ctx)
		})
	}
}
