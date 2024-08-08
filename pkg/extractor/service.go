package extractor

import (
	"github.com/gofiber/fiber/v2/log"
	e "github.com/rs-anantmishra/metubeplus/pkg/entities"
	g "github.com/rs-anantmishra/metubeplus/pkg/global"
)

type IService interface {
	ExtractIngestMetadata(p e.IncomingRequest) bool // here we have an option to dl subs as well, when the metadata is available.
	ExtractIngestMedia()                            //in case it was a metadata only files, youre free to dl video at a later time.
	ExtractSubtitlesOnly(string) bool               // here we are navigating to a Video and downloading subs for it.
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

func (s *service) ExtractIngestMetadata(p e.IncomingRequest) bool {
	metadata, fp := s.download.ExtractMetadata()
	sequencedVideoIds, lstSMI := s.repository.SaveMetadata(metadata, fp)
	//error check here before continuing exec for thumbs and subs

	thumbnail := s.download.ExtractThumbnail(fp, sequencedVideoIds, lstSMI)
	s.repository.SaveThumbnail(thumbnail)

	if p.SubtitlesReq {
		subtitles := s.download.ExtractSubtitles(fp, sequencedVideoIds, lstSMI)
		s.repository.SaveSubtitles(subtitles)
	}
	return true
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
			videoTitle, playlistId, err := s.repository.GetVideoFileInfo(activeItem[0].VideoId)

			if err != nil {
				log.Info(err)
			}

			fileInfo := s.download.GetDownloadedMediaFileInfo(videoTitle, playlistId)
			dbResult := s.repository.SaveMediaContent(fileInfo)

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
