package models

import (
	"errors"

	"github.com/asaskevich/govalidator"
	"github.com/jinzhu/gorm"
)

type Account struct {
	gorm.Model
	Username       string `valid:"required,length(4|15)" schema:username`
	Password       string `gorm:"-" valid:"required,length(5|15)" schema:"password"`
	Email          string `valid:"required,email" schema:"email"`
	VerifyPass     string `gorm:"-" schema:"verifypass"`
	HashedPassword string
}

func (u *Account) Validate() error {
	_, err := govalidator.ValidateStruct(u)
	if err != nil {
		return err
	}
	if u.Password != u.VerifyPass {
		return errors.New("password missmatch")
	}
	return err
}
