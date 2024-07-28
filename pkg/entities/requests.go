package entities

// Field names should start with an uppercase letter
type IncomingRequest struct {
	Indicator    string `json:"Indicator"`
	SubtitlesReq bool   `json:"SubtitlesReq"`
	IsAudioOnly  bool   `json:"IsAudioOnly"`
}

type QueueDownloads struct {
	DownloadVideos []string `json:"DownloadVideos"`
}
