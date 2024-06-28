package extractor

import e "github.com/rs-anantmishra/metubeplus/pkg/entities"

type IService interface {
	ExtractIngestMetadata() bool
	ExtractIngestVideo(meta e.MediaInformation) bool
}

type service struct {
	repository IRepository
	download   IDownload
}

func Instantiate(r IRepository, d IDownload) IService {
	return &service{
		repository: r,
		download:   d,
	}
}

func (s *service) ExtractIngestMetadata() bool {
	metadata := s.download.ExtractMetadata()
	saveMetaData := s.repository.SaveMetadata(metadata)

	if saveMetaData {
		s.download.ExtractThumbnail(metadata)
		s.download.ExtractSubtitles()
	}

	return true
}

func (s *service) ExtractIngestVideo(m e.MediaInformation) bool {
	return false
}
