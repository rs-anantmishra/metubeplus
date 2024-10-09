package presenter

type CardsInfoResponse struct {
	VideoId            int      `json:"video_id"`
	Title              string   `json:"title"`
	Duration           int      `json:"duration"`
	WebpageURL         string   `json:"webpage_url"`
	Thumbnail          string   `json:"thumbnail"`
	MediaURL           string   `json:"media_url"`
	Channel            string   `json:"channel"`
	SubtitlesURL       string   `json:"subs_url"`
	Playlist           string   `json:"playlist"`
	PlaylistVideoIndex int      `json:"playlist_video_index"`
	Domain             string   `json:"domain"`
	VideoFormat        string   `json:"video_format"`
	WatchCount         int      `json:"watch_count"`
	ViewsCount         int      `json:"views_count"`
	LikesCount         int      `json:"likes_count"`
	FileSize           int      `json:"filesize"`
	UploadDate         string   `json:"upload_date"`
	IsDeleted          bool     `json:"is_deleted"`
	Categories         []string `json:"categories"`
	Tags               []string `json:"tags"`
	CreatedDate        int      `json:"created_date"`
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
	FileSize      int    `json:"filesize"`
	ViewsCount    int    `json:"views_count"`
	LikesCount    int    `json:"likes_count"`
	UploadDate    string `json:"upload_date"`
}
