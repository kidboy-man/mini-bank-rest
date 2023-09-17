package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kidboy-man/mini-bank-rest/configs"
	"github.com/kidboy-man/mini-bank-rest/databases"
)

func main() {
	err := configs.AppConfig.Setup()
	if err != nil {
		panic(err)
	}

	err = databases.Init()
	if err != nil {
		panic(err)
	}

	r := gin.Default()
	r.ForwardedByClientIP = true
	r.SetTrustedProxies([]string{"127.0.0.1"})
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run(fmt.Sprintf("localhost:%s", configs.AppConfig.AppPort))
}
