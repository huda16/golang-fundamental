package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Recipes struct {
    GormModel
    Title       string     `gorm:"not null" json:"title" form:"title" valid:"required~Title of your recipe is required"`
    Description string     `gorm:"not null" json:"description" form:"description" valid:"required~Description of your recipe is required"`
    Ingredients string     `gorm:"not null" json:"ingredients" form:"ingredients" valid:"required~Ingredients of your recipe is required"`
    Steps       string     `gorm:"not null" json:"steps" form:"steps" valid:"required~Steps of your recipe is required"`
	PictureUrl string `gorm:"not null" json:"picture_url" form:"picture_url" valid:"required~URL of your picture is required"`
    Category    *string    `json:"category" form:"category"`
	Tags        *string    `json:"tags" form:"tags"`
	UserID      uint       `json:"user_id"`
	User        *Users     `gorm:"foreignKey:UserID"`
	Comments    []Comments `gorm:"foreignKey:RecipeID"`
	Likes       []Likes    `gorm:"foreignKey:RecipeID"`
}

func (p *Recipes) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

func (p *Recipes) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}