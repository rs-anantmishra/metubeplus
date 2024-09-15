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
	VideoId            int
	PlaylistId         int
	YoutubeVideoId     string
	Domain             string
	Channel            string
	Title              string
	PlaylistTitle      string
	PlaylistVideoIndex int
}

// Playlist: Videos, Subtitles, Thumbnails
//const OutputPlaylistVideoFile string = `-o "%(webpage_url_domain)s/%(channel)s/%(playlist)s/%(playlist_index)s - %(title)s [%(id)s].%(ext)s"`
//const OutputPlaylistSubtitleFile string = `-o "subtitle:%(webpage_url_domain)s/%(channel)s/%(playlist)s/Subtitles/%(playlist_index)s - %(title)s [%(id)s].%(ext)s"`
//const OutputPlaylistThumbnailFile string = `-o "%(webpage_url_domain)s/%(channel)s/%(playlist)s/Thumbnails/%(playlist_index)s - %(title)s [%(id)s].%(ext)s"`

// Videos: Videos, Subtitles, Thumbnails
//const OutputVideoFile string = `-o "%(webpage_url_domain)s/%(channel)s/Videos/%(title)s [%(id)s].%(ext)s"`
//const OutputSubtitleFile string = `-o "subtitle:%(webpage_url_domain)s/%(channel)s/Videos/Subtitles/%(title)s [%(id)s].%(ext)s"`
//const OutputThumbnailFile string = `-o "thumbnail:%(webpage_url_domain)s/%(channel)s/Videos/Thumbnails/%(title)s [%(id)s].%(ext)s"`
