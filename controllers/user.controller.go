package controllers

import (
	"encoding/json"
	"log"
	"time"

	"github.com/Leon4rdoMonteiro/gofiber.api/models"
	"github.com/Leon4rdoMonteiro/gofiber.api/services"
	"github.com/Leon4rdoMonteiro/gofiber.api/validations"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func CreateUser(ctx *fiber.Ctx) error {
	body := ctx.Body()

	user := models.User{}

	user.Id = utils.UUIDv4()
	user.CreatedAt = time.Now().UTC()

	var err error

	if err = json.Unmarshal(body, &user); err != nil {
		log.Fatal(err)
	}

	validationErrors := validations.ValidateStruct(user)

	if validationErrors != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(validationErrors)
	}

	return ctx.Status(fiber.StatusCreated).JSON(user)
}

func GetUsers(ctx *fiber.Ctx) error {
	response := services.GetUserService()

	return ctx.Status(fiber.StatusOK).JSON(response)
}
