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
	download.Post("/metadata", handler.NetworkIngestMetadata) //Download Metadata + Thumbnail and save to db for [Playlists, Videos]
	download.Post("/media", handler.NetworkIngestMedia)       //Download Media File(s) and update db for [Playlists, Videos]

	//Specific to network Video and Audio only
	download.Post("/autosubs", handler.NetworkIngestAutoSubs)   //Download auto-subs for a video that exists in library [Videos]
	download.Post("/thumbnail", handler.NetworkIngestThumbnail) //Download thumbnail for a video that exists in library [Videos]

	//Todo: Homepage
	//Todo: Tags & Categories
	//Todo: Patterns
	//Todo: Files
}
