package handler

import (
	"time"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"

	sql "github.com/rs-anantmishra/metubeplus/database"
	en "github.com/rs-anantmishra/metubeplus/pkg/entities"
	ex "github.com/rs-anantmishra/metubeplus/pkg/extractor"
	g "github.com/rs-anantmishra/metubeplus/pkg/global"
)

func NetworkIngestMetadata(c *fiber.Ctx) error {

	//bind incoming data
	params := new(en.IncomingRequest)
	if err := c.BodyParser(params); err != nil {
		return err
	}

	//log incoming data
	log.Info("Request Params:", params)

	//Instantiate
	svcDownloads := ex.NewDownload(*params)
	svcRepo := ex.NewDownloadRepo(sql.DB)
	svcVideos := ex.NewDownloadService(svcRepo, svcDownloads)

	//Process Request
	// No validations for URL/Playlist are needed.
	// If Metadata is not fetched, and there is an error message from yt-dlp
	// just show that error on the UI
	svcVideos.ExtractIngestMetadata(*params)

	return nil
}

func NetworkIngestMedia(c *fiber.Ctx) error {

	//bind incoming data
	params := new(en.QueueDownloads)
	if err := c.BodyParser(params); err != nil {
		return err
	}

	//log incoming data
	log.Info("Request Params:", params)

	//global MPI
	messages := g.NewMessage()

	//Instantiate
	svcDownloads := ex.NewMediaDownload(params.DownloadVideos, &messages[0])
	svcRepo := ex.NewDownloadRepo(sql.DB)
	svcVideos := ex.NewDownloadService(svcRepo, svcDownloads)

	svcVideos.ExtractIngestMedia(params.DownloadVideos)

	return nil
}

func WSTest(c *websocket.Conn) {
	var (
		mt  int
		msg []byte
		err error
	)
	//global MPI
	messages := g.NewMessage()

	for {
		log.Info("msg:", messages[0])

		msg = []byte(`{"download":"` + messages[0] + `"}`)
		mt = websocket.TextMessage

		if err = c.WriteMessage(mt, msg); err != nil {
			log.Info("write:", err)
			break
		}
		duration := time.Second
		time.Sleep(duration)
	}

	// for {
	// 	if mt, msg, err = c.ReadMessage(); err != nil {
	// 		log.Info("read:", err)
	// 		break
	// 	}
	// 	log.Info("recv: %s", string(msg))

	// 	if err = c.WriteMessage(mt, msg); err != nil {
	// 		log.Info("write:", err)
	// 		break
	// 	}
	// }
}

func NetworkIngestAutoSubs(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "success", "message": "Hello i'm ok!", "data": "nil"})
}

func NetworkIngestThumbnail(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "success", "message": "Hello i'm ok!", "data": "nil"})
}
