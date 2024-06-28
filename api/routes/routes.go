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

	//Network Downloads: Playlist, Videos or Audios
	download := app.Group("/download", logger.New())
	download.Post("/metadata", handler.ExtractMetadata) //Get Metadata Only and add to library [Playlists, Videos]
	download.Post("/media", handler.ExtractMedia)       //Download Video File and update app db [Playlists, Videos]

	//Specific to network Video and Audio only
	download.Post("/autosubs", handler.ExtractAutoSubs)   //Download auto-subs from yt for a video that exists in library [Videos]
	download.Post("/thumbnail", handler.ExtractThumbnail) //Download thumbnail for a video existing in library [Videos]

	//Todo: Homepage
	//Todo: Tags & Categories
	//Todo: Patterns
	//Todo: Files
}
