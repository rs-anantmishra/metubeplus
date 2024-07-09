package extractor

import (
	"strings"

	c "github.com/rs-anantmishra/metubeplus/config"
	e "github.com/rs-anantmishra/metubeplus/pkg/entities"
)

type CSwitch struct {
	Index     int
	Name      string
	Value     string
	DataField bool
	Group     FxGroups
}

type FxGroups struct {
	Playlist Functions
	Video    Functions
	Generic  Functions
}

const (
	Generic  = iota
	Video    = iota
	Playlist = iota
)

type Functions struct {
	Metadata  bool
	Download  bool
	Subtitle  bool
	Thumbnail bool
}

func BuilderOptions() []CSwitch {

	//these true false patterns are talking about default download options
	//this forms the basis of the execute-custom-commands that may be implemented later on
	//flexibility of cutom commands may still be a question mark
	//ideally this should be moved to a db or read from a config file.
	defaults := []CSwitch{
		{Index: 1, Name: `Filepath`, Value: Filepath, DataField: true, Group: FxGroups{
			Playlist: Functions{Metadata: false, Download: true, Subtitle: false, Thumbnail: false},
			Video:    Functions{Metadata: false, Download: true, Subtitle: false, Thumbnail: false}},
		},
		{Index: 2, Name: `Channel`, Value: Channel, DataField: true, Group: FxGroups{
			Playlist: Functions{Metadata: true, Download: false, Subtitle: false, Thumbnail: false},
			Video:    Functions{Metadata: true, Download: false, Subtitle: false, Thumbnail: false}},
		},
		{Index: 3, Name: `Title`, Value: Title, DataField: true, Group: FxGroups{
			Playlist: Functions{Metadata: true, Download: false, Subtitle: false, Thumbnail: false},
			Video:    Functions{Metadata: true, Download: false, Subtitle: false, Thumbnail: false}},
		},
		{Index: 4, Name: `Description`, Value: Description, DataField: true, Group: FxGroups{
			Playlist: Functions{Metadata: true, Download: false, Subtitle: false, Thumbnail: false},
			Video:    Functions{Metadata: true, Download: false, Subtitle: false, Thumbnail: false}},
		},
		{Index: 5, Name: `Extension`, Value: Extension, DataField: true, Group: FxGroups{
			Playlist: Functions{Metadata: true, Download: false, Subtitle: false, Thumbnail: false},
			Video:    Functions{Metadata: true, Download: false, Subtitle: false, Thumbnail: false}},
		},
		{Index: 6, Name: `Duration`, Value: Duration, DataField: true, Group: FxGroups{
			Playlist: Functions{Metadata: true, Download: false, Subtitle: false, Thumbnail: false},
			Video:    Functions{Metadata: true, Download: false, Subtitle: false, Thumbnail: false}},
		},
		{Index: 7, Name: `URLDomain`, Value: URLDomain, DataField: true, Group: FxGroups{
			Playlist: Functions{Metadata: true, Download: false, Subtitle: false, Thumbnail: false},
			Video:    Functions{Metadata: true, Download: false, Subtitle: false, Thumbnail: false}},
		},
		{Index: 8, Name: `OriginalURL`, Value: OriginalURL, DataField: true, Group: FxGroups{
			Playlist: Functions{Metadata: true, Download: false, Subtitle: false, Thumbnail: false},
			Video:    Functions{Metadata: true, Download: false, Subtitle: false, Thumbnail: false}},
		},
		{Index: 9, Name: `PlaylistTitle`, Value: PlaylistTitle, DataField: true, Group: FxGroups{
			Playlist: Functions{Metadata: true, Download: false, Subtitle: false, Thumbnail: false},
			Video:    Functions{Metadata: false, Download: false, Subtitle: false, Thumbnail: false}},
		},
		{Index: 10, Name: `PlaylistIndex`, Value: PlaylistIndex, DataField: true, Group: FxGroups{
			Playlist: Functions{Metadata: true, Download: true, Subtitle: false, Thumbnail: false},
			Video:    Functions{Metadata: false, Download: false, Subtitle: false, Thumbnail: false}},
		},
		{Index: 11, Name: `PlaylistCount`, Value: PlaylistCount, DataField: true, Group: FxGroups{
			Playlist: Functions{Metadata: true, Download: false, Subtitle: false, Thumbnail: false},
			Video:    Functions{Metadata: false, Download: false, Subtitle: false, Thumbnail: false}},
		},
		{Index: 12, Name: `Tags`, Value: Tags, DataField: true, Group: FxGroups{
			Playlist: Functions{Metadata: true, Download: false, Subtitle: false, Thumbnail: false},
			Video:    Functions{Metadata: true, Download: false, Subtitle: false, Thumbnail: false}},
		},
		{Index: 13, Name: `YTFormatString`, Value: YTFormatString, DataField: true, Group: FxGroups{
			Playlist: Functions{Metadata: true, Download: false, Subtitle: false, Thumbnail: false},
			Video:    Functions{Metadata: true, Download: false, Subtitle: false, Thumbnail: false}},
		},
		{Index: 14, Name: `FileSizeApprox`, Value: FileSizeApprox, DataField: true, Group: FxGroups{
			Playlist: Functions{Metadata: true, Download: false, Subtitle: false, Thumbnail: false},
			Video:    Functions{Metadata: true, Download: false, Subtitle: false, Thumbnail: false}},
		},
		{Index: 15, Name: `FormatNote`, Value: FormatNote, DataField: true, Group: FxGroups{
			Playlist: Functions{Metadata: true, Download: false, Subtitle: false, Thumbnail: false},
			Video:    Functions{Metadata: true, Download: false, Subtitle: false, Thumbnail: false}},
		},
		{Index: 16, Name: `Resolution`, Value: Resolution, DataField: true, Group: FxGroups{
			Playlist: Functions{Metadata: true, Download: false, Subtitle: false, Thumbnail: false},
			Video:    Functions{Metadata: true, Download: false, Subtitle: false, Thumbnail: false}},
		},
		{Index: 17, Name: `Categories`, Value: Categories, DataField: true, Group: FxGroups{
			Playlist: Functions{Metadata: true, Download: false, Subtitle: false, Thumbnail: false},
			Video:    Functions{Metadata: true, Download: false, Subtitle: false, Thumbnail: false}},
		},
		{Index: 18, Name: `YoutubeVideoId`, Value: YoutubeVideoId, DataField: true, Group: FxGroups{
			Playlist: Functions{Metadata: true, Download: true, Subtitle: false, Thumbnail: false},
			Video:    Functions{Metadata: true, Download: true, Subtitle: false, Thumbnail: false}},
		},
		{Index: 19, Name: `Availability`, Value: Availability, DataField: true, Group: FxGroups{
			Playlist: Functions{Metadata: true, Download: false, Subtitle: false, Thumbnail: false},
			Video:    Functions{Metadata: true, Download: false, Subtitle: false, Thumbnail: false}},
		},
		{Index: 20, Name: `LiveStatus`, Value: LiveStatus, DataField: true, Group: FxGroups{
			Playlist: Functions{Metadata: true, Download: false, Subtitle: false, Thumbnail: false},
			Video:    Functions{Metadata: true, Download: false, Subtitle: false, Thumbnail: false}},
		},
		{Index: 21, Name: `ChannelId`, Value: ChannelId, DataField: true, Group: FxGroups{
			Playlist: Functions{Metadata: true, Download: false, Subtitle: false, Thumbnail: false},
			Video:    Functions{Metadata: true, Download: false, Subtitle: false, Thumbnail: false}},
		},
		{Index: 22, Name: `ChannelURL`, Value: ChannelURL, DataField: true, Group: FxGroups{
			Playlist: Functions{Metadata: true, Download: false, Subtitle: false, Thumbnail: false},
			Video:    Functions{Metadata: true, Download: false, Subtitle: false, Thumbnail: false}},
		},
		{Index: 23, Name: `PlaylistId`, Value: PlaylistId, DataField: true, Group: FxGroups{
			Playlist: Functions{Metadata: true, Download: false, Subtitle: false, Thumbnail: false},
			Video:    Functions{Metadata: false, Download: false, Subtitle: false, Thumbnail: false}},
		},
		{Index: 24, Name: `ShowProgress`, Value: ShowProgress, DataField: false, Group: FxGroups{
			Playlist: Functions{Metadata: false, Download: true, Subtitle: true, Thumbnail: true},
			Video:    Functions{Metadata: false, Download: true, Subtitle: true, Thumbnail: true}},
		},
		{Index: 25, Name: `ProgressDelta`, Value: ProgressDelta, DataField: false, Group: FxGroups{
			Playlist: Functions{Metadata: false, Download: true, Subtitle: true, Thumbnail: true},
			Video:    Functions{Metadata: false, Download: true, Subtitle: true, Thumbnail: true}},
		},
		{Index: 26, Name: `QuietDownload`, Value: QuietDownload, DataField: false, Group: FxGroups{
			Playlist: Functions{Metadata: false, Download: true, Subtitle: false, Thumbnail: false},
			Video:    Functions{Metadata: false, Download: true, Subtitle: false, Thumbnail: false}},
		},
		{Index: 27, Name: `ProgressNewline`, Value: ProgressNewline, DataField: false, Group: FxGroups{
			Playlist: Functions{Metadata: false, Download: true, Subtitle: true, Thumbnail: true},
			Video:    Functions{Metadata: false, Download: true, Subtitle: true, Thumbnail: true}},
		},
		{Index: 28, Name: `SkipDownload`, Value: SkipDownload, DataField: false, Group: FxGroups{
			Playlist: Functions{Metadata: true, Download: false, Subtitle: true, Thumbnail: true},
			Video:    Functions{Metadata: true, Download: false, Subtitle: true, Thumbnail: true}},
		},
		{Index: 29, Name: `WriteSubtitles`, Value: WriteSubtitles, DataField: false, Group: FxGroups{
			Playlist: Functions{Metadata: false, Download: false, Subtitle: true, Thumbnail: false},
			Video:    Functions{Metadata: false, Download: false, Subtitle: true, Thumbnail: false}},
		},
		{Index: 30, Name: `WriteThumbnail`, Value: WriteThumbnail, DataField: false, Group: FxGroups{
			Playlist: Functions{Metadata: false, Download: false, Subtitle: false, Thumbnail: true},
			Video:    Functions{Metadata: false, Download: false, Subtitle: false, Thumbnail: true}},
		},
		{Index: 31, Name: `MediaDirectory`, Value: GetMediaDirectory(), DataField: false, Group: FxGroups{
			Playlist: Functions{Metadata: false, Download: true, Subtitle: true, Thumbnail: true},
			Video:    Functions{Metadata: false, Download: true, Subtitle: true, Thumbnail: true}},
		},
		{Index: 32, Name: `OutputPlaylistVideoFile`, Value: OutputPlaylistVideoFile, DataField: false, Group: FxGroups{
			Playlist: Functions{Metadata: false, Download: true, Subtitle: false, Thumbnail: false},
			Video:    Functions{Metadata: false, Download: false, Subtitle: false, Thumbnail: false}},
		},
		{Index: 33, Name: `OutputPlaylistSubtitleFile`, Value: OutputPlaylistSubtitleFile, DataField: false, Group: FxGroups{
			Playlist: Functions{Metadata: false, Download: false, Subtitle: true, Thumbnail: false},
			Video:    Functions{Metadata: false, Download: false, Subtitle: false, Thumbnail: false}},
		},
		{Index: 34, Name: `OutputPlaylistThumbnailFile`, Value: OutputPlaylistThumbnailFile, DataField: false, Group: FxGroups{
			Playlist: Functions{Metadata: false, Download: false, Subtitle: false, Thumbnail: true},
			Video:    Functions{Metadata: false, Download: false, Subtitle: false, Thumbnail: false}},
		},
		{Index: 35, Name: `OutputVideoFile`, Value: OutputVideoFile, DataField: false, Group: FxGroups{
			Playlist: Functions{Metadata: false, Download: false, Subtitle: false, Thumbnail: false},
			Video:    Functions{Metadata: false, Download: true, Subtitle: false, Thumbnail: false}},
		},
		{Index: 36, Name: `OutputSubtitleFile`, Value: OutputSubtitleFile, DataField: false, Group: FxGroups{
			Playlist: Functions{Metadata: false, Download: false, Subtitle: false, Thumbnail: false},
			Video:    Functions{Metadata: false, Download: false, Subtitle: true, Thumbnail: false}},
		},
		{Index: 37, Name: `OutputThumbnailFile`, Value: OutputThumbnailFile, DataField: false, Group: FxGroups{
			Playlist: Functions{Metadata: false, Download: false, Subtitle: false, Thumbnail: false},
			Video:    Functions{Metadata: false, Download: false, Subtitle: false, Thumbnail: true}},
		},
		//Audio only file options to be added later
	}

	return defaults
}

