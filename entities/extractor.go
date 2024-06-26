package entities

type MediaInformation struct {
	URL           string   `json:"url"`
	Filepath      string   `json:"filepath"`
	Channel       string   `json:"channel"`
	Title         string   `json:"title"`
	Description   string   `json:"description"`
	Extension     string   `json:"ext"`
	Duration      string   `json:"duration"`
	Domain        string   `json:"webpage_url_domain"`
	OriginalURL   string   `json:"original_url"`
	PlaylistTitle string   `json:"playlist_title"`
	PlaylistCount string   `json:"playlist_count"`
	PlaylistIndex string   `json:"playlist_index"`
	Tags          []string `json:"tags"`
	Format        string   `json:"format"`
	Filesize      string   `json:"filesize_approx"`
	FormatNote    string   `json:"format_note"`
	Resolution    string   `json:"resolution"`
	Categories    string   `json:"categories"`
	ChannelId     string   `json:"channel_id"`
	ChannelURL    string   `json:"channel_url"`
	PlaylistId    string   `json:"playlist_id"`
}

type DownloadProgressIndicator struct {
	Progress string
}
