package routes

import "github.com/gofiber/fiber/v2"

func Hello(c *fiber.Ctx) error {
	return c.SendString("Hola desde la funcion Hello")
}
