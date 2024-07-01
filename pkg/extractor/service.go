package extractor

import e "github.com/rs-anantmishra/metubeplus/pkg/entities"

type IService interface {
	ExtractIngestMetadata() bool                     // here we have an option to dl subs as well, when the metadata is available.
	ExtractIngestVideo(meta e.MediaInformation) bool //in case it was a metadata only files, youre free to dl video at a later time.
	ExtractSubtitlesOnly(string) bool                // here we are navigating to a Video and downloading subs for it.
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
		thumbnail := s.download.ExtractThumbnail(metadata)
		s.repository.SaveThumbnail(thumbnail)

		subtitles := s.download.ExtractSubtitles()
		s.repository.SaveSubtitles(subtitles)
	}

	return true
}

func (s *service) ExtractIngestVideo(m e.MediaInformation) bool {
	return false
}

func (s *service) ExtractSubtitlesOnly(videoId string) bool {
	return false
}
