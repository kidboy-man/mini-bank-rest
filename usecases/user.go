package usecases

import (
	"context"

	"github.com/kidboy-man/mini-bank-rest/models"
	"github.com/kidboy-man/mini-bank-rest/repositories"
	"gorm.io/gorm"
)

type UserUsecase interface {
	Create(ctx context.Context, user *models.User) (err error)
	Delete(ctx context.Context, userID uint) (err error)
	GetByUsername(ctx context.Context, username string) (user *models.User, err error)
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

func (u *userUsecase) Register(ctx context.Context) (err error) {
	// TODO
	return
}
