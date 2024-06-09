package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	m "github.com/rsanantmishra/metubeplus/api/middleware"
	r "github.com/rsanantmishra/metubeplus/api/routes"
	c "github.com/rsanantmishra/metubeplus/config"
)

func main() {
	app := fiber.New(fiber.Config{
		Prefork:       false,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       "MeTube+",
	})

	port := ":" + c.Config("APP_PORT")

	m.SetupMiddleware(app)
	r.SetupRoutes(app)

	log.Fatal(app.Listen(port))
}
