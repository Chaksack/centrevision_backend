package middleware

import (
	"strconv"

	"errors"

	"github.com/Chaksack/centrevision_backend/database"
	"github.com/Chaksack/centrevision_backend/models"
	"github.com/Chaksack/centrevision_backend/utils"
	"github.com/gofiber/fiber/v2"
)

func IsAuthorized(c *fiber.Ctx, page string) error {
	cookie := c.Cookies("jwt")

	Id, err := utils.ParseJwt(cookie)

	if err != nil {
		return err
	}

	userId, _ := strconv.Atoi(Id)

	user := models.User{
		Id: uint(userId),
	}
	database.Database.Db.Preload("Role").Find(&user)

	role := models.Role{
		Id: user.RoleId,
	}

	database.Database.Db.Preload("Permissions").Find(&role)
	if c.Method() == "Get" {
		for _, permission := range role.Permission {
			if permission.Name == "view_"+page || permission.Name == "edit_"+page {
				return nil
			}
		}
	} else {
		for _, permission := range role.Permission {
			if permission.Name == "edit_"+page {
				return nil
			}
		}
	}
	c.Status(fiber.StatusUnauthorized)
	return errors.New("unauthorized")
}
