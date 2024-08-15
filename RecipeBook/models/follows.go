package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Follows struct {
    GormModel
    FollowerID uint      `gorm:"not null" json:"follower_id" form:"follower_id" valid:"required~follower Id is required"`
    FollowedID uint      `gorm:"not null" json:"followed_id" form:"followed_id" valid:"required~followed Id is required"`
    Follower   *Users     
    Followed   *Users     
}

func (p *Follows) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

func (p *Follows) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}