package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Admin struct {
	Id        uint      `json:"id"`
	Username  string    `json:"username" validate:"required" gorm:"not null"`
	Email     string    `gorm:"unique" validate:"required,email"`
	Password  string    `json:"password" gorm:"not null" validate:"required"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (admin *Admin) ValidateAdmin() error {
	validate := validator.New()
	return validate.Struct(admin)
}
