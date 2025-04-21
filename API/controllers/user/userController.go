package controllers

import (
	"api/database"
	"api/models"
	"api/services"
	"net/http"
	"regexp"
	"strconv"
	"time"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

var UserSecretKey = "USER_SECRET_KEY"

type UserController struct {
	userService *services.UserService
}

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{userService: userService}
}

// RegisterUser handles user registration
func (uc *UserController) RegisterUser(ctx *fiber.Ctx) error {
	var data struct {
		Name     string `json:"name"`
		Username string `json:"username"`
		Email    string `json:"email"`
		PhoneNo  string `json:"phone_no"`
		Password string `json:"password"`
	}

	// Parse request body
	if err := ctx.BodyParser(&data); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Format request tidak valid",
			"error":   err.Error(),
		})
	}

	// Validasi input
	if data.Name == "" || data.Username == "" || data.Email == "" || data.PhoneNo == "" || data.Password == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Semua field harus diisi",
		})
	}

	// Validasi email format
	if !regexp.MustCompile(`^[a-zA-Z0-9+_.-]+@[a-zA-Z0-9.-]+$`).MatchString(data.Email) {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Format email tidak valid",
		})
	}

	// Cek apakah email sudah digunakan
	var existingUser models.User
	if err := database.DB.Where("email = ?", data.Email).First(&existingUser).Error; err == nil {
		return ctx.Status(fiber.StatusConflict).JSON(fiber.Map{
			"message": "Email sudah terdaftar",
		})
	}

	// nomor telepon
	if !regexp.MustCompile(`^\d{10,15}$`).MatchString(data.PhoneNo) {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Format nomor telepon tidak valid",
		})
	}
	

	// Enkripsi password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal mengenkripsi password",
		})
	}

	newUser := models.User{
		Name:      data.Name,
		Username:  data.Username,
		Email:     data.Email,
		Phone:     data.PhoneNo,
		Password:  string(hashedPassword),
	}


	if err := database.DB.Create(&newUser).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal menyimpan data pengguna",
			"error":   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Pendaftaran berhasil",
		"email":   newUser.Email,
	})
}

func (uc *UserController) LoginUser(ctx *fiber.Ctx) error {
	var data map[string]string
	if err := ctx.BodyParser(&data); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Format request tidak valid",
		})
	}

	// Debug log untuk memastikan input diterima dengan benar
	fmt.Println("Input email:", data["email"])
	fmt.Println("Password input:", data["password"])

	var user models.User
	database.DB.Where("email = ?", data["email"]).First(&user)

	if user.Email != data["email"] {
		ctx.Status(fiber.StatusNotFound)
		return ctx.JSON(fiber.Map{
			"message": "user not found",
		})
	}

	// Debug log untuk memastikan password dari database (terenkripsi)
	fmt.Println("Password dari database (hashed):", user.Password)

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data["password"])); err != nil {
		return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"message": "Password salah",
		})
	}

	// Pastikan UserSecretKey sudah didefinisikan
	if UserSecretKey == "" {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Token secret key belum didefinisikan",
		})
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.Id)),
		ExpiresAt: time.Now().Add(30 * time.Minute).Unix(),
		Subject:   "user",
	}).SignedString([]byte(UserSecretKey))

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal membuat token",
		})
	}

	ctx.Cookie(&fiber.Cookie{
		Name:     "jwtUser",
		Value:    token,
		Expires:  time.Now().Add(30 * time.Minute),
		HTTPOnly: true,
		Secure:   true,
	})

	return ctx.JSON(fiber.Map{
		"status":  "success",
		"message": "Login berhasil",
		"token":   token,
	})
}





// UserProfile returns user data by token
func (uc *UserController) UserProfile(ctx *fiber.Ctx) error {
	cookie := ctx.Cookies("jwtUser")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(UserSecretKey), nil
	})

	if err != nil || !token.Valid {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthenticated",
		})
	}

	claims := token.Claims.(*jwt.StandardClaims)

	var user models.User
	if err := database.DB.First(&user, claims.Issuer).Error; err != nil {
		return ctx.Status(http.StatusNotFound).JSON(fiber.Map{
			"message": "User tidak ditemukan",
		})
	}

	return ctx.JSON(fiber.Map{
		"status": "success",
		"data":   user,
	})
}

// LogoutUser clears the user token
func (uc *UserController) LogoutUser(ctx *fiber.Ctx) error {
	ctx.Cookie(&fiber.Cookie{
		Name:     "jwtUser",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
		Secure:   true,
	})

	return ctx.JSON(fiber.Map{
		"message": "Logout berhasil",
	})
}
