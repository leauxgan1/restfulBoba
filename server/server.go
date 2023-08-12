package server

import (
	"fmt"
	"restfulBoba/data_stores"
  "errors"
	"github.com/gofiber/fiber/v2"
)

func getLocations(c *fiber.Ctx) error {
  err := c.JSON(data_stores.Locations)
  return err
}

func getLocation(c *fiber.Ctx) error {
  id := c.Params("id")
  for i := 0; i < len(data_stores.Locations); i++ {
    currStore := data_stores.Locations[i].ID
    if currStore == id {
      err := c.JSON(currStore)
      return err
    }
  }
  return errors.New(fmt.Sprintf("Location with id %s not found.",id))
}

func addMenuItem(c *fiber.Ctx) error {
  var err error
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

	if err := app.Get("/locations", getLocations); err != nil {
	  fmt.Println("Locations not found")
  }
  //	app.Get("/:value", func(c *fiber.Ctx) error {
	//	return c.SendString(c.Params("value") + " Bazoingas!?!!?!")
	//})

	// GET http://localhost:8080/hello%20world
	//app.Get("/:name?", func(c *fiber.Ctx) error {
	//	return c.SendString("Hello there, " + c.Params("name"))
	//})

	app.Listen(":3000")

}
