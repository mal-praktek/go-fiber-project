package controllers

import "github.com/gofiber/fiber/v2"

func TestRender(ctx *fiber.Ctx) error {
	return ctx.Render("index", fiber.Map{
		"FiberTitle": "Hello From Fiber Html Engine",
	})
}
