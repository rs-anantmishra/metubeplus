package entities

type MediaInformation struct {
	YoutubeVideoId string   `json:"id"`
	Channel        string   `json:"channel"`
	Title          string   `json:"title"`
	Description    string   `json:"description"`
	Extension      string   `json:"ext"`
	Duration       int      `json:"duration"`
	Domain         string   `json:"webpage_url_domain"`
	OriginalURL    string   `json:"original_url"`
	PlaylistTitle  string   `json:"playlist_title"`
	PlaylistCount  int      `json:"playlist_count"`
	PlaylistIndex  int      `json:"playlist_index"`
	Tags           []string `json:"tags"`
	Format         string   `json:"format"`
	Filesize       int      `json:"filesize_approx"`
	FormatNote     string   `json:"format_note"`
	Resolution     string   `json:"resolution"`
	Categories     []string `json:"categories"`
	ChannelId      string   `json:"channel_id"`
	ChannelURL     string   `json:"channel_url"`
	PlaylistId     string   `json:"playlist_id"`
	Availability   string   `json:"availability"`
	LiveStatus     string   `json:"live_status"`
}

// type DownloadProgressIndicator struct {
// 	Filepath string `json:"filepath"`
// 	Title    string
// 	Progress string
// }

type SavedMediaInformation struct {
	VideoId        int
	Title          string
	YoutubeVideoId string
}
