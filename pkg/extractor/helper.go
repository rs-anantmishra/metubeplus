package extractor

import (
	"encoding/base64"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2/log"
	p "github.com/rs-anantmishra/metubeplus/api/presenter"
	e "github.com/rs-anantmishra/metubeplus/pkg/entities"
	g "github.com/rs-anantmishra/metubeplus/pkg/global"
)

func falsifyQueueAlive() {
	qa := g.NewQueueAlive()
	qa[0] = 0
}

func clearActiveItem(activeItem []g.DownloadStatus) []g.DownloadStatus {

	//empty out active item
	activeItem[0].VideoURL = ""
	activeItem[0].VideoId = 0
	activeItem[0].StatusMessage = ""
	activeItem[0].State = 0

	return activeItem
}

func createMetadataResponse(lstSavedInfo []e.SavedInfo, subtitles []e.Files, subtitlesReq bool, thumbnails []e.Files) []p.CardsInfoResponse {
	//bind here to presenter entity
	var cardMetaDataInfoList []p.CardsInfoResponse
	const _blank string = ""

	for _, elem := range lstSavedInfo {
		var cardMetaDataInfo p.CardsInfoResponse

		cardMetaDataInfo.Channel = elem.MediaInfo.Channel
		cardMetaDataInfo.CreatedDate = int(time.Now().Unix())
		cardMetaDataInfo.Description = elem.MediaInfo.Description
		cardMetaDataInfo.Domain = elem.MediaInfo.Domain
		cardMetaDataInfo.Duration = elem.MediaInfo.Duration
		cardMetaDataInfo.IsDeleted = false
		cardMetaDataInfo.IsFileDownloaded = false
		cardMetaDataInfo.MediaURL = _blank
		cardMetaDataInfo.OriginalURL = elem.MediaInfo.OriginalURL
		cardMetaDataInfo.Playlist = elem.MediaInfo.PlaylistTitle
		cardMetaDataInfo.PlaylistVideoIndex = elem.MediaInfo.PlaylistIndex
		cardMetaDataInfo.Title = elem.MediaInfo.Title
		cardMetaDataInfo.VideoFormat = elem.MediaInfo.Format
		cardMetaDataInfo.VideoId = elem.VideoId
		cardMetaDataInfo.WatchCount = 0

		cardMetaDataInfoList = append(cardMetaDataInfoList, cardMetaDataInfo)
	}

	//subtitles
	if subtitlesReq {
		for i, elem := range subtitles {
			cardMetaDataInfoList[i].SubtitlesURL = elem.FilePath + elem.FileName
		}
	}

	//thumbnails
	for i := range thumbnails {
		if lstSavedInfo[0].MediaInfo.PlaylistTitle == _blank {
			cardMetaDataInfoList[i].Thumbnail = getImagesFromURL(thumbnails[i])
		} else {
			cardMetaDataInfoList[i].Thumbnail = getImagesFromURL(thumbnails[i])
			//set playlist thumbnail whenever Video Index is 1
			if cardMetaDataInfoList[i].PlaylistVideoIndex == 1 {
				playlistThumbnail := getImagesFromURL(thumbnails[i])
				for k := 0; k < len(thumbnails); k++ {
					cardMetaDataInfoList[k].PlaylistThumbnail = playlistThumbnail
				}
			}
		}
	}

	return cardMetaDataInfoList
}

func getImagesFromURL(file e.Files) string {
	var base64EncodedImage string

	filepath := file.FilePath + "\\" + file.FileName
	// Read the entire file into a byte slice
	bytes, err := os.ReadFile(filepath)
	if err != nil {
		log.Info(err)
	}

	switch file.Extension {
	case "jpeg":
		base64EncodedImage += "data:image/jpeg;base64,"
	case "jpg":
		base64EncodedImage += "data:image/jpg;base64,"
	case "png":
		base64EncodedImage += "data:image/png;base64,"
	case "webp":
		base64EncodedImage += "data:image/webp;base64,"
	}

	base64EncodedImage += base64.StdEncoding.EncodeToString(bytes)
	return base64EncodedImage
}

