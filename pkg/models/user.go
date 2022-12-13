package models

import "time"

type User struct {
	Id        string     `json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	Username  string     `json:"username" binding:"required"`
	Email     string     `json:"email" binding:"required"`
	Password  string     `json:"password" binding:"required"`
	Name      string     `json:"name" binding:"required"`
	Age       int        `json:"age" binding:"required"`
}
