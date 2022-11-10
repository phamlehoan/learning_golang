package books_controller

import (
	"hello/helpers"
	"hello/models"
	"net/http"

	"gorm.io/gorm"

	"github.com/gofiber/fiber/v2"
)

func Index(c *fiber.Ctx) error {
	var books []models.Book
	models.DB.Find(&books)

	return c.JSON(helpers.BuildResponse(true, "Success", books))
}

func Show(c *fiber.Ctx) error {

	id := c.Params("id")
	var book models.Book
	if err := models.DB.First(&book, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(http.StatusNotFound).JSON(
				helpers.BuildErrorResponse("Data not found", err.Error(), book),
			)
		}

		return c.Status(http.StatusInternalServerError).JSON(
			helpers.BuildErrorResponse("Data not found", err.Error(), book),
		)
	}

	return c.JSON(helpers.BuildResponse(true, "Success", book))
}

func Create(c *fiber.Ctx) error {

	var book models.Book
	if err := c.BodyParser(&book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			helpers.BuildErrorResponse("Error", err.Error(), book),
		)
	}

	if err := models.DB.Create(&book).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			helpers.BuildErrorResponse("Error", err.Error(), book),
		)
	}

	return c.JSON(helpers.BuildResponse(true, "Success", book))
}

func Update(c *fiber.Ctx) error {

	id := c.Params("id")

	var book models.Book
	if err := c.BodyParser(&book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			helpers.BuildErrorResponse("Error", err.Error(), book),
		)
	}

	if models.DB.Where("id = ?", id).Updates(&book).RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(
			helpers.BuildErrorResponse("Unable to update data", "", book),
		)
	}

	return c.JSON(helpers.BuildResponse(true, "Success", book))
}

func Delete(c *fiber.Ctx) error {

	id := c.Params("id")

	var book models.Book
	if models.DB.Delete(&book, id).RowsAffected == 0 {
		return c.Status(http.StatusNotFound).JSON(
			helpers.BuildErrorResponse("Unable to delete data", "", book),
		)
	}

	return c.JSON(helpers.BuildResponse(true, "Success", book))
}
