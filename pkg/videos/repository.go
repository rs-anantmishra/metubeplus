package videos

import (
	"database/sql"
	"fmt"

	p "github.com/rs-anantmishra/metubeplus/database"
	"github.com/rs-anantmishra/metubeplus/pkg/entities"
)

type IRepository interface {
	GetAllVideos() ([]entities.Videos, error)
}

type repository struct {
	db *sql.DB
}

func NewVideoRepo(Database *sql.DB) IRepository {
	return &repository{
		db: Database,
	}
}

// GetAllVideos implements IRepository.
func (r *repository) GetAllVideos() ([]entities.Videos, error) {

	var lstVideos []entities.Videos

	rows, err := r.db.Query(p.GetAllVideos_Info)
	if err != nil {
		return nil, fmt.Errorf("error fetching Videos: %v", err)
	}
	defer rows.Close()

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var v entities.Videos
		if err := rows.Scan(&v.Id, &v.Title, &v.Description, &v.DurationSeconds, &v.OriginalURL,
			&v.WebpageURL, &v.IsFileDownloaded, &v.IsDeleted, &v.Channel.Name, &v.Playlist.Title,
			&v.LiveStatus, &v.Domain.Domain, &v.Availability, &v.Format.Format, &v.YoutubeVideoId,
			&v.CreatedDate, &v.ThumbnailFilePath, &v.VideoFilePath); err != nil {
			return nil, fmt.Errorf("error fetching videos: %v", err)
		}
		lstVideos = append(lstVideos, v)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error fetching Videos: %v", err)
	}

	return lstVideos, nil
}
