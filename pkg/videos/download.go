package videos

import (
	"bytes"
	"fmt"
	"io"
	"net/url"
	"os/exec"
	"strings"
	"syscall"

	"github.com/gofiber/fiber/v2/log"
	e "github.com/rs-anantmishra/metubeplus/entities"
)

type IDownload interface {
	GetMetadata(verbose bool) e.Metadata
	GetVideo()
}

type download struct {
	url string
}

func InstantiateDownload(URL string) IDownload {
	return &download{
		url: URL,
	}
}

func getIndicatorType(url string) ([]string, []string, int) {

	arguments := "\"" + url + "\"" + Space + Title + Space + SkipDownload
	cmd, stdout := ProcessBuilder(arguments, GetCommandString())

	err := cmd.Start()
	handleErrors(err, "ValidateRequest - Cmd.Start")

	pResult := ProcessResult(stdout)

	var warnings []string
	var errors []string

	for index, elem := range pResult {
		if val := strings.Index(elem, WARNING); val == 0 {
			warnings = append(warnings, elem)
		} else if val := strings.Index(elem, ERROR); val == 0 {
			errors = append(errors, elem)
		} else {
			_ = index
			if index == len(pResult)-1 {
				return warnings, errors, Video
			} else if len(pResult[index:len(pResult)]) > 1 {
				warnings = append(warnings, errors...)
				return warnings, errors, Playlist
			} else {
				warnings = append(warnings, errors...)
				return warnings, errors, Generic
			}
		}

	}

	return warnings, errors, Generic
}

func (d *download) GetMetadata(verbose bool) e.Metadata {

	warnings, errors, indicatorType := getIndicatorType(d.url)

	if len(warnings) > 0 {
		log.Info(warnings)
	}
	if len(errors) > 0 {
		log.Info(errors)
	}

	//in case of errors just terminate the flow and return the error to UI
	fmt.Println(indicatorType)
	//in case of errors just terminate the flow and return the error to UI

	metadata := e.Metadata{}
	args, command := cmdBuilderMetadata(d.url, indicatorType, false)
	metadata.Command = command + Space + args

	//log executed command
	fmt.Println(metadata.Command)

	cmd, stdout := ProcessBuilder(args, GetCommandString())
	//fmt.Println(cmd.String())

	err := cmd.Start()
	handleErrors(err, "Metadata - Cmd.Start")

	var pResult []string
	if indicatorType == Video {
		pResult = ProcessResult(stdout)
	}

	if indicatorType == Playlist {
		pResult = ProcessResult(stdout)
	}

	fmt.Println(pResult)

	return metadata
}

func (d *download) GetVideo() {

}

// if string list in URL then its a playlist.
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
func ProcessBuilder(args string, command string) (*exec.Cmd, io.ReadCloser) {

	cmd := exec.Command(command)
	cmd.SysProcAttr = &syscall.SysProcAttr{}
	cmd.SysProcAttr.CmdLine = command + Space + args

	stdout, err := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout
	handleErrors(err, "Metadata - StdoutPipe")

	return cmd, stdout
}

func ProcessResult(stdout io.ReadCloser) []string {
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

	//split by newlines and remove and from the end.
	result := strings.Split(b.String(), "\n")
	if result[len(result)-1] == "" {
		result = result[:len(result)-1]
	}

	return result
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
