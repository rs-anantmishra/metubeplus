package entities

type Videos struct {
	Id               int
	Title            string
	Description      string
	Duration         int
	Channel          Channel
	Playlist         Playlist
	WebpageURL       string
	Domain           Domain
	IsFileDownloaded int
	IsDeleted        int
	AvailabilityId   int
	LiveStatusId     int
	YoutubeVideoId   int
	CreatedDate      int
	Files            []Files
	Formats          []Formats
	Tags             []Tags
}

type Channel struct {
	Id               int
	Title            string
	WebpageURL       string
	YoutubeChannelId int
	CreatedDate      int
}

type Playlist struct {
	Id                int
	Title             string
	VideoCount        int
	File              Files
	ChannelId         int
	YoutubePlaylistId int
	CreatedDate       int
}

type Domain struct {
	Id          int
	Domain      string
	CreatedDate int
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

type Formats struct {
	Id         int
	Format     string
	StreamType string //Audio or Video
}

type VideoFormat struct {
	Id       int
	VideoId  int
	FormatId int
}

//Source - yt-dlp only
//Operations - DOWNLOAD, READ, DELETE
