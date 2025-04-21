package controllers

import (
	"fmt"
	"os"
	"api/database"
	"api/models"
	"path/filepath"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

const ImageProduct = "./ProductImg"

func init() {
	if _, err := os.Stat(ImageProduct); os.IsNotExist(err) {
		os.Mkdir(ImageProduct, os.ModePerm)
	}
}

func CreateProduct(ctx *fiber.Ctx) error {
	name := ctx.FormValue("name")
	description := ctx.FormValue("description")
	status := ctx.FormValue("status")
	title := ctx.FormValue("title")
	priceStr := ctx.FormValue("price")
	idCategoryStr := ctx.FormValue("id_category")

	// Validasi field wajib
	if name == "" || description == "" || status == "" || title == "" || priceStr == "" || idCategoryStr == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "All fields are required",
		})
	}

	// Konversi price & id_category
	price, err := strconv.Atoi(priceStr)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid price format",
		})
	}
	idCategory, err := strconv.Atoi(idCategoryStr)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid category format",
		})
	}

	// Upload gambar
	image, err := ctx.FormFile("image")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Image is required",
		})
	}
	filename := fmt.Sprintf("Product_%s%s", name, filepath.Ext(image.Filename))
	if err := ctx.SaveFile(image, filepath.Join(ImageProduct, filename)); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to save image",
		})
	}

	// Simpan ke database
	product := models.Product{
		
		Description: description,
		Status:     status,
		Title:      title,
		Price:      float64(price),
		Image:      filename,
		IdCategory: uint(idCategory),
	}

	database.DB.Create(&product)

	return ctx.JSON(product)
}


func IndexProduct(ctx *fiber.Ctx) error {
	var product []models.Product

	database.DB.Find(&product)

	if len(product) == 0 {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Product not found",
		})
	}

	return ctx.JSON(product)
}

func ShowProduct(ctx *fiber.Ctx) error {
	productIDStr := ctx.Params("id")

	productID, err := strconv.Atoi(productIDStr)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid product ID",
		})
	}

	var product models.Product

	database.DB.Where("id = ?", productID).First(&product)

	if productID != int(product.Id) {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Product not found",
		})
	}

	return ctx.JSON(product)
}

func UpdateProduct(ctx *fiber.Ctx) error {
	productIDStr := ctx.Params("id")
	productID, err := strconv.Atoi(productIDStr)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid product ID",
		})
	}

	var product models.Product
	database.DB.First(&product, productID)
	if product.Id == 0 {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Product not found",
		})
	}

	// Ambil nilai baru
	name := ctx.FormValue("name")
	description := ctx.FormValue("description")
	status := ctx.FormValue("status")
	title := ctx.FormValue("title")
	priceStr := ctx.FormValue("price")
	idCategoryStr := ctx.FormValue("id_category")

	price, err := strconv.Atoi(priceStr)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid price format",
		})
	}
	idCategory, err := strconv.Atoi(idCategoryStr)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid category format",
		})
	}

	newImage, err := ctx.FormFile("image")
	if err == nil {
		if product.Image != "" {
			oldImagePath := filepath.Join(ImageProduct, product.Image)
			os.Remove(oldImagePath)
		}

		filename := fmt.Sprintf("Product_%s%s", name, filepath.Ext(newImage.Filename))
		if err := ctx.SaveFile(newImage, filepath.Join(ImageProduct, filename)); err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to save image",
			})
		}
		product.Image = filename
	}

	// Update data
	product.Description = description
	product.Status = status
	product.Title = title
	product.Price = float64(price)
	product.IdCategory = uint(idCategory)

	database.DB.Save(&product)

	return ctx.JSON(product)
}


func DeleteProduct(ctx *fiber.Ctx) error {
	productIDStr := ctx.Params("id")

	productID, err := strconv.Atoi(productIDStr)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid product ID",
		})
	}

	var product models.Product

	database.DB.Where("id = ?", productID).First(&product)

	if productID != int(product.Id) {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Product not found",
		})
	}

	if err := database.DB.Delete(&product).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to Delete Table",
		})
	}

	if product.Image != "" {
		imagePath := filepath.Join(ImageProduct, product.Image)
		if err := os.Remove(imagePath); err != nil {
			fmt.Printf("Failde to delete image File : %v\n", err)
		}
	}

	return ctx.JSON(fiber.Map{
		"message": "Product deleted successfully",
	})
}
