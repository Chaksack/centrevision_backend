package controllers

import (
	"strconv"

	"github.com/Chaksack/centrevision_backend/database"
	"github.com/Chaksack/centrevision_backend/models"
	"github.com/gofiber/fiber/v2"
)

func AllCategorys(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))

	return c.JSON(models.Paginate(database.Database.Db, &models.Category{}, page))

}

func CreateCategory(c *fiber.Ctx) error {
	var category models.Category

	if err := c.BodyParser(&category); err != nil {
		return err
	}

	database.Database.Db.Create(&category)

	return c.JSON(category)

}

func GetCategory(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	category := models.Category{
		Id: uint(id),
	}
	database.Database.Db.Find(&category)
	return c.JSON(category)
}

func UpdateCategory(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	category := models.Category{
		Id: uint(id),
	}
	if err := c.BodyParser(&category); err != nil {
		return err
	}
	database.Database.Db.Model(&category).Updates(category)
	return c.JSON(category)
}

func DeleteCategory(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	category := models.Category{
		Id: uint(id),
	}

	database.Database.Db.Delete(&category)
	return nil
}
