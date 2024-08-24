package extractor

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"syscall"

	"github.com/gofiber/fiber/v2/log"
	c "github.com/rs-anantmishra/metubeplus/config"
	e "github.com/rs-anantmishra/metubeplus/pkg/entities"
	g "github.com/rs-anantmishra/metubeplus/pkg/global"
)

type IDownload interface {
	ExtractMetadata() ([]e.MediaInformation, e.Filepath)
	ExtractMediaContent() int
	ExtractThumbnail(fp e.Filepath, videoId []int, lstSMI []e.SavedMediaInformation) []e.Files
	ExtractSubtitles(fp e.Filepath, videoId []int, lstSMI []e.SavedMediaInformation) []e.Files
	GetDownloadedMediaFileInfo(smi e.SavedMediaInformation, fp e.Filepath) []e.Files
	Cleanup()
}

type download struct {
	p             e.IncomingRequest
	indicatorType int
	lstDownloads  []g.DownloadStatus
}

func NewDownload(params e.IncomingRequest) IDownload {
	return &download{
		p: params,
	}
}

func (d *download) Cleanup() {
	var updated []g.DownloadStatus
	for i := 0; i < len(d.lstDownloads); i++ {
		ds := d.lstDownloads
		if ds[i].State != g.Completed {
			updated = append(updated, ds[i])
		}
	}
	d.lstDownloads = updated
}

func (d *download) ExtractMetadata() ([]e.MediaInformation, e.Filepath) {

	indicatorType, itemCount := getIndicatorType(d.p.Indicator)
	d.indicatorType = indicatorType

	args, command := cmdBuilderMetadata(d.p.Indicator, indicatorType)
	logCommand := command + Space + args

	//log executed command - in activity log later
	fmt.Println(logCommand)

	cmd, stdout := buildProcess(args, GetCommandString())

	err := cmd.Start()
	handleErrors(err, "Metadata - Cmd.Start")

	var pResult []string
	var mediaInfo []e.MediaInformation
	if indicatorType == Video {
		pResult = executeProcess(stdout)
		video := parseResults(pResult, VideoMetadata, itemCount)

		//dump data to db and return result from here
		fmt.Println(video)
		mediaInfo = video
	}

	if indicatorType == Playlist {
		pResult = executeProcess(stdout)
		playlist := parseResults(pResult, PlaylistMetadata, itemCount)

		//dump data to db and return result from here
		fmt.Println(playlist)
		mediaInfo = playlist
	}

	fp := e.Filepath{Domain: mediaInfo[0].Domain, Channel: mediaInfo[0].Channel, PlaylistTitle: mediaInfo[0].PlaylistTitle}

	return mediaInfo, fp
}

func (d *download) ExtractMediaContent() int {

	activeItem := g.NewActiveItem()

	args, command := cmdBuilderDownload(activeItem[0].VideoURL, Video)
	logCommand := command + Space + args

	//log executed command - in activity log later
	_ = logCommand
	cmd, stdout := buildProcess(args, GetCommandString())

	err := cmd.Start()
	handleErrors(err, "Download - Cmd.Start")

	pResult := executeDownloadProcess(stdout, activeItem)
	_, errors, results := stripResultSections(pResult)

	//results are not really needed - except maybe to check for errors.
	_ = errors
	_ = results

	return activeItem[0].State

}

