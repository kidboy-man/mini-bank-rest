package models

import (
	"net/http"
	"strings"
	"time"

	"github.com/kidboy-man/mini-bank-rest/constants"
	"github.com/kidboy-man/mini-bank-rest/schemas"
	"gorm.io/gorm"
)

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Username  string    `gorm:"index;unique;type:varchar(255)" validate:"required" json:"username"`
	Password  string    `gorm:"type:varchar(255)" validate:"required" json:"-"`
	Email     string    `gorm:"index;unique;type:varchar(255)" validate:"required,email" json:"email"`
	CreatedAt time.Time `gorm:"autoCreateTime;<-:create" json:"createdAt"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updatedAt"`

	Token string `gorm:"-" json:"token,omitempty"`
}

func (User) TableName() string {
	return "users"
}

func (u *User) setAttr() {
	u.Username = strings.ToLower(strings.TrimSpace(u.Username))
	u.Email = strings.ToLower(strings.TrimSpace(u.Email))
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	// should be handled at max in usecase,
	// so, if this passes through here, we are lacking of validations
	if u.Username == "" || u.Password == "" || u.Email == "" {
		err = &schemas.CustomError{
			Code:       constants.OrmHookValidationErrCode,
			HTTPStatus: http.StatusInternalServerError,
			Message:    "Incomplete user data, missing required values",
		}
		return
	}
	u.setAttr()
	return
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	u.setAttr()
	return
}
