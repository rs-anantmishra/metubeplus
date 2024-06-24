package entities

type Metadata struct {
	URL         string `json:"url"`
	Progress    string `json:"progress"`
	Filepath    string `json:"filepath"`
	Channel     string `json:"channel"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Extension   string `json:"ext"`
	Duration    string `json:"duration"`
	Domain      string `json:"webpage_url_domain"`
	OriginalURL string `json:"original_url"`
	Playlist    PlaylistMeta
	Thumbnail   string
	Format      string   `json:"format"`
	Tags        []string `json:"tags"`
	Files       FilesMeta
	Command     string
	Filesize    string `json:"filesize_approx"`
	FormatNote  string `json:"format_note"`
	Resolution  string `json:"resolution"`
	Categories  string `json:"categories"`
}

type PlaylistMeta struct {
	PlaylistTitle string `json:"playlist_title"`
	PlaylistCount string `json:"playlist_count"`
	PlaylistIndex string `json:"playlist_index"`
}

type FilesMeta struct {
	VideoDir  string
	SubsDir   string
	ThumbsDir string
}
