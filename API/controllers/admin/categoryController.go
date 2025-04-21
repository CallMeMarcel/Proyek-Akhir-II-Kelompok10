package controllers

import (
	"api/database"
	"api/models"
	"strconv"

	fiber "github.com/gofiber/fiber/v2"
)

func CreateCategory(ctx *fiber.Ctx) error {
	name := ctx.FormValue("name")
	if name == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Name is required",
		})
	}

	description := ctx.FormValue("description")
	if description == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Description is required",
		})
	}

	adminID := ctx.Locals("id").(string)

	var admin models.Admin
	database.DB.Where("id = ?", adminID).Find(&admin)

	category := models.Category{
		Name: name,
	}

	database.DB.Create(&category)

	return ctx.JSON(category)
}
func IndexCategory(ctx *fiber.Ctx) error {

	var category []models.Category

	database.DB.Find(&category)

	if len(category) == 0 {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Category not found",
		})
	}

	return ctx.JSON(category)
}

func ShowCategory(ctx *fiber.Ctx) error {
	categoryIDStr := ctx.Params("id")

	categoryID, err := strconv.Atoi(categoryIDStr)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid ID Format",
		})
	}

	var category models.Category

	database.DB.Where("id = ?", categoryID).Find(&category)

	return ctx.JSON(category)
}

func UpdateCategory(ctx *fiber.Ctx) error {
	categoryIDStr := ctx.Params("id")

	categoryID, err := strconv.Atoi(categoryIDStr)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid ID Format",
		})
	}

	var category models.Category

	database.DB.Where("id = ?", categoryID).Find(&category)

	if categoryID != int(category.Id) {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Category Not Found",
		})
	}

	nameUpdate := ctx.FormValue("name")
	if nameUpdate == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Category Name is Required",
		})
	}

	descriptionUpdate := ctx.FormValue("description")
	if descriptionUpdate == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Category Description is Required",
		})
	}

	updates := map[string]interface{}{
		"Name":        nameUpdate,
		"Description": descriptionUpdate,
	}

	result := database.DB.Model(&category).Updates(updates)
	if result.Error != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Couldn't update category",
		})
	}

	return ctx.JSON(category)
}

func DeleteCategory(ctx *fiber.Ctx) error {
	categoryIDStr := ctx.Params("id")

	categoryID, err := strconv.Atoi(categoryIDStr)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid ID Format",
		})
	}

	var category models.Category

	database.DB.Where("id = ?", categoryID).Find(&category)

	if categoryID != int(category.Id) {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Category Not Found",
		})
	}

	if err := database.DB.Delete(&category).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not delete category",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "Delete category successfully",
	})
}