func getImagesFromURLString(filepath string) string {
	var base64EncodedImage string
	splitter := "."

	// Read the entire file into a byte slice
	bytes, err := os.ReadFile(filepath)
	if err != nil {
		log.Info(err)
	}

	if len(bytes) == 0 {
		filepath = `..\utils\noimage.png`
		bytes, err = os.ReadFile(filepath)
		if err != nil {
			log.Info(err)
		}
	}

	splits := strings.SplitN(filepath, splitter, -1)
	extension := splits[len(splits)-1]

	switch extension {
	case "jpeg":
		base64EncodedImage += "data:image/jpeg;base64,"
	case "jpg":
		base64EncodedImage += "data:image/jpg;base64,"
	case "png":
		base64EncodedImage += "data:image/png;base64,"
	case "webp":
		base64EncodedImage += "data:image/webp;base64,"
	}

	base64EncodedImage += base64.StdEncoding.EncodeToString(bytes)
	return base64EncodedImage
}

// Remove Characters that are not allowed in folder-names
// Only specific fields to be handled which are used for generating folder names
// Folder naming rules:
// Windows (FAT32, NTFS): Any Unicode except NUL, \, /, :, *, ?, ", <, >, |. Also, no space character at the start or end, and no period at the end.
// Mac(HFS, HFS+): Any valid Unicode except : or /
// Linux(ext[2-4]): Any byte except NUL or /
func removeForbiddenChars(metadata []e.MediaInformation) []e.MediaInformation {

	// / = 10744

	forbiddenChars := []string{"\\", "/", ":", "*", "?", "\"", "<", ">", "|"}
	emptyString := ""
	singleSpace := " "
	doubleSpaces := "  "

	for i := 0; i < len(metadata); i++ {
		for _, elem := range forbiddenChars {

			//handle Domain
			if strings.Contains(metadata[i].Domain, elem) {
				metadata[i].Domain = strings.ReplaceAll(metadata[i].Domain, elem, emptyString)
				metadata[i].Domain = strings.TrimSpace(metadata[i].Domain)                             //Trim leading and trailing spaces
				metadata[i].Domain = strings.TrimRight(metadata[i].Domain, ".")                        //Trim trailing period
				metadata[i].Domain = strings.ReplaceAll(metadata[i].Domain, doubleSpaces, singleSpace) //Replace any double spaces that may have occurred as a result of removing characters
			}

			//handle Video Channel
			if strings.Contains(metadata[i].Channel, elem) {
				metadata[i].Channel = strings.ReplaceAll(metadata[i].Channel, elem, emptyString)
				metadata[i].Channel = strings.TrimSpace(metadata[i].Channel)
				metadata[i].Channel = strings.TrimRight(metadata[i].Channel, ".")
				metadata[i].Channel = strings.ReplaceAll(metadata[i].Channel, doubleSpaces, singleSpace)
			}

			//handle Video Title
			if strings.Contains(metadata[i].Title, elem) {
				metadata[i].Title = strings.ReplaceAll(metadata[i].Title, elem, emptyString)
				metadata[i].Title = strings.TrimSpace(metadata[i].Title)
				metadata[i].Title = strings.TrimRight(metadata[i].Title, ".")
				metadata[i].Title = strings.ReplaceAll(metadata[i].Title, doubleSpaces, singleSpace)
			}

			if strings.Contains(metadata[i].PlaylistTitle, elem) {
				metadata[i].PlaylistTitle = strings.ReplaceAll(metadata[i].PlaylistTitle, elem, emptyString)
				metadata[i].PlaylistTitle = strings.TrimSpace(metadata[i].PlaylistTitle)
				metadata[i].PlaylistTitle = strings.TrimRight(metadata[i].PlaylistTitle, ".")
				metadata[i].PlaylistTitle = strings.ReplaceAll(metadata[i].PlaylistTitle, doubleSpaces, singleSpace)
			}
		}
	}

	return metadata
}

func getFilepaths(playlistId int, fPath e.Filepath, pathType int) string {
	var fp string

	if playlistId < 0 {
		fp = GetVideoFilepath(fPath, pathType)
	} else if playlistId > 0 {
		fp = GetPlaylistFilepath(fPath, pathType)
	}

	return fp
}

func handleSingleChannelPlaylist(lstSavedInfo []e.MediaInformation) bool {

	//handle video
	if len(lstSavedInfo) == 1 && lstSavedInfo[0].PlaylistId == "" {
		return true
	}

	result := true
	//private playlists will have videos from various channels
	prevChannel := ""
	currentChannel := ""
	for chIndex := 1; chIndex < len(lstSavedInfo); chIndex++ {
		prevChannel = lstSavedInfo[chIndex-1].Channel
		currentChannel = lstSavedInfo[chIndex].Channel

		if prevChannel != currentChannel {
			result = false
		}
	}

	return result
}
