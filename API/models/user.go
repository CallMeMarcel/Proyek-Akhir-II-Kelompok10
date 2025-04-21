package models

import (
	"regexp"
	"time"

	"github.com/go-playground/validator/v10"
)

type User struct {
	Id        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" validate:"required"`
	Username  string    `json:"username" validate:"required"`
	Email     string    `json:"email" validate:"required,email"`
	Phone     string    `json:"phone" validate:"required"`
	Password  string    `json:"password" validate:"required,min=6"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

// CustomValidator digunakan untuk validasi tambahan di luar tag
func (user *User) ValidateUser() error {
	validate := validator.New()

	// Custom regex untuk Phone: hanya angka 10 sampai 15 digit
	validate.RegisterValidation("phone_regex", func(fl validator.FieldLevel) bool {
		re := regexp.MustCompile(`^\d{10,15}$`)
		return re.MatchString(fl.Field().String())
	})

	// Daftarkan validasi khusus phone_regex
	validate.RegisterStructValidation(func(sl validator.StructLevel) {
		u := sl.Current().Interface().(User)
		if !regexp.MustCompile(`^\d{10,15}$`).MatchString(u.Phone) {
			sl.ReportError(u.Phone, "Phone", "phone", "phone_regex", "")
		}
	}, User{})

	return validate.Struct(user)
}
