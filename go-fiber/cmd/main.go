package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/naixatwork/trivia/database"
	"os"
)

const port = ":3000"

func main() {
	database.ConnectDb()
	app := fiber.New()

	setupRoutes(app)
	fmt.Printf("Listening to port %v", port)

	if err := app.Listen(port); err != nil {
		fmt.Printf("error: %s \n", err)
		os.Exit(1)
	}
}
