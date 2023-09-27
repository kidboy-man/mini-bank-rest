package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/kidboy-man/mini-bank-rest/constants"
	"github.com/kidboy-man/mini-bank-rest/schemas"
)

type Controllers struct{}

func (c *Controllers) ReturnOK(ctx *gin.Context, httpStatus int, message string, object interface{}) {
	if message == "" {
		message = "Successful"
	}
	ctx.JSON(httpStatus, gin.H{
		"success": true,
		"code":    fmt.Sprintf("%d0001", httpStatus),
		"message": message,
		"data":    object,
	})
}

func (c *Controllers) ReturnNotOK(ctx *gin.Context, err error) {
	var message string
	var code string
	var httpStatus int

	switch err := err.(type) {
	case *schemas.CustomError:
		message = err.Error()
		httpStatus = err.HTTPStatus
		code = fmt.Sprintf("%d%d", httpStatus, err.Code)

	case validator.ValidationErrors:
		message = err.Error()
		httpStatus = http.StatusBadRequest
		code = fmt.Sprintf("%d%d", httpStatus, constants.BadRequestErrCode)

	default:
		message = err.Error()
		httpStatus = http.StatusInternalServerError
		code = fmt.Sprintf("%d%d", httpStatus, constants.InternalServerErrCode)
	}

	ctx.JSON(httpStatus, gin.H{
		"success": false,
		"code":    code,
		"message": message,
	})
}
