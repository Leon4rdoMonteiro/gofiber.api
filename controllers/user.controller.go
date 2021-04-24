package controllers

import (
	"github.com/Leon4rdoMonteiro/gofiber.api/services"

	"github.com/gofiber/fiber/v2"
)

func GetUsers(ctx *fiber.Ctx) error {
	response := services.GetUserService()

	return ctx.Status(200).JSON(response)
}
