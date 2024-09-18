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

// Service is an interface from which our api module can access our repository of all our models
type IService interface {
	GetVideos() ([]presenter.LimitedCardsInfoResponse, error)

	//TODO:
	//InsertVideo(book *entities.Videos) (*entities.Videos, error)
	//UpdateVideo(book *entities.Videos) (*entities.Videos, error)
	//RemoveVideo(ID string) error
}

type service struct {
	repository IRepository
}

func NewVideoService(r IRepository) IService {
	return &service{
		repository: r,
	}
}

// GetVideos implements IService.
func (s *service) GetVideos() ([]presenter.LimitedCardsInfoResponse, error) {
	allVideos, err := s.repository.GetAllVideos()
	if err != nil {
		return nil, err
	}
	result := GetVideosPageInfo(allVideos)

	return result, nil
}

func GetVideosPageInfo(videos []entities.Videos) []presenter.LimitedCardsInfoResponse {

	var lstCardsInfo []presenter.LimitedCardsInfoResponse
	for _, elem := range videos {
		var cardsInfo presenter.LimitedCardsInfoResponse

		cardsInfo.VideoId = elem.Id
		cardsInfo.Title = elem.Title
		cardsInfo.Channel = elem.Channel.Name
		cardsInfo.Description = elem.Description
		cardsInfo.Duration = elem.DurationSeconds
		cardsInfo.OriginalURL = elem.WebpageURL
		cardsInfo.Thumbnail = getImagesFromURL(elem.ThumbnailFilePath)
		cardsInfo.VideoFilepath = strings.Replace(elem.VideoFilePath, extractor.GetMediaDirectory(false), "http://localhost:3000/", -1)
		cardsInfo.VideoFilepath = strings.Replace(elem.VideoFilePath, "\\", "/", -1)

		lstCardsInfo = append(lstCardsInfo, cardsInfo)
	}

	return lstCardsInfo
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
