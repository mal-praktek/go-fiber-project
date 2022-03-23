package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tegarsubkhan236/go-fiber-project/database"
	"log"
)

func main() {
	app := fiber.New()
	database.ConnectDB()

	app.Get("/", hallo)
	app.Get("/:name", halloRandomName)

	log.Fatal(app.Listen(":3000"))
}

func hallo(c *fiber.Ctx) error {
	return c.SendString("Hallo World")
}

func halloRandomName(c *fiber.Ctx) error {
	return c.SendString("Hallo " + c.Params("name"))
}
