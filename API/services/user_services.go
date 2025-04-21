package services


import (
	"api/dto"
	"api/models"
	"errors"
)

// UserService bertanggung jawab untuk mengelola logika bisnis user
type UserService struct {}

// NewUserService menginisialisasi dan mengembalikan instance UserService
func NewUserService() *UserService {
	return &UserService{}
}

// CreateUser bertanggung jawab untuk membuat user baru
func (s *UserService) CreateUser(input dto.RequestUserRegister) error {
	// Lakukan validasi atau logika lainnya jika perlu

	// Simulasi penyimpanan user ke database
	user := models.User{
		Name:     input.Name,
		Username: input.Username,
		Email:    input.Email,
		Phone:    input.Phone,
		Password: input.Password, // Jangan lupa mengenkripsi password sebelum menyimpannya!
	}

	// Simulasikan penyimpanan ke database
	err := saveUserToDatabase(user)
	if err != nil {
		return errors.New("failed to create user: " + err.Error())
	}

	return nil
}

// Simulasi penyimpanan user ke database
func saveUserToDatabase(user models.User) error {
	// Implementasikan penyimpanan ke database, misalnya:
	// db.Create(&user) 

	// Kembalikan nil untuk simulasikan berhasil
	return nil
}
