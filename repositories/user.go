package repositories

import (
	"context"
	"net/http"
	"strings"

	"github.com/kidboy-man/mini-bank-rest/constants"
	"github.com/kidboy-man/mini-bank-rest/models"
	"github.com/kidboy-man/mini-bank-rest/schemas"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepo interface {
	Create(ctx context.Context, db *gorm.DB, user *models.User) (err error)
	Delete(ctx context.Context, db *gorm.DB, userID uint) (err error)
	GetByEmail(ctx context.Context, email string) (user *models.User, err error)
	GetByUsername(ctx context.Context, username string) (user *models.User, err error)
	Update(ctx context.Context, db *gorm.DB, user *models.User) (err error)
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &userRepo{db: db}
}

func (r *userRepo) Create(ctx context.Context, db *gorm.DB, user *models.User) (err error) {
	err = db.Omit(clause.Associations).Model(user).Create(user).Error
	if err != nil {
		err = &schemas.CustomError{
			Code:       constants.InternalServerErrCode,
			HTTPStatus: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}
	return
}

func (r *userRepo) Update(ctx context.Context, db *gorm.DB, user *models.User) (err error) {
	row := db.Omit(clause.Associations).Model(user).Updates(user)
	err = row.Error
	if err != nil {
		err = &schemas.CustomError{
			Code:       constants.InternalServerErrCode,
			HTTPStatus: http.StatusInternalServerError,
			Message:    err.Error(),
		}
		return
	}

	if row.RowsAffected == 0 {
		err = &schemas.CustomError{
			Code:       constants.QueryNotFoundErrCode,
			HTTPStatus: http.StatusNotFound,
			Message:    gorm.ErrRecordNotFound.Error(),
		}
		return
	}
	return
}

func (r *userRepo) Delete(ctx context.Context, db *gorm.DB, userID uint) (err error) {
	user := &models.User{ID: userID}
	row := db.Omit(clause.Associations).Model(user).Delete(user)
	err = row.Error
	if err != nil {
		err = &schemas.CustomError{
			Code:       constants.InternalServerErrCode,
			HTTPStatus: http.StatusInternalServerError,
			Message:    err.Error(),
		}
		return
	}

	if row.RowsAffected == 0 {
		err = &schemas.CustomError{
			Code:       constants.QueryNotFoundErrCode,
			HTTPStatus: http.StatusNotFound,
			Message:    gorm.ErrRecordNotFound.Error(),
		}
		return
	}
	return
}

func (r *userRepo) GetByUsername(ctx context.Context, username string) (user *models.User, err error) {
	qs := r.db.Where("username ILIKE ?", username)
	err = qs.First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = &schemas.CustomError{
				Code:       constants.QueryNotFoundErrCode,
				HTTPStatus: http.StatusNotFound,
				Message:    err.Error(),
			}
			return nil, err
		}
		err = &schemas.CustomError{
			Code:       constants.InternalServerErrCode,
			HTTPStatus: http.StatusInternalServerError,
			Message:    err.Error(),
		}
		return nil, err
	}
	return
}

func (r *userRepo) GetByEmail(ctx context.Context, email string) (user *models.User, err error) {
	qs := r.db.Where("email = ?", strings.ToLower(strings.TrimSpace(email)))
	err = qs.First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = &schemas.CustomError{
				Code:       constants.QueryNotFoundErrCode,
				HTTPStatus: http.StatusNotFound,
				Message:    err.Error(),
			}
			return nil, err
		}
		err = &schemas.CustomError{
			Code:       constants.InternalServerErrCode,
			HTTPStatus: http.StatusInternalServerError,
			Message:    err.Error(),
		}
		return nil, err
	}
	return
}