func GetCommandString() string {
	cmdPath := c.Config("YTDLP_PATH")
	return cmdPath + CommandName
}

func GetVideoFilepath(fp e.Filepath, fType int) string {
	var result string
	if fType == e.Thumbnail {
		result = strings.Join([]string{c.Config("MEDIA_PATH"), fp.Domain, fp.Channel, "Videos", "Thumbnails"}, "\\")
	} else if fType == e.Subtitles {
		result = strings.Join([]string{c.Config("MEDIA_PATH"), fp.Domain, fp.Channel, "Videos", "Subtitles"}, "\\")
	}
	return result
}

func GetPlaylistFilepath(fp e.Filepath, fType int) string {

	var result string
	if fType == e.Thumbnail {
		result = strings.Join([]string{c.Config("MEDIA_PATH"), fp.Domain, fp.Channel, fp.PlaylistTitle, "Thumbnails"}, "\\")
	} else if fType == e.Subtitles {
		result = strings.Join([]string{c.Config("MEDIA_PATH"), fp.Domain, fp.Channel, fp.PlaylistTitle, "Subtitles"}, "\\")
	} else {
		result = strings.Join([]string{c.Config("MEDIA_PATH"), fp.Domain, fp.Channel, fp.PlaylistTitle}, "\\")
	}
	return result
}

