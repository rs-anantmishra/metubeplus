package presenter

type CardsInfo struct {
	VideoId            int    `json:"video_id"`
	Title              string `json:"title"`
	Description        string `json:"description"`
	Duration           int    `json:"duration"`
	OriginalURL        string `json:"original_url"`
	ThumbnailURL       string `json:"thumbnail_url"`
	SubtitlesURL       string `json:"subs_url"`
	MediaURL           string `json:"media_url"`
	IsFileDownloaded   bool   `json:"is_file_downloaded"`
	Channel            string `json:"channel"`
	Playlist           string `json:"playlist"`
	PlaylistVideoIndex int    `json:"playlist_video_index"`
	Domain             string `json:"domain"`
	VideoFormat        string `json:"video_format"`
	WatchCount         int    `json:"watch_count"`
	IsDeleted          bool   `json:"is_deleted"`
	CreatedDate        int    `json:"created_date"`
}
