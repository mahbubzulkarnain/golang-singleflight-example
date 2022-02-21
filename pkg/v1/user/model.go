package user

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Entity
}

type Users []*User

func (m User) TableName() string {
	return "users"
}
