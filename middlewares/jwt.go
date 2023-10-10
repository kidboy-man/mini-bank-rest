package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/kidboy-man/mini-bank-rest/configs"
	"github.com/kidboy-man/mini-bank-rest/constants"
	"github.com/kidboy-man/mini-bank-rest/schemas"
)

type JWTClaims struct {
	jwt.RegisteredClaims
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
}

type Auth struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
}

func VerifyToken(ctx *gin.Context) {
	auth, err := doVerifyToken(ctx.Request)
	if err != nil {
		errAuth(ctx, err)
		return
	}

	ctx.Set("user_id", auth.UserID)
	ctx.Set("username", auth.Username)
}

func doVerifyToken(r *http.Request) (result *Auth, err error) {

	token, err := getToken(r)
	if err != nil {
		err = &schemas.CustomError{
			Code:       constants.InternalServerErrCode,
			HTTPStatus: http.StatusInternalServerError,
			Message:    err.Error(),
		}
		return
	}

	isVerified, claims, err := parseTokenJWT(token)
	if err != nil {
		err = &schemas.CustomError{
			Code:       constants.InternalServerErrCode,
			HTTPStatus: http.StatusInternalServerError,
			Message:    err.Error(),
		}
		return
	}

	if !isVerified {
		err = &schemas.CustomError{
			Code:       constants.NotAuthorizedErrCode,
			HTTPStatus: http.StatusUnauthorized,
			Message:    "INVALID_TOKEN",
		}
		return
	}

	result = &Auth{
		UserID:   claims.UserID,
		Username: claims.Username,
	}
	return

}

func getToken(r *http.Request) (token string, err error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		err = &schemas.CustomError{
			Code:       constants.NotAuthorizedErrCode,
			HTTPStatus: http.StatusUnauthorized,
			Message:    "EMPTY_TOKEN",
		}
		return
	}

	s := strings.Split(authHeader, " ")
	if len(s) != 2 {
		err = &schemas.CustomError{
			Code:       constants.NotAuthorizedErrCode,
			HTTPStatus: http.StatusUnauthorized,
			Message:    "INVALID_TOKEN",
		}
		return
	}

	token = s[1]
	return
}

func parseTokenJWT(token string) (isVerified bool, result *JWTClaims, err error) {
	result = &JWTClaims{}
	jwtClaims, err := jwt.ParseWithClaims(token, result, func(token *jwt.Token) (interface{}, error) {
		return configs.AppConfig.JWTSignatureKey, nil
	})

	if result == nil || jwtClaims == nil || !jwtClaims.Valid || err != nil {
		return
	}
	isVerified = true
	return
}

func errAuth(ctx *gin.Context, err error) {
	var message string
	var httpStatus int
	var code string
	switch err := err.(type) {
	case *schemas.CustomError:
		message = err.Error()
		httpStatus = err.HTTPStatus
		code = fmt.Sprintf("%d%d", httpStatus, err.Code)

	default:
		message = err.Error()
		httpStatus = http.StatusInternalServerError
		code = fmt.Sprintf("%d%d", httpStatus, constants.InternalServerErrCode)
	}

	ctx.AbortWithStatusJSON(httpStatus, gin.H{
		"success": false,
		"code":    code,
		"message": message,
	})
}
