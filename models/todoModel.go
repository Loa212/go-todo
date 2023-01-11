package models

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	UserID uint `json:"UserId"`
	Title string 
	Done bool
}