package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kidboy-man/mini-bank-rest/configs"
	"github.com/kidboy-man/mini-bank-rest/schemas"
	"github.com/kidboy-man/mini-bank-rest/usecases"
)

type AuthPublicController struct {
	Controllers
	userUsecase usecases.UserUsecase
}

func (c *AuthPublicController) Prepare() {
	c.userUsecase = usecases.NewUserUsecase(configs.AppConfig.DBConnection)
}

// Register      godoc
// @Summary      Register
// @Description  Register user to our system account.
// @Tags         authentication
// @Accept 		 json
// @Param 		 registrationData body schemas.Register true "user registration data"
// @Produce      json
// @Success      201  {object}  schemas.GeneralResponse
// @Router       /v1/public/auth/register [post]
func (apc *AuthPublicController) Register(c *gin.Context) {
	registration := schemas.Register{}
	err := c.BindJSON(&registration)
	if err != nil {
		apc.ReturnNotOK(c, err)
		return
	}

	apc.ReturnOK(c, http.StatusCreated, "registration success", nil)
}
