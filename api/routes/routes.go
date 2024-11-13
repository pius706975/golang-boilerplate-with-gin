package routes

import (
	_ "go-gin/docs"
	"go-gin/modules/auth"
	"go-gin/modules/user"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

const (
	APIPrefix = "/api"
)

func homeHandler(ctx *gin.Context) {
	ctx.JSON(404, gin.H{
		"status":  404,
		"message": "Sorry! Page not found",
	})
}

func RouteApp(router *gin.Engine, db *gorm.DB) error {
	router.GET(APIPrefix+"/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET(APIPrefix, homeHandler)

	auth.AuthRoutesModule(router, db, APIPrefix)
	user.UserRoutesModule(router, db, APIPrefix)

	return nil
}