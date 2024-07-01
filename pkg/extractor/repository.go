package extractor

import (
	e "github.com/rs-anantmishra/metubeplus/pkg/entities"
)

type IRepository interface {
	SaveMetadata([]e.MediaInformation) bool
	SaveThumbnail([]e.Files) bool
	SaveSubtitles([]e.Files) bool
	SaveMediaContent([]e.Files) bool
}

type repository struct {
	//here we have the db connection object (or the connection string?) to execute queries
	Connection string
}

func InstantiateRepo(conn string) IRepository {
	return &repository{
		Connection: conn,
	}
}

func (r *repository) SaveMetadata(metadata []e.MediaInformation) bool {

	return true
}

func (r *repository) SaveThumbnail(file []e.Files) bool {

	return true
}

func (r *repository) SaveSubtitles(file []e.Files) bool {

	return true
}

func (r *repository) SaveMediaContent(file []e.Files) bool {
	return true
}
