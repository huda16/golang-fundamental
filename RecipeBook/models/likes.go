package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Likes struct {
    GormModel
    UserID    uint      `gorm:"not null" json:"user_id" form:"user_id" valid:"required~User Id of your like is required"`
    RecipeID  uint      `gorm:"not null" json:"recipe_id" form:"recipe_id" valid:"required~Recipe Id of your like is required"`
    User      *Users     
    Recipe    *Recipes    
}

func (p *Likes) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

func (p *Likes) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}