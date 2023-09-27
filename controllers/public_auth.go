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

	err = apc.userUsecase.Register(c.Request.Context(), registration)
	if err != nil {
		apc.ReturnNotOK(c, err)
		return
	}

	apc.ReturnOK(c, http.StatusCreated, "registration success", nil)
}

// Login      	 godoc
// @Summary      Login
// @Description  Login user to our system account.
// @Tags         authentication
// @Accept 		 json
// @Param 		 credential body schemas.Login true "user credential"
// @Produce      json
// @Success      200  {object}  schemas.GeneralResponse
// @Router       /v1/public/auth/login [post]
func (apc *AuthPublicController) Login(c *gin.Context) {
	credential := schemas.Login{}
	err := c.BindJSON(&credential)
	if err != nil {
		apc.ReturnNotOK(c, err)
		return
	}

	user, err := apc.userUsecase.Login(c.Request.Context(), credential)
	if err != nil {
		apc.ReturnNotOK(c, err)
		return
	}

	apc.ReturnOK(c, http.StatusCreated, "login successful", user)
}
