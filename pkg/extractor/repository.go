package extractor

import (
	"database/sql"

	e "github.com/rs-anantmishra/metubeplus/pkg/entities"
)

type IRepository interface {
	SaveMetadata([]e.MediaInformation) (bool, int)
	SaveThumbnail([]e.Files) (bool, int)
	SaveSubtitles([]e.Files) (bool, int)
	SaveMediaContent([]e.Files) (bool, int)
}

type repository struct {
	db *sql.DB
}

func NewDownloadRepo(Database *sql.DB) IRepository {
	return &repository{
		db: Database,
	}
}

func (r *repository) SaveMetadata(metadata []e.MediaInformation) (bool, int) {

	if len(metadata) < 1 {
		return false, 0
	}

	return true, 1
}

func (r *repository) SaveThumbnail(file []e.Files) (bool, int) {

	return true, 1
}

func (r *repository) SaveSubtitles(file []e.Files) (bool, int) {

	return true, 1
}

func (r *repository) SaveMediaContent(file []e.Files) (bool, int) {
	return true, 1
}
