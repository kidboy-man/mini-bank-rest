package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kidboy-man/mini-bank-rest/configs"
	"github.com/kidboy-man/mini-bank-rest/usecases"
)

type UserPublicController struct {
	Controllers
	userUsecase usecases.UserUsecase
}

func (c *UserPublicController) Prepare() {
	c.userUsecase = usecases.NewUserUsecase(configs.AppConfig.DBConnection)
}

// GetUser       godoc
// @Summary      Get One User
// @Description  Returns the user who matches the username.
// @Tags         users
// @Produce      json
// @Success      200  {object}  models.User
// @Router       /users/{username} [get]
func (upc *UserPublicController) GetUser(c *gin.Context) {
	username := c.Param("username")
	user, err := upc.userUsecase.GetByUsername(c.Request.Context(), username)
	if err != nil {
		upc.ReturnNotOK(c, err)
		return
	}

	upc.ReturnOK(c, http.StatusOK, "", user)
}
