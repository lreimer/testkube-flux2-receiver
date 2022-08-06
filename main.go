package main

import "github.com/gofiber/fiber/v2"

var version = ""
var commit = ""

func main() {
	app := fiber.New()

	app.Get("/version", func(c *fiber.Ctx) error {
		return c.JSON(map[string]string{"version": version, "commit": commit})
	})

	app.Listen(":8080")
}
