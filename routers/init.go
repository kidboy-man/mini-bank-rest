package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/kidboy-man/mini-bank-rest/controllers"
	"github.com/kidboy-man/mini-bank-rest/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Setup() (r *gin.Engine) {
	r = gin.Default()
	r.ForwardedByClientIP = true
	docs.SwaggerInfo.BasePath = ""

	upc := controllers.UserPublicController{}
	upc.Prepare()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.GET("/users/:username", upc.GetUser)
	return
}
