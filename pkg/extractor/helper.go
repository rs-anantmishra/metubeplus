package extractor

import (
	"encoding/base64"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2/log"
	p "github.com/rs-anantmishra/metubeplus/api/presenter"
	e "github.com/rs-anantmishra/metubeplus/pkg/entities"
	g "github.com/rs-anantmishra/metubeplus/pkg/global"
)

func falsifyQueueAlive() {
	qa := g.NewQueueAlive()
	qa[0] = 0
}

func clearActiveItem(activeItem []g.DownloadStatus) []g.DownloadStatus {

	//empty out active item
	activeItem[0].VideoURL = ""
	activeItem[0].VideoId = 0
	activeItem[0].StatusMessage = ""
	activeItem[0].State = 0

	return activeItem
}

func createMetadataResponse(lstSavedInfo []e.SavedInfo, subtitles []e.Files, subtitlesReq bool, thumbnails []e.Files) []p.CardsInfoResponse {
	//bind here to presenter entity
	var cardMetaDataInfoList []p.CardsInfoResponse
	const _blank string = ""

	for _, elem := range lstSavedInfo {
		var cardMetaDataInfo p.CardsInfoResponse

		cardMetaDataInfo.Channel = elem.MediaInfo.Channel
		cardMetaDataInfo.CreatedDate = int(time.Now().Unix())
		cardMetaDataInfo.Description = elem.MediaInfo.Description
		cardMetaDataInfo.Domain = elem.MediaInfo.Domain
		cardMetaDataInfo.Duration = elem.MediaInfo.Duration
		cardMetaDataInfo.IsDeleted = false
		cardMetaDataInfo.IsFileDownloaded = false
		cardMetaDataInfo.MediaURL = _blank
		cardMetaDataInfo.OriginalURL = elem.MediaInfo.OriginalURL
		cardMetaDataInfo.Playlist = elem.MediaInfo.PlaylistTitle
		cardMetaDataInfo.PlaylistVideoIndex = elem.MediaInfo.PlaylistIndex
		cardMetaDataInfo.Title = elem.MediaInfo.Title
		cardMetaDataInfo.VideoFormat = elem.MediaInfo.Format
		cardMetaDataInfo.VideoId = elem.VideoId
		cardMetaDataInfo.WatchCount = 0

		cardMetaDataInfoList = append(cardMetaDataInfoList, cardMetaDataInfo)
	}

	//subtitles
	if subtitlesReq {
		for i, elem := range subtitles {
			cardMetaDataInfoList[i].SubtitlesURL = elem.FilePath + elem.FileName
		}
	}

	//thumbnails
	for i := range thumbnails {
		if lstSavedInfo[0].MediaInfo.PlaylistTitle == _blank {
			cardMetaDataInfoList[i].Thumbnail = getImagesFromURL(thumbnails[i])
		} else {
			if i == 0 {
				cardMetaDataInfoList[i].PlaylistThumbnail = getImagesFromURL(thumbnails[0])
				continue
			}
			cardMetaDataInfoList[i-1].Thumbnail = getImagesFromURL(thumbnails[i])
			cardMetaDataInfoList[i-1].PlaylistThumbnail = getImagesFromURL(thumbnails[0])
		}
	}

	return cardMetaDataInfoList
}

func getImagesFromURL(file e.Files) string {
	var base64EncodedImage string

	filepath := file.FilePath + "\\" + file.FileName
	// Read the entire file into a byte slice
	bytes, err := os.ReadFile(filepath)
	if err != nil {
		log.Info(err)
	}

	switch file.Extension {
	case "jpeg":
		base64EncodedImage += "data:image/jpeg;base64,"
	case "jpg":
		base64EncodedImage += "data:image/jpg;base64,"
	case "png":
		base64EncodedImage += "data:image/png;base64,"
	case "webp":
		base64EncodedImage += "data:image/webp;base64,"
	}

	base64EncodedImage += base64.StdEncoding.EncodeToString(bytes)
	return base64EncodedImage
}

func getImagesFromURLString(filepath string) string {
	var base64EncodedImage string
	splitter := "."

	// Read the entire file into a byte slice
	bytes, err := os.ReadFile(filepath)
	if err != nil {
		log.Info(err)
	}

	if len(bytes) == 0 {
		filepath = `..\utils\noimage.png`
		bytes, err = os.ReadFile(filepath)
		if err != nil {
			log.Info(err)
		}
	}

	splits := strings.SplitN(filepath, splitter, -1)
	extension := splits[len(splits)-1]

	switch extension {
	case "jpeg":
		base64EncodedImage += "data:image/jpeg;base64,"
	case "jpg":
		base64EncodedImage += "data:image/jpg;base64,"
	case "png":
		base64EncodedImage += "data:image/png;base64,"
	case "webp":
		base64EncodedImage += "data:image/webp;base64,"
	}

	base64EncodedImage += base64.StdEncoding.EncodeToString(bytes)
	return base64EncodedImage
}
