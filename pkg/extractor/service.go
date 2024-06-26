package extractor

import e "github.com/rs-anantmishra/metubeplus/entities"

type IService interface {
	GetMetadata(verobse bool)
	GetVideo(meta e.Metadata) bool
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

func (s *service) GetMetadata(verbose bool) {
	s.download.GetMetadata(true)
}

func (s *service) GetVideo(m e.Metadata) bool {
	return false
}
