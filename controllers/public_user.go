package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetUser       godoc
// @Summary      Get One User
// @Description  Returns the user who matches the username.
// @Tags         users
// @Produce      json
// @Success      200  {object}  models.User
// @Router       /users [get]
func GetUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
