package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model // Adds some metadata fields to the table
	Name       string
	Email      string `json:"username" gorm:"unique;" validate:"required,email,min=6,max=32"`
	Password   string `json:"-" gorm:"type:text;" validate:"required,min=4"`
	Phone      string
	Status     bool
}
