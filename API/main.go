package main

import (
	"log"

	"api/database"
	"api/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// Koneksi ke database
	database.Connect()

	// Inisialisasi Fiber
	app := fiber.New()

	// Middleware CORS
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowMethods:     "GET, POST, PUT, DELETE",
		AllowHeaders:     "Content-Type, Authorization, Origin, Accept",
		AllowOrigins:     "*",
	}))

  
	// Setup semua route
	routes.Setup(app)

	// Jalankan server pada port 8000
	if err := app.Listen(":8000"); err != nil {
		log.Fatal(err)
	}
}