func (d *download) ExtractThumbnail(fPath e.Filepath, videoId []int, lstSMI []e.SavedMediaInformation) []e.Files {

	args, command := cmdBuilderThumbnails(d.p.Indicator, d.indicatorType)
	logCommand := command + Space + args

	//log executed command - in activity log later
	_ = logCommand
	cmd, stdout := buildProcess(args, GetCommandString())

	err := cmd.Start()
	handleErrors(err, "Thumbnail - Cmd.Start")

	pResult := executeProcess(stdout)
	_, errors, results := stripResultSections(pResult)

	//results are not really needed - except maybe to check for errors.
	_ = errors
	_ = results

	if len(errors) > 0 {
		//Show error on UI
		log.Error(errors)
		return []e.Files{}
	}

	var fp string
	//Get FilePaths
	if d.indicatorType == Video {
		fp = GetVideoFilepath(fPath, e.Thumbnail)
	} else if d.indicatorType == Playlist {
		fp = GetPlaylistFilepath(fPath, e.Thumbnail)
	}
	fp = strings.ReplaceAll(fp, "../media/", "..\\media")

	c, err := os.ReadDir(fp)
	handleErrors(err, "network - ExtractThumbnail")

	var files []e.Files
	//Todo: re-write this to compare filenames after removing all special characters.
	//if there is a match then do the assignment.
	for i, entry := range c {

		info, _ := entry.Info()
		splits := strings.SplitN(info.Name(), ".", -1)
		fs_filename := info.Name()

		thumbnailVideoId := -1
		thumbnailPlaylistId := lstSMI[0].PlaylistId
		//for playlists, 1st Video Id should be -1
		if d.indicatorType == Playlist {
			if i > 0 {
				thumbnailVideoId = lstSMI[i-1].VideoId
			}
			//if VideoId doesn't exist for this thumbnail
			//can include logic to remove thumbnail too?
			if i > 0 && lstSMI[i-1].VideoId < 0 {
				continue
			}

			f := e.Files{
				VideoId:      thumbnailVideoId,
				PlaylistId:   thumbnailPlaylistId,
				FileType:     "Thumbnail",
				SourceId:     e.Downloaded,
				FilePath:     fp,
				FileName:     info.Name(),
				Extension:    splits[len(splits)-1],
				FileSize:     int(info.Size()),
				FileSizeUnit: "bytes",
				NetworkPath:  "",
				IsDeleted:    0,
				CreatedDate:  info.ModTime().Unix(),
			}

			files = append(files, f)
		} else if d.indicatorType == Video {
			smiIndex := 0
			thumbnailVideoId = lstSMI[smiIndex].VideoId
			for _, saved := range lstSMI {
				if strings.Contains(fs_filename, saved.YoutubeVideoId) {
					f := e.Files{
						VideoId:      thumbnailVideoId,
						PlaylistId:   thumbnailPlaylistId,
						FileType:     "Thumbnail",
						SourceId:     e.Downloaded,
						FilePath:     fp,
						FileName:     info.Name(),
						Extension:    splits[len(splits)-1],
						FileSize:     int(info.Size()),
						FileSizeUnit: "bytes",
						NetworkPath:  "",
						IsDeleted:    0,
						CreatedDate:  info.ModTime().Unix(),
					}
					files = append(files, f)
				}
			}
		}
	}
	return files
}

func (d *download) ExtractSubtitles(fPath e.Filepath, videoId []int, lstSMI []e.SavedMediaInformation) []e.Files {

	args, command := cmdBuilderSubtitles(d.p.Indicator, d.indicatorType)
	logCommand := command + Space + args

	//log executed command - in activity log later
	_ = logCommand
	cmd, stdout := buildProcess(args, GetCommandString())

	err := cmd.Start()
	handleErrors(err, "Subtitles - Cmd.Start")

	pResult := executeProcess(stdout)
	_, errors, results := stripResultSections(pResult)

	//results are not really needed - except maybe to check for errors.
	_ = errors
	_ = results

	if len(errors) > 0 {
		//Show error on UI
		log.Error(errors)
		return []e.Files{}
	}

	var fp string
	//Get FilePaths
	if d.indicatorType == Video {
		fp = GetVideoFilepath(fPath, e.Subtitles)
	} else if d.indicatorType == Playlist {
		fp = GetPlaylistFilepath(fPath, e.Subtitles)
	}
	fp = strings.ReplaceAll(fp, "../media/", "..\\media")

	c, err := os.ReadDir(fp)
	handleErrors(err, "network - ExtractSubtitles")

	smiIndex := 0
	var files []e.Files
	//Todo: re-write this to compare filenames after removing all special characters.
	//if there is a match then do the assignment.
	for _, entry := range c {
		info, _ := entry.Info()
		splits := strings.SplitN(info.Name(), ".", -1)
		fs_filename := info.Name()

		for _, saved := range lstSMI {
			if strings.Contains(fs_filename, saved.YoutubeVideoId) {
				f := e.Files{
					VideoId:      lstSMI[smiIndex].VideoId,
					PlaylistId:   lstSMI[smiIndex].PlaylistId,
					FileType:     "Subtitles",
					SourceId:     e.Downloaded,
					FilePath:     fp,
					FileName:     info.Name(),
					Extension:    splits[len(splits)-1],
					FileSize:     int(info.Size()),
					FileSizeUnit: "bytes",
					NetworkPath:  "",
					IsDeleted:    0,
					CreatedDate:  info.ModTime().Unix(),
				}
				files = append(files, f)
				smiIndex++
			}
		}
	}
	return files
}

