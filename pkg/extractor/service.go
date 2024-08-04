package extractor

import (
	"time"

	"github.com/gofiber/fiber/v2/log"
	e "github.com/rs-anantmishra/metubeplus/pkg/entities"
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

	i := 0
	for {
		//cleanup of processed
		// s.download.Cleanup()

		//download file
		result := s.download.ExtractMediaContent()
		_ = result //save file details in DB
		//something := s.repository.SaveMediaContent(result)

		if i%10 == 0 {
			log.Info(i, " seconds have passed")
		}
		i++
		duration := time.Second
		time.Sleep(duration)
	}
}

func (s *service) ExtractSubtitlesOnly(videoId string) bool {
	return false
}
