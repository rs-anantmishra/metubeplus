package handler

import (
	"encoding/json"
	"time"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"

	res "github.com/rs-anantmishra/metubeplus/api/presenter"
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
	lstDownloads := g.NewDownloadStatus()

	for idx := range params.DownloadVideos {
		lstDownloads = append(lstDownloads, &g.DownloadStatus{VideoId: params.DownloadVideos[idx].VideoId, VideoURL: params.DownloadVideos[idx].VideoURL, StatusMessage: "", State: g.Queued})
	}

	//Instantiate
	svcDownloads := ex.NewDownload(en.IncomingRequest{}) // this is just a placeholder
	svcRepo := ex.NewDownloadRepo(sql.DB)
	svcVideos := ex.NewDownloadService(svcRepo, svcDownloads)

	go svcVideos.ExtractIngestMedia(lstDownloads)

	return nil
}

func DownloadStatus(c *websocket.Conn) {
	var (
		mt  int
		msg []byte
		err error
	)
	//global MPI
	lstDownloads := g.NewDownloadStatus()
	mt = websocket.TextMessage

	for {

		dsr := res.DownloadStatusResponse{Message: lstDownloads[0].StatusMessage}
		jsonData, e := json.Marshal(dsr)
		if e != nil {
			log.Info(e)
		}
		msg = []byte(jsonData)

		if err = c.WriteMessage(mt, msg); err != nil {
			log.Info("write:", err)
			break
		}
		duration := time.Second
		time.Sleep(duration)
	}
}

func NetworkIngestAutoSubs(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "success", "message": "Hello i'm ok!", "data": "nil"})
}

func NetworkIngestThumbnail(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "success", "message": "Hello i'm ok!", "data": "nil"})
}
