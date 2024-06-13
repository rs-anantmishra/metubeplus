package videos

type CSwitch struct {
	Index int
	Name  string
	Value string
	Group FxGroups
}

type FxGroups struct {
	Playlist Functions
	Video    Functions
	Audio    Functions
}

const (
	Playlist = iota
	Video    = iota
	Audio    = iota
)

type Functions struct {
	Metadata  bool
	Download  bool
	Subtitle  bool
	Thumbnail bool
}

func BuilderOptions() []CSwitch {

	//these true false patterns are talking about default download options
	//this forms the basis of the exectue-custom-commands that may be implemented later on
	//flexibility of cutom commands may still be a question mark
	defaults := []CSwitch{

		{Index: 1, Name: `Filepath`, Value: Filepath, Group: FxGroups{
			Playlist: Functions{Metadata: false, Download: false, Subtitle: false, Thumbnail: false},
			Video:    Functions{Metadata: false, Download: false, Subtitle: false, Thumbnail: false},
			Audio:    Functions{Metadata: false, Download: false, Subtitle: false, Thumbnail: false}},
		},
		{Index: 2, Name: `Channel`, Value: Channel, Group: FxGroups{
			Playlist: Functions{Metadata: true, Download: false, Subtitle: false, Thumbnail: false},
			Video:    Functions{Metadata: true, Download: false, Subtitle: false, Thumbnail: false},
			Audio:    Functions{Metadata: true, Download: false, Subtitle: false, Thumbnail: false}},
		},
		{Index: 3, Name: `Title`, Value: Title, Group: FxGroups{
			Playlist: Functions{Metadata: true, Download: false, Subtitle: false, Thumbnail: false},
			Video:    Functions{Metadata: true, Download: false, Subtitle: false, Thumbnail: false},
			Audio:    Functions{Metadata: true, Download: false, Subtitle: false, Thumbnail: false}},
		},
		{Index: 4, Name: `Description`, Value: Description, Group: FxGroups{
			Playlist: Functions{Metadata: true, Download: false, Subtitle: false, Thumbnail: false},
			Video:    Functions{Metadata: true, Download: false, Subtitle: false, Thumbnail: false},
			Audio:    Functions{Metadata: true, Download: false, Subtitle: false, Thumbnail: false}},
		},
		{Index: 5, Name: `Extension`, Value: Extension, Group: FxGroups{
			Playlist: Functions{Metadata: true, Download: false, Subtitle: false, Thumbnail: false},
			Video:    Functions{Metadata: true, Download: false, Subtitle: false, Thumbnail: false},
			Audio:    Functions{Metadata: true, Download: false, Subtitle: false, Thumbnail: false}},
		},
		{Index: 6, Name: `Duration`, Value: Duration, Group: FxGroups{
			Playlist: Functions{Metadata: true, Download: false, Subtitle: false, Thumbnail: false},
			Video:    Functions{Metadata: true, Download: false, Subtitle: false, Thumbnail: false},
			Audio:    Functions{Metadata: true, Download: false, Subtitle: false, Thumbnail: false}},
		},
		{Index: 7, Name: `URLDomain`, Value: URLDomain, Group: FxGroups{
			Playlist: Functions{Metadata: true, Download: false, Subtitle: false, Thumbnail: false},
			Video:    Functions{Metadata: true, Download: false, Subtitle: false, Thumbnail: false},
			Audio:    Functions{Metadata: true, Download: false, Subtitle: false, Thumbnail: false}},
		},
		{Index: 8, Name: `OriginalURL`, Value: OriginalURL, Group: FxGroups{
			Playlist: Functions{Metadata: true, Download: false, Subtitle: false, Thumbnail: false},
			Video:    Functions{Metadata: true, Download: false, Subtitle: false, Thumbnail: false},
			Audio:    Functions{Metadata: true, Download: false, Subtitle: false, Thumbnail: false}},
		},
		{Index: 9, Name: `PlaylistTitle`, Value: PlaylistTitle, Group: FxGroups{
			Playlist: Functions{Metadata: true, Download: false, Subtitle: false, Thumbnail: false},
			Video:    Functions{Metadata: false, Download: false, Subtitle: false, Thumbnail: false},
			Audio:    Functions{Metadata: false, Download: false, Subtitle: false, Thumbnail: false}},
		},
		{Index: 10, Name: `PlaylistIndex`, Value: PlaylistIndex, Group: FxGroups{
			Playlist: Functions{Metadata: true, Download: true, Subtitle: false, Thumbnail: false},
			Video:    Functions{Metadata: false, Download: false, Subtitle: false, Thumbnail: false},
			Audio:    Functions{Metadata: false, Download: false, Subtitle: false, Thumbnail: false}},
		},
		{Index: 11, Name: `PlaylistCount`, Value: PlaylistCount, Group: FxGroups{
			Playlist: Functions{Metadata: true, Download: false, Subtitle: false, Thumbnail: false},
			Video:    Functions{Metadata: false, Download: false, Subtitle: false, Thumbnail: false},
			Audio:    Functions{Metadata: false, Download: false, Subtitle: false, Thumbnail: false}},
		},
		{Index: 12, Name: `Tags`, Value: Tags, Group: FxGroups{
			Playlist: Functions{Metadata: true, Download: false, Subtitle: false, Thumbnail: false},
			Video:    Functions{Metadata: true, Download: false, Subtitle: false, Thumbnail: false},
			Audio:    Functions{Metadata: true, Download: false, Subtitle: false, Thumbnail: false}},
		},
		{Index: 13, Name: `YTFormatString`, Value: YTFormatString, Group: FxGroups{
			Playlist: Functions{Metadata: true, Download: false, Subtitle: false, Thumbnail: false},
			Video:    Functions{Metadata: true, Download: false, Subtitle: false, Thumbnail: false},
			Audio:    Functions{Metadata: true, Download: false, Subtitle: false, Thumbnail: false}},
		},
		{Index: 14, Name: `ShowProgress`, Value: ShowProgress, Group: FxGroups{
			Playlist: Functions{Metadata: false, Download: true, Subtitle: false, Thumbnail: false},
			Video:    Functions{Metadata: false, Download: true, Subtitle: false, Thumbnail: false},
			Audio:    Functions{Metadata: false, Download: true, Subtitle: false, Thumbnail: false}},
		},
		{Index: 15, Name: `ProgressDelta`, Value: ProgressDelta, Group: FxGroups{
			Playlist: Functions{Metadata: false, Download: true, Subtitle: false, Thumbnail: false},
			Video:    Functions{Metadata: false, Download: true, Subtitle: false, Thumbnail: false},
			Audio:    Functions{Metadata: false, Download: true, Subtitle: false, Thumbnail: false}},
		},
		{Index: 16, Name: `QuietDownload`, Value: QuietDownload, Group: FxGroups{
			Playlist: Functions{Metadata: false, Download: true, Subtitle: false, Thumbnail: false},
			Video:    Functions{Metadata: false, Download: true, Subtitle: false, Thumbnail: false},
			Audio:    Functions{Metadata: false, Download: true, Subtitle: false, Thumbnail: false}},
		},
		{Index: 17, Name: `ProgressNewline`, Value: ProgressNewline, Group: FxGroups{
			Playlist: Functions{Metadata: false, Download: true, Subtitle: false, Thumbnail: false},
			Video:    Functions{Metadata: false, Download: true, Subtitle: false, Thumbnail: false},
			Audio:    Functions{Metadata: false, Download: true, Subtitle: false, Thumbnail: false}},
		},
		{Index: 18, Name: `SkipDownload`, Value: SkipDownload, Group: FxGroups{
			Playlist: Functions{Metadata: true, Download: false, Subtitle: true, Thumbnail: true},
			Video:    Functions{Metadata: true, Download: false, Subtitle: true, Thumbnail: true},
			Audio:    Functions{Metadata: true, Download: false, Subtitle: true, Thumbnail: true}},
		},
		{Index: 19, Name: `WriteSubtitles`, Value: WriteSubtitles, Group: FxGroups{
			Playlist: Functions{Metadata: false, Download: true, Subtitle: true, Thumbnail: false},
			Video:    Functions{Metadata: false, Download: true, Subtitle: true, Thumbnail: false},
			Audio:    Functions{Metadata: false, Download: false, Subtitle: true, Thumbnail: false}},
		},
		{Index: 20, Name: `WriteThumbnail`, Value: WriteThumbnail, Group: FxGroups{
			Playlist: Functions{Metadata: true, Download: false, Subtitle: false, Thumbnail: true},
			Video:    Functions{Metadata: true, Download: false, Subtitle: false, Thumbnail: true},
			Audio:    Functions{Metadata: true, Download: false, Subtitle: false, Thumbnail: true}},
		},
		{Index: 21, Name: `OutputPlaylistVideoFile`, Value: OutputPlaylistVideoFile, Group: FxGroups{
			Playlist: Functions{Metadata: false, Download: true, Subtitle: false, Thumbnail: false},
			Video:    Functions{Metadata: false, Download: false, Subtitle: false, Thumbnail: false},
			Audio:    Functions{Metadata: false, Download: false, Subtitle: false, Thumbnail: false}},
		},
		{Index: 22, Name: `OutputPlaylistSubtitleFile`, Value: OutputPlaylistSubtitleFile, Group: FxGroups{
			Playlist: Functions{Metadata: false, Download: true, Subtitle: true, Thumbnail: false},
			Video:    Functions{Metadata: false, Download: false, Subtitle: false, Thumbnail: false},
			Audio:    Functions{Metadata: false, Download: false, Subtitle: false, Thumbnail: false}},
		},
		{Index: 23, Name: `OutputPlaylistThumbnailFile`, Value: OutputPlaylistThumbnailFile, Group: FxGroups{
			Playlist: Functions{Metadata: true, Download: false, Subtitle: false, Thumbnail: true},
			Video:    Functions{Metadata: false, Download: false, Subtitle: false, Thumbnail: false},
			Audio:    Functions{Metadata: false, Download: false, Subtitle: false, Thumbnail: false}},
		},
		{Index: 24, Name: `OutputVideoFile`, Value: OutputVideoFile, Group: FxGroups{
			Playlist: Functions{Metadata: false, Download: false, Subtitle: false, Thumbnail: false},
			Video:    Functions{Metadata: false, Download: true, Subtitle: false, Thumbnail: false},
			Audio:    Functions{Metadata: false, Download: false, Subtitle: false, Thumbnail: false}},
		},
		{Index: 25, Name: `OutputSubtitleFile`, Value: OutputSubtitleFile, Group: FxGroups{
			Playlist: Functions{Metadata: false, Download: false, Subtitle: false, Thumbnail: false},
			Video:    Functions{Metadata: false, Download: true, Subtitle: true, Thumbnail: false},
			Audio:    Functions{Metadata: false, Download: false, Subtitle: false, Thumbnail: false}},
		},
		{Index: 26, Name: `OutputThumbnailFile`, Value: OutputThumbnailFile, Group: FxGroups{
			Playlist: Functions{Metadata: false, Download: false, Subtitle: false, Thumbnail: false},
			Video:    Functions{Metadata: true, Download: false, Subtitle: false, Thumbnail: true},
			Audio:    Functions{Metadata: false, Download: false, Subtitle: false, Thumbnail: false}},
		},
		//Audio only file options to be added later
	}

	return defaults
}
