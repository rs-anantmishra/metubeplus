package handler

import (
	"github.com/gofiber/fiber/v2/log"

	"github.com/gofiber/fiber/v2"
	en "github.com/rs-anantmishra/metubeplus/entities"
	ex "github.com/rs-anantmishra/metubeplus/pkg/extractor"
)

func Hello(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "success", "message": "Hello i'm ok!", "data": nil})
}

func MetadataCheck(c *fiber.Ctx) error {

	//bind incoming data
	params := new(en.IncomingRequest)
	if err := c.BodyParser(params); err != nil {
		return err
	}

	//log incoming data
	log.Info("Request Params:", params)

	//Instantiate
	svcDownloads := ex.InstantiateDownload(params.DataIdReq)
	svcRepo := ex.InstantiateRepo("")
	svcVideos := ex.Instantiate(svcRepo, svcDownloads)

	//Process Request
	// No validations for URL/Playlist are needed.
	// If Metadata is not fetched, and there is an error message from yt-dlp
	// just show that error on the UI
	// Get Metadata will be called from inside Get Video also.
	// Separate API only for cases when the use wants to not download video, and play it rom yt
	// but have it metadata downloaded.
	svcVideos.GetMetadata(true)

	return nil
}