func (d *download) GetDownloadedMediaFileInfo(smi e.SavedMediaInformation, fPath e.Filepath) []e.Files {

	var fp string
	//Get FilePaths
	if smi.PlaylistId == -1 {
		fp = GetVideoFilepath(fPath, e.Video)
	} else if smi.PlaylistId > -1 {
		fp = GetPlaylistFilepath(fPath, e.Video)
	}
	fp = strings.ReplaceAll(fp, "../media/", "..\\media")

	c, err := os.ReadDir(fp)
	handleErrors(err, "network - ExtractMediaContent")

	smiIndex := 0
	var files []e.Files
	//Todo: re-write this to compare filenames after removing all special characters.
	//if there is a match then do the assignment.
	for _, entry := range c {
		info, _ := entry.Info()
		splits := strings.SplitN(info.Name(), ".", -1)
		fs_filename := info.Name()

		if strings.Contains(fs_filename, smi.YoutubeVideoId) {
			f := e.Files{
				VideoId:      smi.VideoId,
				PlaylistId:   smi.PlaylistId,
				FileType:     "Video",
				SourceId:     e.Downloaded,
				FilePath:     fp,
				FileName:     info.Name(),
				Extension:    splits[len(splits)-1],
				FileSize:     int(info.Size()),
				FileSizeUnit: "bytes",
				NetworkPath:  "",
				IsDeleted:    0,
				CreatedDate:  info.ModTime().Unix(),
			}
			files = append(files, f)
			smiIndex++
		}
	}

	return files
}

func getIndicatorType(url string) (int, int) {

	arguments := "\"" + url + "\"" + Space + Title + Space + SkipDownload
	cmd, stdout := buildProcess(arguments, GetCommandString())

	err := cmd.Start()
	handleErrors(err, "ValidateRequest - Cmd.Start")

	pResult := executeProcess(stdout)
	_, _, results := stripResultSections(pResult)

	switch {
	case len(results) == 1:
		return Video, len(results)
	case len(results) > 1:
		return Playlist, len(results)
	default:
		return Generic, len(results)
	}
}

// Build a Process to execute & attach pipe to it here
func buildProcess(args string, command string) (*exec.Cmd, io.ReadCloser) {

	cmd := exec.Command(command)
	cmd.SysProcAttr = &syscall.SysProcAttr{}
	cmd.SysProcAttr.CmdLine = command + Space + args

	stdout, err := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout
	handleErrors(err, "Metadata - StdoutPipe")

	return cmd, stdout
}

func executeProcess(stdout io.ReadCloser) []string {
	// var result string
	var b bytes.Buffer
	for {
		//Read data from pipe into temp
		temp := make([]byte, 2048)
		n, err := stdout.Read(temp)

		//handle verbosity - it will always be there (similar to dpkg)
		verbose := true
		if verbose {
			b.WriteString(string(temp[:n]))
		}

		//terminate loop at eof
		if err != nil {
			log.Info("Error Reading:", err)
			break
		}
	}

	//sanitize and return
	results := sanitizeResults(b)
	return results
}

func executeDownloadProcess(stdout io.ReadCloser, activeItem []g.DownloadStatus) []string {

	//Update State
	activeItem[0].State = g.Downloading

	// var result string
	var b bytes.Buffer
	for {
		//Read data from pipe into temp
		temp := make([]byte, 2048)
		n, err := stdout.Read(temp)

		//handle verbosity - it will always be there (similar to dpkg)
		verbose := true
		if verbose {
			b.WriteString(string(temp[:n]))
			result := b.String()
			results := strings.Split(result, "\n")

			if len(results)-2 >= 0 {
				activeItem[0].StatusMessage = results[len(results)-2]
				log.Info("MESSAGE VALUE: ", activeItem[0].StatusMessage)
			}
		}

		//terminate loop at eof
		if err != nil {
			log.Info("Error Reading:", err)
			if err == io.EOF {
				activeItem[0].StatusMessage = "Download completed successfully."
			} else {
				activeItem[0].StatusMessage = "Error:" + err.Error()
			}
			break
		}
	}
	//change state to stop reprocessing.
	activeItem[0].State = g.Completed

	result := b.String()
	results := strings.Split(result, "\n")

	return results
}

func sanitizeResults(b bytes.Buffer) []string {

	//split result by newlines
	result := b.String()
	results := strings.Split(result, "\n")

	for i := range results {
		//valid json require keys and values to be enclosed in double quotes, not single quotes
		results[i] = proximityQuoteReplacement(results[i])
	}

	//remove newlines from the end
	if results[len(results)-1] == "" {
		results = results[:len(results)-1]
	}
	return results
}

// valid json require keys and values to be enclosed in double quotes, not single quotes
func proximityQuoteReplacement(data string) string {

	dQ := []byte("\"")[0]
	b := []byte(data)

	seqArraryCheck1 := strings.Index(data, ": ['")
	seqArraryCheck2 := strings.LastIndex(data, "']")
	if seqArraryCheck1 >= 0 && seqArraryCheck2 >= 0 {
		data = strings.ReplaceAll(data, "'", "\"")
		return data
	}

	if seq1 := strings.Index(data, "{'"); seq1 >= 0 {
		b[seq1+1] = dQ
	}

	if seq2 := strings.Index(data, "':"); seq2 >= 0 {
		b[seq2] = dQ
	}

	if seq3 := strings.Index(data, ": '"); seq3 >= 0 {
		b[seq3+2] = dQ
	}

	if seq4 := strings.LastIndex(data, "'}"); seq4 >= 0 {
		b[seq4] = dQ
	}

	data = string(b)
	return data
}

