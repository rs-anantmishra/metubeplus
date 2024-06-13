package entities

type Files struct {
	Id              int
	FileTypeId      int
	SourceId        int
	FilePath        string
	FileName        string
	Extension       string
	FileSize        int
	FileSizeUnit    string
	Resolution      string
	ParentDirectory string
	IsDeleted       int
	CreatedDate     int
}

// FileType Constants
const (
	Audio     = iota
	Video     = iota
	Thumbnail = iota
	Subtitles = iota
)

// SourceType Constants
const (
	Downloaded   = iota
	Uploaded     = iota
	Local        = iota
	MetadataOnly = iota
)

//operation include - UPLOAD/INSERT, READ, DELETE, METADATA INSERT (for files existing locally)
//Sources - UI, yt-dlp, local, ffmpeg (local file thumbnails), static files - default thumbnails for audio/video files

//operations - socket connection to UI for updates
