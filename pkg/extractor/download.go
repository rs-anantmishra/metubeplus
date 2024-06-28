package extractor

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"os/exec"
	"strings"
	"syscall"

	"github.com/gofiber/fiber/v2/log"
	e "github.com/rs-anantmishra/metubeplus/pkg/entities"
)

type IDownload interface {
	ExtractMetadata(verbose bool) []e.MediaInformation
	ExtractVideo()
}

type download struct {
	url string
}

func InstantiateDownload(URL string) IDownload {
	return &download{
		url: URL,
	}
}

func (d *download) ExtractMetadata(verbose bool) []e.MediaInformation {

	indicatorType, itemCount := getIndicatorType(d.url)

	//in case of errors just terminate the flow and return the error to UI
	fmt.Println(indicatorType)
	//in case of errors just terminate the flow and return the error to UI

	args, command := cmdBuilderMetadata(d.url, indicatorType, false)
	logCommand := command + Space + args

	//log executed command
	fmt.Println(logCommand)

	cmd, stdout := buildProcess(args, GetCommandString())
	//fmt.Println(cmd.String())

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

func (d *download) ExtractVideo() {

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

// [deprecated] if string list in URL then its a playlist.
func EvaluateFxGroup(link string) int {

	const playlistKey string = "list"
	result := Video

	u, err := url.Parse(link)
	handleErrors(err, "EvaluateFxGroup")

	q := u.Query()
	_, ok := q[playlistKey]
	if ok {
		result = Playlist
	}

	return result
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
			// result = (result, string(temp[:n]))
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
