package entities

type Metadata struct {
	URL         string `url`
	Progress    string `progress`
	Filepath    string `filepath`
	Channel     string `channel`
	Title       string `title`
	Description string `description`
	Extension   string `ext`
	Duration    string `duration`
	Domain      string `webpage_url_domain`
	OriginalURL string `original_url`
	Playlist    PlaylistMeta
	Thumbnail   string
	Format      string   `format`
	Tags        []string `tags`
	Files       FilesMeta
	Command     string
	Filesize    string `filesize_approx`
	FormatNote  string `format_note`
	Resolution  string `resolution`
	Categories  string `categories`

	// PlaylistTitle string
	// PlaylistCount string
	// PlaylistIndex string
	// VideoDir  string
	// SubsDir   string
	// ThumbsDir string
}

type PlaylistMeta struct {
	PlaylistTitle string `playlist_title`
	PlaylistCount string `playlist_count`
	PlaylistIndex string `playlist_index`
}

type FilesMeta struct {
	VideoDir  string
	SubsDir   string
	ThumbsDir string
}
