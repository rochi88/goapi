package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model // Adds some metadata fields to the table
	Name       string
	Email      string
	Phone      string
	Password   string
	Status     bool
}
