package entities

type Tags struct {
	Id          int
	Tag         string
	IsUsed      int
	Source      string //Youtube or User?
	CreatedDate int
}

type VideoFileTags struct {
	Id      int
	VideoId int
	FileId  int
	TagId   int
}

// can be created through UI or through yt-dlp
// operations include - INSERT, DELETE, READ
