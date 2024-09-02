package entities

type Videos struct {
	Id                int
	Title             string
	Description       string
	DurationSeconds   int
	WebpageURL        string
	IsFileDownloaded  int
	IsDeleted         int
	Channel           Channel
	Playlist          Playlist
	LiveStatus        string
	Domain            Domain
	Availability      string
	Format            Format
	ThumbnailFilePath string
	VideoFilePath     string
	Files             []Files
	Tags              []Tags
	Categories        []Categories
	YoutubeVideoId    string
	CreatedDate       int
}

type Channel struct {
	Id               int
	Name             string
	ChannelURL       string
	YoutubeChannelId int
	CreatedDate      int
}

type Playlist struct {
	Id                int
	Title             string
	VideoCount        int
	Directory         string
	ChannelId         int
	File              Files //for thumbnail file
	YoutubePlaylistId int
	CreatedDate       int
}

type Domain struct {
	Id          int
	Domain      string
	CreatedDate int
}

type Format struct {
	Id          int
	Format      string
	FormatNote  string
	Resolution  string
	StreamType  string //Audio or Video
	CreatedDate string
}

// AvailabilityType Constants
const (
	Private        = iota
	PremiumOnly    = iota
	SubscriberOnly = iota
	NeedsAuth      = iota
	Unlisted       = iota
	Public         = iota
)

// LiveStatusType Constants
const (
	NotLive    = iota
	IsLive     = iota
	IsUpcoming = iota
	WasLive    = iota
	PostLive   = iota
)

//Source - yt-dlp only
//Operations - DOWNLOAD, READ, DELETE
