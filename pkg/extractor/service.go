package extractor

import (
	"encoding/base64"
	"os"
	"time"

	"github.com/gofiber/fiber/v2/log"
	p "github.com/rs-anantmishra/metubeplus/api/presenter"
	e "github.com/rs-anantmishra/metubeplus/pkg/entities"
	g "github.com/rs-anantmishra/metubeplus/pkg/global"
)

type IService interface {
	ExtractIngestMetadata(p e.IncomingRequest) []p.CardsInfoResponse // here we have an option to dl subs as well, when the metadata is available.
	ExtractIngestMedia()                                             //in case it was a metadata only files, youre free to dl video at a later time.
	ExtractSubtitlesOnly(string) bool                                // here we are navigating to a Video and downloading subs for it.
}

type service struct {
	repository IRepository
	download   IDownload
}

func NewDownloadService(r IRepository, d IDownload) IService {
	return &service{
		repository: r,
		download:   d,
	}
}

func (s *service) ExtractIngestMetadata(params e.IncomingRequest) []p.CardsInfoResponse {
	metadata, fp := s.download.ExtractMetadata()
	sequencedVideoIds, lstSMI := s.repository.SaveMetadata(metadata, fp)
	//error check here before continuing exec for thumbs and subs

	thumbnails := s.download.ExtractThumbnail(fp, sequencedVideoIds, lstSMI)
	s.repository.SaveThumbnail(thumbnails)

	var subtitles []e.Files
	if params.SubtitlesReq {
		subtitles = s.download.ExtractSubtitles(fp, sequencedVideoIds, lstSMI)
		s.repository.SaveSubtitles(subtitles)
	}

	response := createMetadataResponse(metadata, sequencedVideoIds, subtitles, params.SubtitlesReq, thumbnails)
	return response
}

func (s *service) ExtractIngestMedia() {
	defer falsifyQueueAlive()

	//cleanup of processed
	s.download.Cleanup()

	lstDownloads := g.NewDownloadStatus()
	activeItem := g.NewActiveItem()

	if len(lstDownloads) > 0 {
		for i := 0; i < len(lstDownloads); i++ {

			//skip empties
			if lstDownloads[i].State == g.Completed || lstDownloads[i].VideoURL == "" {
				continue
			}

			//copy to active-item
			activeItem[0] = lstDownloads[i]

			//download file
			lstDownloads[i].State = s.download.ExtractMediaContent()
			smi, fp, err := s.repository.GetVideoFileInfo(activeItem[0].VideoId)

			if err != nil {
				log.Info(err)
			}

			fileInfo := s.download.GetDownloadedMediaFileInfo(smi, fp)
			dbResult := s.repository.SaveMediaContent(fileInfo)

			//wait before next
			duration := time.Second
			time.Sleep(duration)

			activeItem = clearActiveItem(activeItem)

			_ = dbResult
		}
	}
}

func (s *service) ExtractSubtitlesOnly(videoId string) bool {
	return false
}

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

func createMetadataResponse(metadata []e.MediaInformation, sequencedVideoIds []int, subtitles []e.Files, subtitlesReq bool, thumbnails []e.Files) []p.CardsInfoResponse {
	//bind here to presenter entity
	var cardMetaDataInfoList []p.CardsInfoResponse
	const _blank string = ""

	for i, elem := range metadata {
		var cardMetaDataInfo p.CardsInfoResponse

		cardMetaDataInfo.Channel = elem.Channel
		cardMetaDataInfo.CreatedDate = int(time.Now().Unix())
		cardMetaDataInfo.Description = elem.Description
		cardMetaDataInfo.Domain = elem.Domain
		cardMetaDataInfo.Duration = elem.Duration
		cardMetaDataInfo.IsDeleted = false
		cardMetaDataInfo.IsFileDownloaded = false
		cardMetaDataInfo.MediaURL = _blank
		cardMetaDataInfo.OriginalURL = elem.OriginalURL
		cardMetaDataInfo.Playlist = elem.PlaylistTitle
		cardMetaDataInfo.PlaylistVideoIndex = elem.PlaylistIndex
		cardMetaDataInfo.Title = elem.Title
		cardMetaDataInfo.VideoFormat = elem.Format
		cardMetaDataInfo.VideoId = sequencedVideoIds[i]
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
		if metadata[0].PlaylistTitle == _blank {
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
	case "png":
		base64EncodedImage += "data:image/png;base64,"
	case "webp":
		base64EncodedImage += "data:image/webp;base64,"
	}

	base64EncodedImage += base64.StdEncoding.EncodeToString(bytes)

	return base64EncodedImage
}
