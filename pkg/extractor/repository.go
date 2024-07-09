package extractor

import (
	"database/sql"
	"time"

	p "github.com/rs-anantmishra/metubeplus/database"
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

	//save channel
	result, err := r.db.Exec(p.InsertChannel, metadata[0].Channel, metadata[0].ChannelURL, metadata[0].ChannelId, time.Now().Unix())

	if err != nil {
		return false, 0
	}

	var id int64
	if id, err = result.LastInsertId(); err != nil {
		return false, 0
	}
	return true, int(id)

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
