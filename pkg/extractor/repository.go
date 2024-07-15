package extractor

import (
	"database/sql"
	"fmt"
	"time"

	p "github.com/rs-anantmishra/metubeplus/database"
	e "github.com/rs-anantmishra/metubeplus/pkg/entities"
)

type IRepository interface {
	SaveMetadata([]e.MediaInformation, e.Filepath) int
	SaveThumbnail([]e.Files) int
	SaveSubtitles([]e.Files) int
	SaveMediaContent([]e.Files) int
}

type repository struct {
	db *sql.DB
}

func NewDownloadRepo(Database *sql.DB) IRepository {
	return &repository{
		db: Database,
	}
}

func (r *repository) SaveMetadata(metadata []e.MediaInformation, fp e.Filepath) int {

	resultId := -1

	for _, elem := range metadata {

		//Channel will be same for all items in playlist.
		channelId := genericCheck(*r, elem.ChannelId, "Channel", p.InsertChannelCheck)
		if channelId <= 0 {
			var args []any
			args = append(args, elem.Channel)
			args = append(args, elem.ChannelURL)
			args = append(args, elem.ChannelId)
			args = append(args, time.Now().Unix())

			channelId = genericSave(*r, args, p.InsertChannel)
			_ = channelId
		}

		//playlist info will be same for all in the playlist.
		playlistId := genericCheck(*r, elem.PlaylistId, "Playlist", p.InsertPlaylistCheck)
		if playlistId <= 0 && (elem.PlaylistTitle != "" && elem.PlaylistCount > 0) {
			var args []any
			args = append(args, elem.PlaylistTitle)
			args = append(args, elem.PlaylistCount)
			args = append(args, GetPlaylistFilepath(fp, -1))
			args = append(args, elem.ChannelId)
			args = append(args, 0)
			args = append(args, elem.PlaylistId)
			args = append(args, time.Now().Unix())

			playlistId = genericSave(*r, args, p.InsertPlaylist)
			_ = playlistId
		}

		//Domain will be same for all items in playlist.
		domainId := genericCheck(*r, elem.Domain, "Domain", p.InsertDomainCheck)
		if domainId <= 0 {
			var args []any
			args = append(args, elem.Domain)
			args = append(args, time.Now().Unix())

			domainId = genericSave(*r, args, p.InsertDomain)
			_ = domainId
		}

		//Format will NOT be same for all items in playlist.
		formatId := genericCheck(*r, elem.Format, "Format", p.InsertFormatCheck)
		if formatId <= 0 {
			var args []any
			args = append(args, elem.Format)
			args = append(args, elem.FormatNote)
			args = append(args, elem.Resolution)
			args = append(args, "Video") //It should be Audio for audio only files
			args = append(args, time.Now().Unix())

			formatId = genericSave(*r, args, p.InsertFormat)
			_ = formatId
		}

		ytVideoId := genericCheck(*r, elem.YoutubeVideoId, "Metadata", p.InsertMetadataCheck)
		if ytVideoId < 0 {
			var args []any

			args = append(args, elem.Title)
			args = append(args, elem.Description)
			args = append(args, elem.Duration)
			args = append(args, elem.OriginalURL)
			args = append(args, elem.LiveStatus)
			args = append(args, elem.Availability)
			args = append(args, elem.PlaylistIndex)
			args = append(args, 0) //IsFileDownloaded
			args = append(args, channelId)
			args = append(args, playlistId)
			args = append(args, domainId)
			args = append(args, formatId)
			args = append(args, elem.YoutubeVideoId)
			args = append(args, 0)                 //WatchCount
			args = append(args, 0)                 //IsDeleted
			args = append(args, time.Now().Unix()) //CreatedDate

			ytVideoId = genericSave(*r, args, p.InsertMetadata)
			_ = ytVideoId
		}

		//Tags will NOT be same for all items in playlist.
		var lstTagId []int
		for _, element := range elem.Tags {
			tagId := genericCheck(*r, element, "Tag", p.InsertTagsCheck)
			if tagId <= 0 {
				var args []any
				args = append(args, element)
				args = append(args, 1) // IsUsed
				args = append(args, time.Now().Unix())

				tagId = genericSave(*r, args, p.InsertTags)
				_ = tagId
				//here, we should have a map between tag Id,
				//VideoId which should be used to populate VideoFileTags
				lstTagId = append(lstTagId, tagId)
			}

			videoFileTagId := checkTagsOrCategories(*r, tagId, "VideoFileTag", p.InsertVideoFileTagsCheck, ytVideoId)
			if videoFileTagId < 0 {
				var args []any
				args = append(args, tagId)
				args = append(args, ytVideoId)
				args = append(args, 0)
				args = append(args, time.Now().Unix())
				videoFileTagId = genericSave(*r, args, p.InsertVideoFileTags)
				_ = videoFileTagId
			}
		}

		//Categories will NOT be same for all items in playlist.
		var lstCategoryId []int
		for _, element := range elem.Categories {
			categoryId := genericCheck(*r, element, "Category", p.InsertCategoriesCheck)
			if categoryId <= 0 {
				var args []any
				args = append(args, element)
				args = append(args, 1) // IsUsed
				args = append(args, time.Now().Unix())

				categoryId = genericSave(*r, args, p.InsertCategories)
				_ = categoryId
				//here, we should have a map between categoryId Id,
				//VideoId which should be used to populate VideoFileCategories
				lstCategoryId = append(lstCategoryId, categoryId)
			}

			videoFileCategoryId := checkTagsOrCategories(*r, categoryId, "VideoFileCategory", p.InsertVideoFileCategoriesCheck, ytVideoId)
			if videoFileCategoryId < 0 {
				var args []any
				args = append(args, categoryId)
				args = append(args, ytVideoId)
				args = append(args, 0)
				args = append(args, time.Now().Unix())
				videoFileCategoryId = genericSave(*r, args, p.InsertVideoFileCategories)
				_ = videoFileCategoryId
			}
		}
	}
	return resultId
}

func (r *repository) SaveThumbnail(file []e.Files) int {

	return 1
}

func (r *repository) SaveSubtitles(file []e.Files) int {

	return 1
}

func (r *repository) SaveMediaContent(file []e.Files) int {
	return 1
}

// ////////////////////////////////////////////////////////////////////
// Private Methods ////////////////////////////////////////////////////
// ////////////////////////////////////////////////////////////////////

func genericSave(r repository, args []any, genericQuery string) int {
	resultId := -1

	if resultId < 0 {

		//generic save
		result, err := r.db.Exec(genericQuery, args...)

		//check for errors
		if err != nil {
			fmt.Println("error:", err)
			return resultId
		}

		//get the inserted records Id
		var id int64
		if id, err = result.LastInsertId(); err != nil {
			return resultId
		}
		resultId = int(id)
	}

	return resultId
}

func genericCheck(r repository, Id any, idContext string, genericQuery string) int {
	resultId := -1

	//Check if entry for Id exists?
	chk := r.db.QueryRow(genericQuery, Id)
	if err := chk.Scan(&resultId); err == sql.ErrNoRows {
		fmt.Println("No Rows found for", idContext, "Id", Id)
	}

	return resultId
}

func checkTagsOrCategories(r repository, Id any, idContext string, genericQuery string, videoId int) int {
	resultId := -1

	//Check if entry for Id exists?
	chk := r.db.QueryRow(genericQuery, Id, videoId)
	if err := chk.Scan(&resultId); err == sql.ErrNoRows {
		fmt.Println("No Rows found for", idContext, "Id", Id)
	}

	return resultId
}
