package database

import "gorm.io/gorm"

type UserEntity struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email"`
	UserId   string `json:"user_id" gorm:"primaryKey"`
	Password string `json:"password,omitempty"`
}
