package extractor

import (
	"database/sql"
	"fmt"
	"time"

	p "github.com/rs-anantmishra/metubeplus/database"
	e "github.com/rs-anantmishra/metubeplus/pkg/entities"
)

type IRepository interface {
	SaveMetadata([]e.MediaInformation, e.Filepath) ([]int, []e.SavedMediaInformation)
	SaveThumbnail([]e.Files) []int
	SaveSubtitles([]e.Files) []int
	SaveMediaContent([]e.Files) []int
	GetVideoFileInfo(videoId int) (string, int, error)
}

type repository struct {
	db *sql.DB
}

func NewDownloadRepo(Database *sql.DB) IRepository {
	return &repository{
		db: Database,
	}
}

func (r *repository) SaveMetadata(metadata []e.MediaInformation, fp e.Filepath) ([]int, []e.SavedMediaInformation) {

	var lstSMI []e.SavedMediaInformation
	sequencedVideoIds := make([]int, len(metadata))
	for i := range metadata {
		sequencedVideoIds[i] = -1
	}
	for k, elem := range metadata {
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
		sequencedVideoIds[k] = ytVideoId
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
			sequencedVideoIds[k] = ytVideoId
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
				//So this would be like a User Defined Type in MSSQL which can
				//be sent at once to SQLite
				lstTagId = append(lstTagId, tagId)
			}

			videoFileTagId := tagsOrCategoriesCheck(*r, tagId, "VideoFileTag", p.InsertVideoFileTagsCheck, ytVideoId)
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
				//So this would be like a User Defined Type in MSSQL which can
				//be sent at once to SQLite
				lstCategoryId = append(lstCategoryId, categoryId)
			}

			videoFileCategoryId := tagsOrCategoriesCheck(*r, categoryId, "VideoFileCategory", p.InsertVideoFileCategoriesCheck, ytVideoId)
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

		lstSMI = append(lstSMI, e.SavedMediaInformation{VideoId: ytVideoId, Title: elem.Title, YoutubeVideoId: elem.YoutubeVideoId})
	}
	return sequencedVideoIds, lstSMI
}

func (r *repository) SaveThumbnail(file []e.Files) []int {

	var lstFileIds []int
	for _, elem := range file {
		thumbnailFileId := filesCheck(*r, elem.FileType, elem.VideoId, p.InsertThumbnailFileCheck)
		if thumbnailFileId <= 0 {
			var args []any
			args = append(args, elem.VideoId)
			args = append(args, elem.FileType)
			args = append(args, elem.SourceId)
			args = append(args, elem.FilePath)
			args = append(args, elem.FileName)
			args = append(args, elem.Extension)
			args = append(args, elem.FileSize)
			args = append(args, elem.FileSizeUnit)
			args = append(args, elem.NetworkPath)
			args = append(args, elem.IsDeleted)
			args = append(args, time.Now().Unix())

			thumbnailFileId = genericSave(*r, args, p.InsertFile)
			lstFileIds = append(lstFileIds, thumbnailFileId)
			_ = thumbnailFileId
		}
	}

	return lstFileIds
}

func (r *repository) SaveSubtitles(file []e.Files) []int {

	var lstFileIds []int
	for _, elem := range file {
		//Check is such that app can support multiple lang subs file for 1 video file.
		subsFileId := subsFilesCheck(*r, elem.FileType, elem.VideoId, elem.FileName, p.InsertSubsFileCheck)
		if subsFileId <= 0 {
			var args []any
			args = append(args, elem.VideoId)
			args = append(args, elem.FileType)
			args = append(args, elem.SourceId)
			args = append(args, elem.FilePath)
			args = append(args, elem.FileName)
			args = append(args, elem.Extension)
			args = append(args, elem.FileSize)
			args = append(args, elem.FileSizeUnit)
			args = append(args, elem.NetworkPath)
			args = append(args, elem.IsDeleted)
			args = append(args, time.Now().Unix())

			subsFileId = genericSave(*r, args, p.InsertFile)
			lstFileIds = append(lstFileIds, subsFileId)
			_ = subsFileId
		}
	}

	return lstFileIds
}

func (r *repository) SaveMediaContent(file []e.Files) []int {
	lstFileIds := make([]int, len(file))
	for i := range file {
		lstFileIds[i] = -1
	}

	for _, elem := range file {
		mediaFileId := filesCheck(*r, elem.FileType, elem.VideoId, p.InsertThumbnailFileCheck)
		if mediaFileId <= 0 {
			var args []any
			args = append(args, elem.VideoId)
			args = append(args, elem.FileType)
			args = append(args, elem.SourceId)
			args = append(args, elem.FilePath)
			args = append(args, elem.FileName)
			args = append(args, elem.Extension)
			args = append(args, elem.FileSize)
			args = append(args, elem.FileSizeUnit)
			args = append(args, elem.NetworkPath)
			args = append(args, elem.IsDeleted)
			args = append(args, time.Now().Unix())

			mediaFileId = genericSave(*r, args, p.InsertFile)
			lstFileIds = append(lstFileIds, mediaFileId)
			_ = mediaFileId
		}
	}

	return lstFileIds
}

func (r *repository) GetVideoFileInfo(videoId int) (string, int, error) {

	var videoTitle string
	var playlistId int

	row := r.db.QueryRow(p.GetVideoInformationById, videoId)
	if err := row.Scan(&videoTitle, &playlistId); err != nil {
		if err == sql.ErrNoRows {
			return videoTitle, playlistId, fmt.Errorf("VideoId %d: no such video", videoId)
		}
		return videoTitle, playlistId, fmt.Errorf("VideoById %d: %v", videoId, err)
	}
	return videoTitle, playlistId, nil
}

// Private Methods ////////////////////////////////////////////////////

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

func tagsOrCategoriesCheck(r repository, Id any, idContext string, genericQuery string, videoId int) int {
	resultId := -1

	//Check if entry for Id exists?
	chk := r.db.QueryRow(genericQuery, Id, videoId)
	if err := chk.Scan(&resultId); err == sql.ErrNoRows {
		fmt.Println("No Rows found for", idContext, "Id", Id)
	}

	return resultId
}

func filesCheck(r repository, fileType string, videoId int, filesCheckQuery string) int {
	resultId := -1

	//Check if entry for Id exists?
	chk := r.db.QueryRow(filesCheckQuery, fileType, videoId)
	if err := chk.Scan(&resultId); err == sql.ErrNoRows {
		fmt.Println("No Rows found for FileType:", fileType, "VideoId:", videoId)
	}

	return resultId
}

func subsFilesCheck(r repository, fileType string, videoId int, filename string, filesCheckQuery string) int {
	resultId := -1

	//Check if entry for Id exists?
	chk := r.db.QueryRow(filesCheckQuery, fileType, videoId, filename)
	if err := chk.Scan(&resultId); err == sql.ErrNoRows {
		fmt.Println("No Rows found for FileType:", fileType, "VideoId:", videoId, "FileName:", filename)
	}

	return resultId
}
