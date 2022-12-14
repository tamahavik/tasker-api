package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model `json:"-"`
	ID         uint `json:"id" gorm:"primarykey"`
	Username   string
	Email      string
	Password   string `json:"-"`
	Name       string
}
