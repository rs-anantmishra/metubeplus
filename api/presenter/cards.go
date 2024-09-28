package presenter

type CardsInfoResponse struct {
	VideoId            int    `json:"video_id"`
	Title              string `json:"title"`
	Description        string `json:"description"`
	Duration           int    `json:"duration"`
	OriginalURL        string `json:"original_url"`
	Thumbnail          string `json:"thumbnail"`
	PlaylistThumbnail  string `json:"pl_thumbnail"`
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

type DownloadStatusResponse struct {
	Message  string `json:"download"`
	VideoURL string `json:"video_url"`
}

type LimitedCardsInfoResponse struct {
	VideoId       int    `json:"video_id"`
	Title         string `json:"title"`
	Description   string `json:"description"`
	Duration      int    `json:"duration"`
	WebpageURL    string `json:"original_url"`
	Thumbnail     string `json:"thumbnail"`
	VideoFilepath string `json:"video_filepath"`
	Channel       string `json:"channel"`
}
