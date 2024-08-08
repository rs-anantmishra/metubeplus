package handler

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"

	res "github.com/rs-anantmishra/metubeplus/api/presenter"
	cfg "github.com/rs-anantmishra/metubeplus/config"
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

	maxQueueLength, _ := strconv.Atoi((cfg.Config("MAX_QUEUE")))
	//bind incoming data
	params := new(en.QueueDownloads)

	if err := c.BodyParser(params); err != nil {
		return err
	}

	//log incoming data
	log.Info("Request Params:", params)

	//global MPI
	lstDownloads := g.NewDownloadStatus()
	qAlive := g.NewQueueAlive()
	currentQueueIndex := g.NewCurrentQueueIndex()

	if maxQueueLength-currentQueueIndex[0]-len(params.DownloadVideos) >= 0 {
		for idx := range params.DownloadVideos {
			lstDownloads[currentQueueIndex[0]] = g.DownloadStatus{VideoId: params.DownloadVideos[idx].VideoId, VideoURL: params.DownloadVideos[idx].VideoURL, StatusMessage: "", State: g.Queued}
			currentQueueIndex[0]++
		}
	} else {
		//send error response that queue is full. Please wait for existing downloads to complete.
	}

	//Instantiate
	svcDownloads := ex.NewDownload(en.IncomingRequest{})
	svcRepo := ex.NewDownloadRepo(sql.DB)
	svcVideos := ex.NewDownloadService(svcRepo, svcDownloads)

	if qAlive[0] != 1 {
		qAlive[0] = 1
		go svcVideos.ExtractIngestMedia()
	}

	return nil
}

func DownloadStatus(c *websocket.Conn) {
	var (
		mt  int
		msg []byte
		err error
	)
	//global MPI
	activeItem := g.NewActiveItem()
	mt = websocket.TextMessage

	for {
		if len(activeItem) > 0 {

			dsr := res.DownloadStatusResponse{Message: activeItem[0].StatusMessage, VideoURL: activeItem[0].VideoURL}
			jsonData, e := json.Marshal(dsr)
			if e != nil {
				log.Info(e)
			}
			msg = []byte(jsonData)

			if err = c.WriteMessage(mt, msg); err != nil {
				log.Info("write:", err)
				break
			}
		} else {
			c.Close()
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
