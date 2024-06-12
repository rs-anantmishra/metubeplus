package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/rs-anantmishra/metubeplus/api/handler"
)

func SetupRoutes(app *fiber.App) {

	//Middlewares
	api := app.Group("/api", logger.New())

	api.Get("/hello", handler.Hello)
	api.Get("/meta", handler.MetadataCheck)

	//routes Videos, Files, Tags
}
