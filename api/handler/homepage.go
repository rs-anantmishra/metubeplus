package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

// Get All Videos
func GetAllVideos(c *fiber.Ctx) error {
	//log context
	log.Info("Request Params:", c)

	return c.JSON(fiber.Map{"status": "success", "message": "Hello i'm ok!", "data": nil})
}

// Get All Playlists
func GetAllPlaylists(c *fiber.Ctx) error {
	//log context
	log.Info("Request Params:", c)

	return c.JSON(fiber.Map{"status": "success", "message": "Hello i'm ok!", "data": nil})
}

// Get All Audios
func GetAllAudios(c *fiber.Ctx) error {
	//log context
	log.Info("Request Params:", c)

	return c.JSON(fiber.Map{"status": "success", "message": "Hello i'm ok!", "data": nil})
}

// Get Media By Tags
func GetMediaByTags(c *fiber.Ctx) error {
	//log context
	log.Info("Request Params:", c)

	return c.JSON(fiber.Map{"status": "success", "message": "Hello i'm ok!", "data": nil})
}

// Get Media By Categories
func GetVideosByCategories(c *fiber.Ctx) error {
	//log context
	log.Info("Request Params:", c)

	return c.JSON(fiber.Map{"status": "success", "message": "Hello i'm ok!", "data": nil})
}

// Get Media By Domain
func GetVideosByDomain(c *fiber.Ctx) error {
	//log context
	log.Info("Request Params:", c)

	return c.JSON(fiber.Map{"status": "success", "message": "Hello i'm ok!", "data": nil})
}

// Get Media By Channel
func GetVideosByChannel(c *fiber.Ctx) error {
	//log context
	log.Info("Request Params:", c)

	return c.JSON(fiber.Map{"status": "success", "message": "Hello i'm ok!", "data": nil})
}

// Search Media Files
func GetMediaBySearch(c *fiber.Ctx) error {
	//log context
	log.Info("Request Params:", c)

	return c.JSON(fiber.Map{"status": "success", "message": "Hello i'm ok!", "data": nil})
}

// Search all media by youtube Id
func GetMediaByYoutubeId(c *fiber.Ctx) error {
	//log context
	log.Info("Request Params:", c)

	return c.JSON(fiber.Map{"status": "success", "message": "Hello i'm ok!", "data": nil})
}

func GetMediaByPhysicalLocation(c *fiber.Ctx) error {
	//log context
	log.Info("Request Params:", c)

	return c.JSON(fiber.Map{"status": "success", "message": "Hello i'm ok!", "data": nil})
}
