package handler

import (
	"github.com/gofiber/fiber/v2/log"

	"github.com/gofiber/fiber/v2"
	e "github.com/rs-anantmishra/metubeplus/entities"
	v "github.com/rs-anantmishra/metubeplus/pkg/videos"
)

func Hello(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "success", "message": "Hello i'm ok!", "data": nil})
}

func MetadataCheck(c *fiber.Ctx) error {

	//bind incoming data
	params := new(e.IncomingRequest)
	if err := c.BodyParser(params); err != nil {
		return err
	}

	//log incoming data
	log.Info("params are as follows: ", params)

	svcDownloads := v.InstantiateDownload(params.DataIdReq)
	svcRepo := v.InstantiateRepo("")

	svcVideos := v.Instantiate(svcRepo, svcDownloads)
	svcVideos.GetMetadata(true)

	return nil
}
