package extractor

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2/log"
	c "github.com/rs-anantmishra/metubeplus/config"
	e "github.com/rs-anantmishra/metubeplus/pkg/entities"
)

type IDownload interface {
	ExtractMetadata() []e.MediaInformation
	ExtractMediaContent() bool
	ExtractThumbnail([]e.MediaInformation) []e.Files
	ExtractSubtitles() []e.Files
}

type download struct {
	p             e.IncomingRequest
	indicatorType int
}

func InstantiateDownload(params e.IncomingRequest) IDownload {
	return &download{
		p: params,
	}
}

func (d *download) ExtractMetadata() []e.MediaInformation {

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

	return mediaInfo
}

func (d *download) ExtractMediaContent() bool {
	return false
}

func (d *download) ExtractThumbnail(m []e.MediaInformation) []e.Files {

	args, command := cmdBuilderThumbnails(d.p.Indicator, d.indicatorType)
	logCommand := command + Space + args

	//log executed command - in activity log later
	_ = logCommand
	cmd, stdout := buildProcess(args, GetCommandString())

	err := cmd.Start()
	handleErrors(err, "Metadata - Cmd.Start")

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
	var fn []string

	//names dont require extensions as they are only needed to match name with subtitle filename.
	for i, elem := range m {
		if d.indicatorType == Video {
			fn = append(fn, elem.Title+Space+"["+elem.YoutubeVideoId+"]")
			fp = strings.Join([]string{c.Config("MEDIA_PATH"), elem.Domain, elem.Channel, "Videos", "Thumbnails"}, "\\")
		} else if d.indicatorType == Playlist {
			if i == 0 {
				//add playlist details
				fn = append(fn, strconv.FormatInt(int64(elem.PlaylistIndex)-1, 10)+Space+"-"+Space+elem.PlaylistTitle+Space+"["+elem.PlaylistId+"]")
				fp = strings.Join([]string{c.Config("MEDIA_PATH"), elem.Domain, elem.Channel, elem.PlaylistTitle, "Thumbnails"}, "\\")

				//add Video details at index 0
				fn = append(fn, strconv.FormatInt(int64(elem.PlaylistIndex), 10)+Space+"-"+Space+elem.Title+Space+"["+elem.YoutubeVideoId+"]")
			} else {
				fn = append(fn, strconv.FormatInt(int64(elem.PlaylistIndex), 10)+Space+"-"+Space+elem.Title+Space+"["+elem.YoutubeVideoId+"]")
			}
		}
	}
	c, err := os.ReadDir(fp)
	handleErrors(err, "network - ExtractThumbnail")

	var files []e.Files
	files = append(files, e.Files{FileTypeId: e.Thumbnail,
		SourceId:     e.Downloaded,
		FilePath:     fp,
		FileName:     fn[0],
		Extension:    "",
		FileSize:     0,
		FileSizeUnit: "bytes",
		NetworkPath:  "",
		IsDeleted:    0,
		CreatedDate:  time.Now().Unix()})

	if d.indicatorType == Playlist {
		for i := 1; i < len(c); i++ {
			file := e.Files{FileTypeId: e.Thumbnail,
				SourceId:     e.Downloaded,
				FilePath:     fp,
				FileName:     fn[i],
				Extension:    "",
				FileSize:     0,
				FileSizeUnit: "bytes",
				NetworkPath:  "",
				IsDeleted:    0,
				CreatedDate:  time.Now().Unix()}

			files = append(files, file)
		}
	}

	//Todo: re-write this to compare filenames after removing all special characters.
	//if there is a match then do the assignment.
	for k, entry := range c {
		info, _ := entry.Info()
		files[k].FileName = info.Name()
		files[k].FileSize = int(info.Size())
		files[k].CreatedDate = info.ModTime().Unix()
		splits := strings.SplitN(info.Name(), ".", -1)
		files[k].Extension = splits[len(splits)-1]
	}

	return files
}

func (d *download) ExtractSubtitles() []e.Files {
	return []e.Files{}
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

func sanitizeResults(b bytes.Buffer) []string {

	//sanitize json
	result := strings.ReplaceAll(b.String(), "'", "\"")

	//split by newlines and remove and from the end.
	results := strings.Split(result, "\n")
	if results[len(results)-1] == "" {
		results = results[:len(results)-1]
	}
	return results
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
			json.Unmarshal([]byte(results[i]), &mediaInfo)
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

	for index, elem := range pResult {
		if val := strings.Index(elem, WARNING); val == 0 {
			warnings = append(warnings, elem)
		} else if val := strings.Index(elem, ERROR); val == 0 {
			errors = append(errors, elem)
		} else {
			results = pResult[index:]
			break
		}
	}

	return warnings, errors, results
}

func handleErrors(err error, methodName string) {
	if err != nil {
		fmt.Println("pkg dowonload", methodName, err)
	}
}

// Result Type for entity binding and result parsing
const (
	Indicator        = iota
	VideoMetadata    = iota
	PlaylistMetadata = iota
	Download         = iota
)
