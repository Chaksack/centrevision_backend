package controllers

import (
	"strconv"

	"github.com/Chaksack/centrevision_backend/database"
	"github.com/Chaksack/centrevision_backend/middleware"
	"github.com/Chaksack/centrevision_backend/models"
	"github.com/gofiber/fiber/v2"
)

func AllUsers(c *fiber.Ctx) error {
	if err := middleware.IsAuthorized(c, "users"); err != nil {
		return err
	}

	page, _ := strconv.Atoi(c.Query("page", "1"))
	return c.JSON(models.Paginate(database.Database.Db, &models.User{}, page))

}

func CreateUser(c *fiber.Ctx) error {
	if err := middleware.IsAuthorized(c, "users"); err != nil {
		return err
	}
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return err
	}

	user.SetPassword("1234")

	database.Database.Db.Create(&user)

	return c.JSON(user)

}

func GetUser(c *fiber.Ctx) error {
	if err := middleware.IsAuthorized(c, "users"); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Params("id"))

	user := models.User{
		Id: uint(id),
	}
	database.Database.Db.Preload("Role").Find(&user)
	return c.JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	if err := middleware.IsAuthorized(c, "users"); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Params("id"))

	user := models.User{
		Id: uint(id),
	}
	if err := c.BodyParser(&user); err != nil {
		return err
	}
	database.Database.Db.Model(&user).Updates(user)
	return c.JSON(user)
}

func DeleteUser(c *fiber.Ctx) error {
	if err := middleware.IsAuthorized(c, "users"); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Params("id"))

	user := models.User{
		Id: uint(id),
	}

	database.Database.Db.Delete(&user)
	return nil
}
