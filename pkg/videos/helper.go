package videos

import (
	"encoding/base64"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2/log"
	"github.com/rs-anantmishra/metubeplus/api/presenter"
	"github.com/rs-anantmishra/metubeplus/pkg/entities"
	"github.com/rs-anantmishra/metubeplus/pkg/extractor"
)

func getVideosPageInfo(videos []entities.Videos) []presenter.CardsInfoResponse {

	var lstCardsInfo []presenter.CardsInfoResponse
	for _, elem := range videos {
		var cardsInfo presenter.CardsInfoResponse

		cardsInfo.VideoId = elem.Id
		cardsInfo.Title = elem.Title
		cardsInfo.Description = elem.Description
		cardsInfo.Duration = elem.DurationSeconds
		cardsInfo.WebpageURL = elem.WebpageURL
		cardsInfo.Channel = elem.Channel.Name
		cardsInfo.Domain = elem.Domain.Domain
		cardsInfo.VideoFormat = elem.Format.Format
		cardsInfo.WatchCount = elem.WatchCount
		cardsInfo.ViewsCount = elem.ViewsCount
		cardsInfo.LikesCount = elem.LikesCount
		cardsInfo.UploadDate = elem.UploadDate

		//tags and categories - name only
		cardsInfo.Tags, cardsInfo.Categories = getTagsCategories(elem.Tags, elem.Categories)

		//files - limited info only
		cardsInfo.FileSize, cardsInfo.MediaURL, cardsInfo.Thumbnail, cardsInfo.Extension = getFilesInfo(elem.Files)

		//additional transforms
		cardsInfo.MediaURL = strings.Replace(cardsInfo.MediaURL, extractor.GetMediaDirectory(false), "http://localhost:3000/", -1)
		cardsInfo.MediaURL = strings.Replace(cardsInfo.MediaURL, "\\", "/", -1)

		lstCardsInfo = append(lstCardsInfo, cardsInfo)
	}

	return lstCardsInfo
}

func getFilesInfo(files []entities.Files) (int, string, string, string) {

	filesize := 0
	contentFilepath := ``
	thumbnail := ``
	extension := ``

	for idx := range files {
		if files[idx].FileType == "Video" {
			filesize = files[idx].FileSize
			contentFilepath = files[idx].FilePath + "\\" + files[idx].FileName
			extension = files[idx].Extension
		} else if files[idx].FileType == "Thumbnail" {
			thumbnailURL := files[idx].FilePath + "\\" + files[idx].FileName
			thumbnail = getImagesFromURL(thumbnailURL)
		}
	}

	return filesize, contentFilepath, thumbnail, extension
}

func getTagsCategories(tags []entities.Tags, categories []entities.Categories) ([]string, []string) {
	var resultTags []string
	var resultCategories []string

	for _, elem := range tags {
		resultTags = append(resultTags, elem.Name)
	}

	for _, elem := range categories {
		resultCategories = append(resultCategories, elem.Name)
	}

	return resultTags, resultCategories
}

func getImagesFromURL(filepath string) string {
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

func getContentSearchResponse(list []entities.ContentSearch) []presenter.ContentSearchResponse {
	var result []presenter.ContentSearchResponse
	for idx := range list {
		result = append(result, presenter.ContentSearchResponse{
			VideoId: list[idx].VideoId,
			Title:   list[idx].Title,
			Channel: list[idx].Channel,
		})
	}
	return result
}
