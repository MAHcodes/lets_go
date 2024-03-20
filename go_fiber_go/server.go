package main

import (
	"github.com/MAHcodes/lets_go/go_fiber_go/controllers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

  app.Get("/:name", controllers.GreetUser)

	app.Listen(":3000")
}
