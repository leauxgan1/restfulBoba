package server

import (
	"fmt"
	"restfulBoba/data_stores"

	"github.com/gofiber/fiber/v2"
)

func getLocations(c *fiber.Ctx) error {
	err := c.JSON(data_stores.Locations)
	return err
}

func LoadHandlers() {
	app := fiber.New(fiber.Config{
		ServerHeader:  "Fiber",
		CaseSensitive: true,
		StrictRouting: true,
		AppName:       "RestfulBoba",
	})

	app.Static("/", "./static")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	if err := app.Get("/locations/:id", getLocations); err != nil {
		fmt.Println("The error was totally handled, totally.")
	}
	app.Get("/:value", func(c *fiber.Ctx) error {
		return c.SendString(c.Params("value") + " Bazoingas!?!!?!")
	})

	// GET http://localhost:8080/hello%20world
	app.Get("/:name?", func(c *fiber.Ctx) error {
		return c.SendString("Hello there, " + c.Params("name"))
	})

	app.Listen(":3000")

}
