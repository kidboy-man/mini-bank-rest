package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
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
	message := err.Error()
	code := "5000005"
	httpStatus := http.StatusInternalServerError
	if v, ok := err.(*schemas.CustomError); ok {
		message = v.Error()
		code = fmt.Sprintf("%d%d", v.HTTPStatus, v.Code)
		httpStatus = v.HTTPStatus

	}
	ctx.JSON(httpStatus, gin.H{
		"success": false,
		"code":    code,
		"message": message,
	})
}
