package entities

type Metadata struct {
	URL         string
	Progress    string
	Filepath    string
	Channel     string
	Title       string
	Description string
	Extension   string
	Duration    string
	Domain      string
	OriginalURL string
	Playlist    PlaylistMeta //replace unrolled above
	Thumbnail   string
	Format      string
	Tags        []string
	Files       FilesMeta //replace unrolled above
	Command     string

	// PlaylistTitle string
	// PlaylistCount string
	// PlaylistIndex string
	// VideoDir  string
	// SubsDir   string
	// ThumbsDir string
}

type PlaylistMeta struct {
	PlaylistTitle string
	PlaylistCount string
	PlaylistIndex string
}

type FilesMeta struct {
	VideoDir  string
	SubsDir   string
	ThumbsDir string
}
