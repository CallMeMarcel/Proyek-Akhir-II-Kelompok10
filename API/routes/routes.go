package routes

import (
	// Admin
	adminControllers "api/controllers/admin"

	// User
	userControllers "api/controllers/user"

	// Service
	"api/services"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	// ------------------- ADMIN ROUTES -------------------
	admin := app.Group("/admin")

	// Auth Admin
	admin.Post("/register", adminControllers.Register)
	admin.Post("/login", adminControllers.Login)
	admin.Get("/profile", adminControllers.Profile)
	admin.Post("/logout", adminControllers.Logout)

	// Category Management
	admin.Post("/category", adminControllers.CreateCategory)
	admin.Get("/category/index", adminControllers.IndexCategory)
	admin.Get("/category/:id", adminControllers.ShowCategory)
	admin.Put("/category/:id", adminControllers.UpdateCategory)
	admin.Delete("/category/:id", adminControllers.DeleteCategory)

	// Product Management
	admin.Post("/product", adminControllers.CreateProduct)
	admin.Get("/product", adminControllers.IndexProduct)
	admin.Get("/product/:id", adminControllers.ShowProduct)
	admin.Put("/product/:id", adminControllers.UpdateProduct)
	admin.Delete("/product/:id", adminControllers.DeleteProduct)

	// ------------------- USER ROUTES -------------------
	user := app.Group("/user")

	// Inisialisasi service dan controller user
	userService := services.NewUserService()
	userController := userControllers.NewUserController(userService)

	// Auth User
	user.Post("/register", userController.RegisterUser)
	user.Post("/login", userController.LoginUser)
	user.Get("/profile", userController.UserProfile)
	user.Post("/logout", userController.LogoutUser)
}