func cmdBuilderMetadata(url string, indicatorType int) (string, string) {

	var args []string
	args = append(args, "\""+url+"\"")

	bo := BuilderOptions()
	for _, elem := range bo {

		//Handle Video
		if indicatorType == Video && elem.Group.Video.Metadata {
			args = append(args, elem.Value)
		}

		//Handle Playlist
		if indicatorType == Playlist && elem.Group.Playlist.Metadata {
			args = append(args, elem.Value)
		}

		//Todo: Handle Audio
		//Todo: Handle Audio Playlists
	}

	arguments := strings.Join(args, Space)
	cmdPath := c.Config("YTDLP_PATH")
	cmd := cmdPath + "/" + CommandName

	return arguments, cmd
}

// Download Media Content
func cmdBuilderDownload() (string, string) {

	return "nil", "nil"
}

func cmdBuilderSubtitles(url string, indicatorType int) (string, string) {

	var args []string
	args = append(args, "\""+url+"\"")

	bo := BuilderOptions()
	for _, elem := range bo {

		//Handle Video
		if indicatorType == Video && elem.Group.Video.Subtitle {
			args = append(args, elem.Value)
		}

		//Handle Playlist
		if indicatorType == Playlist && elem.Group.Playlist.Subtitle {
			args = append(args, elem.Value)
		}
	}

	arguments := strings.Join(args, Space)
	cmdPath := c.Config("YTDLP_PATH")
	cmd := cmdPath + "/" + CommandName

	return arguments, cmd
}

func cmdBuilderThumbnails(url string, indicatorType int) (string, string) {

	var args []string
	args = append(args, "\""+url+"\"")

	bo := BuilderOptions()
	for _, elem := range bo {

		//Handle Video
		if indicatorType == Video && elem.Group.Video.Thumbnail {
			args = append(args, elem.Value)
		}

		//Handle Playlist
		if indicatorType == Playlist && elem.Group.Playlist.Thumbnail {
			args = append(args, elem.Value)
		}
	}

	arguments := strings.Join(args, Space)
	cmdPath := c.Config("YTDLP_PATH")
	cmd := cmdPath + "/" + CommandName

	return arguments, cmd
}
