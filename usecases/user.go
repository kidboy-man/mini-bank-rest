package usecases

import (
	"context"
	"net/http"
	"strings"

	"github.com/kidboy-man/mini-bank-rest/constants"
	"github.com/kidboy-man/mini-bank-rest/helpers"
	"github.com/kidboy-man/mini-bank-rest/models"
	"github.com/kidboy-man/mini-bank-rest/repositories"
	"github.com/kidboy-man/mini-bank-rest/schemas"
	"github.com/kidboy-man/mini-bank-rest/utils"
	"gorm.io/gorm"
)

type UserUsecase interface {
	Create(ctx context.Context, user *models.User) (err error)
	Delete(ctx context.Context, userID uint) (err error)
	GetByUsername(ctx context.Context, username string) (user *models.User, err error)
	Login(ctx context.Context, param schemas.Login) (user *models.User, err error)
	Register(ctx context.Context, param schemas.Register) (err error)
	Update(ctx context.Context, user *models.User) (err error)
}

type userUsecase struct {
	db       *gorm.DB
	userRepo repositories.UserRepo
}

func NewUserUsecase(db *gorm.DB) UserUsecase {
	userRepo := repositories.NewUserRepo(db)
	return &userUsecase{
		userRepo: userRepo,
		db:       db,
	}
}

func (u *userUsecase) GetByUsername(ctx context.Context, username string) (user *models.User, err error) {
	user, err = u.userRepo.GetByUsername(ctx, username)
	return
}

func (u *userUsecase) Create(ctx context.Context, user *models.User) (err error) {
	err = u.userRepo.Create(ctx, u.db, user)
	return
}

func (u *userUsecase) Update(ctx context.Context, user *models.User) (err error) {
	err = u.userRepo.Update(ctx, u.db, user)
	return
}

func (u *userUsecase) Delete(ctx context.Context, userID uint) (err error) {
	err = u.userRepo.Delete(ctx, u.db, userID)
	return
}

func (u *userUsecase) Register(ctx context.Context, param schemas.Register) (err error) {
	param.Prepare()
	hashedPassword, err := helpers.HashPassword(param.Password)
	if err != nil {
		return
	}

	err = u.userRepo.Create(ctx, u.db, &models.User{
		Username: param.Username,
		Email:    param.Email,
		Password: hashedPassword,
	})
	if err != nil && strings.Contains(strings.ToLower(err.Error()), "duplicate") {
		if strings.Contains(strings.ToLower(err.Error()), "email") {
			err = &schemas.CustomError{
				Code:       constants.RegisterEmailNotAvailableErrCode,
				HTTPStatus: http.StatusBadRequest,
				Message:    "email address is already taken",
			}
		}

		if strings.Contains(strings.ToLower(err.Error()), "username") {
			err = &schemas.CustomError{
				Code:       constants.RegisterUsernameNotAvailableErrCode,
				HTTPStatus: http.StatusBadRequest,
				Message:    "username is already taken",
			}
		}
	}
	return
}

func (u *userUsecase) Login(ctx context.Context, param schemas.Login) (user *models.User, err error) {
	if utils.IsEmail(param.Identifier) {
		user, err = u.userRepo.GetByEmail(ctx, param.Identifier)
		if err != nil {
			if err.Error() == gorm.ErrRecordNotFound.Error() {
				err = &schemas.CustomError{
					Code:       constants.LoginEmailNotFoundErrCode,
					HTTPStatus: http.StatusBadRequest,
					Message:    "email address is not registered",
				}
			}
			return
		}
	} else {
		user, err = u.userRepo.GetByUsername(ctx, param.Identifier)
		if err != nil {
			if err.Error() == gorm.ErrRecordNotFound.Error() {
				err = &schemas.CustomError{
					Code:       constants.LoginUsernameNotFoundErrCode,
					HTTPStatus: http.StatusBadRequest,
					Message:    "username is not registered",
				}
			}
			return
		}
	}

	if !helpers.CheckPasswordHash(param.Password, user.Password) {
		err = &schemas.CustomError{
			Code:       constants.LoginInvalidPasswordErrCode,
			HTTPStatus: http.StatusBadRequest,
			Message:    "invalid password",
		}
		return
	}

	// TODO: set token
	return
}