func parseResults(pResult []string, metadataType int, vCount int) []e.MediaInformation {

	_, _, results := stripResultSections(pResult)

	metaItemsCount := 0
	for _, elem := range BuilderOptions() {
		if metadataType == VideoMetadata && elem.Group.Video.Metadata && elem.DataField {
			metaItemsCount++
		} else if metadataType == PlaylistMetadata && elem.Group.Playlist.Metadata && elem.DataField {
			metaItemsCount++
		}
	}

	var lstMediaInfo []e.MediaInformation
	for k := 0; k < vCount; k++ {
		mediaInfo := e.MediaInformation{}
		for i := (0 + k*metaItemsCount); i < (k+1)*metaItemsCount; i++ {

			//Unmarshall is unreliable since the json coming from yt-dlp is invalid.
			if results[i][0] == '{' && results[i][len(results[i])-1] == '}' {
				json.Unmarshal([]byte(results[i]), &mediaInfo)
			}
		}

		if c.Config("PATCHING") == "enabled" {
			mediaInfo = patchDataField(mediaInfo)
		}

		lstMediaInfo = append(lstMediaInfo, mediaInfo)
	}

	//Print Properties that were bound
	fmt.Println(lstMediaInfo)

	return lstMediaInfo
}

func stripResultSections(pResult []string) ([]string, []string, []string) {

	var warnings []string
	var errors []string
	var results []string
	var previous string

	for _, elem := range pResult {
		if val := strings.Index(elem, WARNING); val == 0 {
			warnings = append(warnings, elem)
			previous = WARNING
		} else if val := strings.Index(elem, ERROR); val == 0 {
			errors = append(errors, elem)
			previous = ERROR
		} else if val := strings.Index(elem, ANSWER_START); val == 0 {
			results = append(results, elem)
		} else {
			//append to previous entry if nothing matches -- most tested and stable solution
			if previous == WARNING {
				warnings = append(warnings, elem)
			} else if previous == ERROR {
				errors = append(errors, elem)
			}
		}

	}

	return warnings, errors, results
}

func handleErrors(err error, methodName string) {
	if err != nil {
		fmt.Println("pkg dowonload", methodName, err)
	}
}

// avoiding reflection here
func patchDataField(mediaInfo e.MediaInformation) e.MediaInformation {

	const plainChannel string = "Channel: "
	const plainTitle string = "Title: "
	const plainDescription string = "Description: "
	const plainTags string = "Tags: "
	const plainCategories string = "Categories: "

	var queries []string
	//Only checking for title description, Channel Name errors here.
	switch {
	case mediaInfo.Channel == "":
		queries = append(queries, Plaintext_Channel)
	case mediaInfo.Title == "":
		queries = append(queries, Plaintext_Title)
	case mediaInfo.Description == "":
		queries = append(queries, Plaintext_Description)
	case len(mediaInfo.Tags) == 0:
		queries = append(queries, Plaintext_Tags)
	case len(mediaInfo.Categories) == 0:
		queries = append(queries, Plaintext_Categories)
	default:
		break
	}

	for _, elem := range queries {
		var args []string
		args = append(args, mediaInfo.OriginalURL)
		args = append(args, SkipDownload)
		args = append(args, elem)

		options := strings.Join(args, Space)
		cmd, stdout := buildProcess(options, GetCommandString())

		err := cmd.Start()
		handleErrors(err, "patchDataField - Cmd.Start")

		procResult := executeProcess(stdout)

		for i := range procResult {
			if idx := strings.Index(procResult[i], plainChannel); idx == 0 {
				mediaInfo.Channel = procResult[i][len(plainChannel):]
			} else if idx := strings.Index(procResult[i], plainTitle); idx == 0 {
				mediaInfo.Title = procResult[i][len(plainTitle):]
			} else if idx := strings.Index(procResult[i], plainDescription); idx == 0 {
				mediaInfo.Description = procResult[i][len(plainDescription):]
			} else if idx := strings.Index(procResult[i], plainTags); idx == 0 {
				mediaInfo.Description = procResult[i][len(plainTags):]
			} else if idx := strings.Index(procResult[i], plainCategories); idx == 0 {
				mediaInfo.Description = procResult[i][len(plainCategories):]
			}
		}
	}

	return mediaInfo
}

// Result Type for entity binding and result parsing
const (
	Indicator        = iota
	VideoMetadata    = iota
	PlaylistMetadata = iota
	Download         = iota
)
