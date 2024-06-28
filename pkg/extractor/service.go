package extractor

import e "github.com/rs-anantmishra/metubeplus/pkg/entities"

type IService interface {
	ExtractMetadata(verobse bool)
	ExtractVideo(meta e.MediaInformation) bool
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

func (s *service) ExtractMetadata(verbose bool) {
	s.download.ExtractMetadata(true)
}

func (s *service) ExtractVideo(m e.MediaInformation) bool {
	return false
}
