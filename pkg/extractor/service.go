package extractor

import e "github.com/rs-anantmishra/metubeplus/pkg/entities"

type IService interface {
	ExtractIngestMetadata(p e.IncomingRequest) bool  // here we have an option to dl subs as well, when the metadata is available.
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

func (s *service) ExtractIngestMetadata(p e.IncomingRequest) bool {
	metadata := s.download.ExtractMetadata()
	saveMetaData, VideoId := s.repository.SaveMetadata(metadata)
	_ = VideoId //Pass to other methods to save in tblFiles

	fp := e.Filepath{Domain: metadata[0].Domain, Channel: metadata[0].Channel, PlaylistTitle: metadata[0].PlaylistTitle}
	//save thumbnails to disk, then - read files from disk and populate metadata into db
	//save subtitles to disk, then - read files from disk and populate metadata into db
	//If video is also downloaded - then, write Video file path else use network paths.
	if saveMetaData {
		thumbnail := s.download.ExtractThumbnail(fp)
		s.repository.SaveThumbnail(thumbnail)

		if p.SubtitlesReq {
			subtitles := s.download.ExtractSubtitles(fp)
			s.repository.SaveSubtitles(subtitles)
		}
	}

	return true
}

func (s *service) ExtractIngestVideo(m e.MediaInformation) bool {
	return false
}

func (s *service) ExtractSubtitlesOnly(videoId string) bool {
	return false
}
