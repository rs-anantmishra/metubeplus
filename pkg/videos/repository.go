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
		if err := rows.Scan(&v.Id, &v.Title, &v.Description, &v.DurationSeconds, &v.OriginalURL, &v.WebpageURL, &v.IsFileDownloaded,
			&v.IsDeleted, &v.Channel.Name, &v.LiveStatus, &v.Domain.Domain, &v.LikesCount, &v.ViewsCount, &v.WatchCount, &v.UploadDate,
			&v.Availability, &v.Format.Format, &v.YoutubeVideoId, &v.CreatedDate); err != nil {
			return nil, fmt.Errorf("error fetching videos: %v", err)
		}
		lstVideos = append(lstVideos, v)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error fetching Videos: %v", err)
	}

	//tags

	tagsRows, err := r.db.Query(p.GetVideoTags_AllVideos)
	if err != nil {
		return nil, fmt.Errorf("error fetching tags for Video: %v", err)
	}
	defer tagsRows.Close()

	var lstTags []entities.Tags
	for tagsRows.Next() {
		var tags entities.Tags
		if err := tagsRows.Scan(&tags.Id, &tags.Name, &tags.IsUsed, &tags.CreatedDate); err != nil {
			return nil, fmt.Errorf("error fetching tags: %v", err)
		}
		lstTags = append(lstTags, tags)
	}

	//assign Tags to video
	for i := range lstVideos {
		for k, elem := range lstTags {
			if lstVideos[i].Id == lstTags[k].Id {
				lstVideos[i].Tags = append(lstVideos[i].Tags, elem)
			}
		}
	}

	//categories
	categoryRows, err := r.db.Query(p.GetVideoCategories_AllVideos)
	if err != nil {
		return nil, fmt.Errorf("error fetching categories for Video: %v", err)
	}
	defer categoryRows.Close()

	var lstCategories []entities.Categories
	for categoryRows.Next() {
		var categories entities.Categories
		if err := categoryRows.Scan(&categories.Id, &categories.Name, &categories.IsUsed, &categories.CreatedDate); err != nil {
			return nil, fmt.Errorf("error fetching categories: %v", err)
		}
		lstCategories = append(lstCategories, categories)
	}

	//assign categories to video
	for i := range lstVideos {
		for k, elem := range lstCategories {
			if lstVideos[i].Id == lstCategories[k].Id {
				lstVideos[i].Categories = append(lstVideos[i].Categories, elem)
			}
		}
	}

	//files

	filesRows, err := r.db.Query(p.GetVideoFiles_AllVideos)
	if err != nil {
		return nil, fmt.Errorf("error fetching files for Video: %v", err)
	}
	defer filesRows.Close()

	var lstFiles []entities.Files
	for filesRows.Next() {
		var files entities.Files
		if err := filesRows.Scan(&files.VideoId, &files.FileType, &files.FileSize, &files.Extension, &files.FilePath, &files.FileName); err != nil {
			return nil, fmt.Errorf("error fetching files: %v", err)
		}
		lstFiles = append(lstFiles, files)
	}

	//assign Files to video
	for i := range lstVideos {
		for k, elem := range lstFiles {
			if lstVideos[i].Id == lstFiles[k].VideoId {
				lstVideos[i].Files = append(lstVideos[i].Files, elem)
			}
		}
	}

	return lstVideos, nil
}
