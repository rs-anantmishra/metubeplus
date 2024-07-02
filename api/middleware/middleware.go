package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/pprof"
)

func SetupMiddleware(app *fiber.App) {

	//pprof
	app.Use(pprof.New())

	//logger
	app.Use(logger.New())

	//setup CORS
	app.Use(cors.New())

	//Rate-Limiter
	app.Use(limiter.New(limiter.Config{
		Max:               200000,
		Expiration:        30 * time.Second,
		LimiterMiddleware: limiter.SlidingWindow{},
	}))
}
