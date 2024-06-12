package handler

import (
	"github.com/gofiber/fiber/v2"
	v "github.com/rs-anantmishra/metubeplus/pkg/videos"
)

func Hello(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "success", "message": "Hello i'm ok!", "data": nil})
}

func MetadataCheck(c *fiber.Ctx) error {

	link := c.Query("data")

	svcDownloads := v.InstantiateDownload(link)
	svcRepo := v.InstantiateRepo("")

	svcVideos := v.Instantiate(svcRepo, svcDownloads)
	svcVideos.GetMetadata(true)

	return nil
}
