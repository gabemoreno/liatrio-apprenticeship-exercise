package main

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message":   "My name is Gabriel Moreno",
			"timestamp": time.Now().UnixMilli(),
			"did this work?": "yes",
		})
	})

	log.Fatal(app.Listen(":3000"))
}
