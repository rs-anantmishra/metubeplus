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
	args, command := cmdBuilderMetadata(d.url)

	args = args + ` -q`

	cmd := exec.Command(command)
	cmd.SysProcAttr = &syscall.SysProcAttr{}
	cmd.SysProcAttr.CmdLine = command + Space + args

	//log executed command
	metadata.Command = command + Space + args
	fmt.Println(metadata.Command)

	stdout, err := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout
	handleErrors(err, "Metadata - StdoutPipe")

	//fmt.Println(cmd.String())
	err = cmd.Start()
	handleErrors(err, "Metadata - Cmd.Start")

	for {
		tmp := make([]byte, 2048)
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

// never should have made a generic cmdbuilder!
func cmdBuilder(url string, isMetaCmd bool) (string, string) {

	fg := EvaluateFxGroup(url)

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
	args = append(args, Tags)

	//Playlist specific
	if fg == Playlist {
		args = append(args, PlaylistTitle)
		args = append(args, PlaylistIndex)
		args = append(args, PlaylistCount)
	}

	args = append(args, GetMediaDirectory())
	if isMetaCmd {
		//thumbnails will be downloaded as part of Metadata
		if fg == Playlist {
			args = append(args, OutputPlaylistThumbnailFile)
		} else {
			args = append(args, OutputThumbnailFile)
		}
		args = append(args, SkipDownload)
	} else {
		args = append(args, Filepath)
		args = append(args, ProgressDelta)
		args = append(args, ShowProgress)

		// Filepaths for playlist are different
		if fg == Playlist {
			args = append(args, OutputPlaylistVideoFile)
			args = append(args, OutputPlaylistSubtitleFile)
		} else {
			args = append(args, OutputVideoFile)
			args = append(args, OutputSubtitleFile)
		}
		args = append(args, WriteSubtitles)
	}

	arguments := strings.Join(args, Space)

	cmdPath := c.Config("YTDLP_PATH")
	cmd := cmdPath + "/" + CommandName

	return arguments, cmd
}

func cmdBuilderMetadata(url string) (string, string) {

	fg := EvaluateFxGroup(url)

	var args []string
	args = append(args, "\""+url+"\"")

	bo := BuilderOptions()
	for _, elem := range bo {

		//Handle Video
		if fg == Video && elem.Group.Video.Metadata {
			args = append(args, elem.Value)
		}

		//Handle Playlist
		if fg == Playlist && elem.Group.Playlist.Metadata {
			args = append(args, elem.Value)
		}

		//Handle Audio
		if fg == Audio && elem.Group.Audio.Metadata {
			args = append(args, elem.Value)
		}
	}

	arguments := strings.Join(args, Space)
	cmdPath := c.Config("YTDLP_PATH")
	cmd := cmdPath + "/" + CommandName

	return arguments, cmd
}

func cmdBuilderDownload() (string, string) {

	return "nil", "nil"
}

func cmdBuilderSubtitles() (string, string) {

	return "nil", "nil"
}

func cmdBuilderThumbnails() (string, string) {

	return "nil", "nil"
}

func handleErrors(err error, methodName string) {
	if err != nil {
		fmt.Println("pkg dowonload", methodName, err)
	}
}
