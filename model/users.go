package model

import "time"
import _ "github.com/go-playground/validator/v10"

type Users struct {
	ID          uint   `gorm:"primary_key"`
	Name        string `validate:"required"`
	Birthday    string `validate:"required"`
	Phone       string `validate:"required"`
	Email       string `validate:"required,email"`
	Photo       string `validate:"required"`
	Appointment string `validate:"required"`
	Status      int    `gorm:"default:0"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
