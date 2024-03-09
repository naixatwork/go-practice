package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/naixatwork/trivia/handlers"
)

func setupRoutes(app *fiber.App) {
	app.Get("/fact", func(c *fiber.Ctx) error {
		return handlers.ListFacts(c)
	})

	app.Post("/fact", func(c *fiber.Ctx) error {
		return handlers.CreateFact(c)
	})
}
