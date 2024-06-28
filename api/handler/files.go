package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func UploadMedia(c *fiber.Ctx) error {
	//log context
	log.Info("Request Params:", c)

	return c.JSON(fiber.Map{"status": "success", "message": "Hello i'm ok!", "data": nil})
}

// Local or Network Media via NFS or SAMBA
func IngestMedia(c *fiber.Ctx) error {
	//log context
	log.Info("Request Params:", c)

	return c.JSON(fiber.Map{"status": "success", "message": "Hello i'm ok!", "data": nil})
}
