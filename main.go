package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Get("/", hallo)
	app.Get("/:name", halloRandomName)

	app.Listen(":3000")
}

func hallo(c *fiber.Ctx) error {
	return c.SendString("Hallo World")
}

func halloRandomName(c *fiber.Ctx) error {
	return c.SendString("Hallo " + c.Params("name"))
}
