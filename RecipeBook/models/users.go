package models

import (
	"recipebook/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Users struct {
    GormModel
    Username         string    `gorm:"not null;uniqueIndex" json:"username" form:"username" valid:"required~Your username is required"`
    Email            string    `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required~Your email is required,email~Invalid email format"`
    Password         string    `gorm:"not null" json:"password" form:"password" valid:"required~Your password is required,minstringlength(6)~Password has to have a minimum length of 6 characters"`
    Bio              string    `gorm:"not null" json:"bio" form:"bio"  valid:"required~Your bio is required"`
    Recipes          []Recipes  `gorm:"foreignKey:UserID" json:"recipes"`
}

func (u *Users) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	u.Password = helpers.HashPass(u.Password)
	err = nil
	return
}