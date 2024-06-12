package videos

import (
	"fmt"
	"net/url"
	"os/exec"
	"strings"
	"syscall"

	c "github.com/rs-anantmishra/metubeplus/config"
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

func (d *download) GetVideo() {

}

func (d *download) GetMetadata(verbose bool) e.Metadata {

	metadata := e.Metadata{}
	args, command := cmdBuilder(d.url, true)

	cmd := exec.Command(command)
	cmd.SysProcAttr = &syscall.SysProcAttr{}
	cmd.SysProcAttr.CmdLine = command + Space + args

	//log executed command
	metadata.Command = command + Space + args

	stdout, err := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout
	handleErrors(err, "Metadata - StdoutPipe")

	//fmt.Println(cmd.String())
	err = cmd.Start()
	handleErrors(err, "Metadata - Cmd.Start")

	for {
		tmp := make([]byte, 1024)
		n, err := stdout.Read(tmp)

		//handle verbosity
		if verbose {
			// Once the Metadata Object is populated, print that.
			// fmt.Println(command + Space + args) //log Command
			fmt.Println("\r", string(tmp[:n])) //log Metadata
		}

		//terminate loop at eof
		if err != nil {
			break
		}
	}
	return metadata
}

func isPlaylist(link string) bool {

	const playlistKey string = "list"
	const videoKey string = "v"
	result := false

	u, err := url.Parse(link)
	handleErrors(err, "IsPlaylist")

	q := u.Query()
	_, ok := q[playlistKey]
	if ok {
		result = true
	}

	return result
}

func cmdBuilder(url string, meta bool) (string, string) {

	var downloadPath string
	var args []string

	args = append(args, url)
	args = append(args, Channel)
	args = append(args, Title)
	args = append(args, Description)
	args = append(args, Extension)
	args = append(args, Duration)
	args = append(args, URLDomain)
	args = append(args, OriginalURL)
	args = append(args, YTFormatString)

	//Playlist specific
	if isPlaylist(url) {
		args = append(args, PlaylistTitle)
		args = append(args, PlaylistIndex)
		args = append(args, PlaylistCount)
	}

	if meta {
		args = append(args, SkipDownload)
	} else {
		args = append(args, Filepath)
		args = append(args, ProgressDelta)
		args = append(args, ShowProgress)

		//Separate files are created for them
		args = append(args, WriteThumbnail)
		args = append(args, AutoSubtitles)

		//Output filepath only for Download Requests.
		//Video\domain\channel\playlist\filename OR  Video\domain\channel\filename
		//Call Metadata 1st to have output-filepath-name
		_ = downloadPath
	}

	arguments := strings.Join(args, Space)

	cmdPath := c.Config("YTDLP_PATH")
	cmd := cmdPath + "/" + CommandName

	return arguments, cmd
}

func handleErrors(err error, methodName string) {
	if err != nil {
		fmt.Println("pkg dowonload", methodName, err)
	}
}
