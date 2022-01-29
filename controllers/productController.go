package controllers

import (
	"strconv"

	"github.com/Chaksack/centrevision_backend/database"
	"github.com/Chaksack/centrevision_backend/models"
	"github.com/gofiber/fiber/v2"
)

func AllProducts(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))

	return c.JSON(models.Paginate(database.Database.Db, &models.Product{}, page))

}

func CreateProduct(c *fiber.Ctx) error {
	var product models.Product

	if err := c.BodyParser(&product); err != nil {
		return err
	}

	database.Database.Db.Create(&product)

	return c.JSON(product)

}

func GetProduct(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	product := models.Product{
		Id: uint(id),
	}
	database.Database.Db.Find(&product)
	return c.JSON(product)
}

func UpdateProduct(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	product := models.Product{
		Id: uint(id),
	}
	if err := c.BodyParser(&product); err != nil {
		return err
	}
	database.Database.Db.Model(&product).Updates(product)
	return c.JSON(product)
}

func DeleteProduct(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	product := models.Product{
		Id: uint(id),
	}

	database.Database.Db.Delete(&product)
	return nil
}
