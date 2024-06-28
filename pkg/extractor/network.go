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
	"time"

	"github.com/gofiber/fiber/v2/log"
	c "github.com/rs-anantmishra/metubeplus/config"
	e "github.com/rs-anantmishra/metubeplus/pkg/entities"
)

type IDownload interface {
	ExtractMetadata() []e.MediaInformation
	ExtractMedia() bool
	ExtractThumbnail([]e.MediaInformation) []e.Files
	ExtractSubtitles() bool
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

func (d *download) ExtractMedia() bool {
	return false
}

func (d *download) ExtractThumbnail(m []e.MediaInformation) []e.Files {

	args, command := cmdBuilderThumbnails(d.p.Indicator, d.indicatorType)
	logCommand := command + Space + args

	//log executed command - in activity log later
	fmt.Println(logCommand)

	cmd, stdout := buildProcess(args, GetCommandString())

	err := cmd.Start()
	handleErrors(err, "Metadata - Cmd.Start")

	pResult := executeProcess(stdout)

	_, errors, results := stripResultSections(pResult)
	fmt.Println(results)

	if len(errors) > 0 {
		//Show error on UI
		log.Error(errors)
		return []e.Files{}
	}

	var fp string
	var fn string
	if d.indicatorType == Video {
		//%(playlist_index)s - %(title)s [%(id)s].%(ext)s"`
		fn = m[0].Title + Space + "[" + m[0].YoutubeVideoId + "]" + "." + m[0].Extension
		fp = strings.Join([]string{c.Config("MEDIA_PATH"), m[0].Domain, m[0].Channel, "Videos", "Thumbnails"}, "\\")
	} else if d.indicatorType == Playlist {
		fn = string(m[0].PlaylistIndex) + Space + "-" + Space + m[0].PlaylistTitle + Space + "[" + m[0].PlaylistId + "]" + "." + m[0].Extension
		fp = strings.Join([]string{c.Config("MEDIA_PATH"), m[0].Domain, m[0].Channel, m[0].PlaylistTitle, "Thumbnails"}, "\\")
	}

	c, err := os.ReadDir(fp)
	handleErrors(err, "network - ExtractThumbnail")

	for _, entry := range c {
		info, _ := entry.Info()
		info.Size()
		info.Name()
	}

	f := e.Files{FileTypeId: e.Thumbnail,
		SourceId:        e.Downloaded,
		FilePath:        fp,
		FileName:        fn,
		Extension:       "",
		FileSize:        0,
		FileSizeUnit:    "bytes",
		ParentDirectory: "",
		IsDeleted:       0,
		CreatedDate:     time.Now().Unix()}

	_ = f

	return []e.Files{}
}

func (d *download) ExtractSubtitles() bool {
	return false
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

	mediaInfo := e.MediaInformation{}
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
		for i := 0; i < metaItemsCount; i++ {
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
