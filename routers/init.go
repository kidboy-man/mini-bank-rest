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
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	v1 := r.Group("/v1")

	publicUserRoute := v1.Group("/public/users")
	upc := controllers.UserPublicController{}
	upc.Prepare()
	publicUserRoute.GET("/:username", upc.GetUser)

	publicAuthRoute := v1.Group("/public/auth")
	apc := controllers.AuthPublicController{}
	apc.Prepare()
	publicAuthRoute.POST("/register", apc.Register)
	return
}
