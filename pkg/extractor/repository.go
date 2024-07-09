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

	if len(metadata) < 1 {
		return 0
	}

	//Channel will be same for all items in playlist.
	channelId := genericCheck(*r, metadata[0].ChannelId, "Channel", p.InsertChannelCheck)
	if channelId <= 0 {
		var args []any
		args = append(args, metadata[0].Channel)
		args = append(args, metadata[0].ChannelURL)
		args = append(args, metadata[0].ChannelId)
		args = append(args, time.Now().Unix())

		channelId = genericSave(*r, args, p.InsertChannel)
		_ = channelId
	}

	//playlist info will be same for all in the playlist.
	playlistId := genericCheck(*r, metadata[0].PlaylistId, "Playlist", p.InsertPlaylistCheck)
	if playlistId <= 0 {
		var args []any
		args = append(args, metadata[0].PlaylistTitle)
		args = append(args, metadata[0].PlaylistCount)
		args = append(args, GetPlaylistFilepath(fp, -1))
		args = append(args, metadata[0].ChannelId)
		args = append(args, 0)
		args = append(args, metadata[0].PlaylistId)
		args = append(args, time.Now().Unix())

		playlistId = genericSave(*r, args, p.InsertPlaylist)
		_ = playlistId
	}

	//Domain will be same for all items in playlist.
	domainId := genericCheck(*r, metadata[0].Domain, "Domain", p.InsertDomainCheck)
	if domainId <= 0 {
		var args []any
		args = append(args, metadata[0].Domain)
		args = append(args, time.Now().Unix())

		domainId = genericSave(*r, args, p.InsertDomain)
		_ = domainId
	}

	//Format will NOT be same for all items in playlist.
	formatId := genericCheck(*r, metadata[0].Format, "Format", p.InsertFormatCheck)
	if formatId <= 0 {
		var args []any
		args = append(args, metadata[0].Format)
		args = append(args, metadata[0].FormatNote)
		args = append(args, metadata[0].Resolution)
		args = append(args, "Video") //It should be Audio for audio only files
		args = append(args, time.Now().Unix())

		formatId = genericSave(*r, args, p.InsertFormat)
		_ = formatId
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
