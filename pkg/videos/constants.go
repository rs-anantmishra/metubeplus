package videos

import (
	"strings"

	c "github.com/rs-anantmishra/metubeplus/config"
)

// --Options---------------------------------------------------------------------------------------//
const ShowProgress string = `--progress`
const ProgressDelta string = `--progress-delta 1` //seconds
const Filepath string = `--print after_move:filepath`
const Channel string = `--print before_dl:channel`
const Title string = `--print before_dl:title`
const Description string = `--print before_dl:description`
const Extension string = `--print before_dl:ext`
const Duration string = `--print before_dl:duration`
const URLDomain string = `--print before_dl:webpage_url_domain`
const OriginalURL string = `--print before_dl:original_url`
const PlaylistTitle string = `--print after_move:playlist_title`
const PlaylistCount string = `--print after_move:playlist_count`
const PlaylistIndex string = `--print after_move:playlist_index`
const WriteSubtitles string = `--write-auto-subs`
const Tags string = `--print before_dl:tags`

// --Extras----------------------------------------------------------------------------------------//
const WriteThumbnail string = `--write-thumbnail`
const YTFormatString string = `--print before_dl:format`
const SkipDownload string = `--skip-download`
const InfoJSON string = `--write-info-json`
const QuietDownload string = `--quiet`
const ProgressNewline string = `--newline`

// --Commands--------------------------------------------------------------------------------------//
// Playlist: Videos, Subtitles, Thumbnails
const OutputPlaylistVideoFile string = `-o "%(webpage_url_domain)s/%(channel)s/%(playlist)s/%(playlist_index)s - %(title)s [%(id)s].%(ext)s"`
const OutputPlaylistSubtitleFile string = `-o "subtitle:%(webpage_url_domain)s/%(channel)s/%(playlist)s/Subtitles/%(playlist_index)s - %(title)s [%(id)s].%(ext)s"`
const OutputPlaylistThumbnailFile string = `-o "thumbnail:%(webpage_url_domain)s/%(channel)s/%(playlist)s/Thumbnails/%(playlist_index)s - %(title)s [%(id)s].%(ext)s"`

// Videos: Videos, Subtitles, Thumbnails
const OutputVideoFile string = `-o "%(webpage_url_domain)s/%(channel)s/Videos/%(title)s [%(id)s].%(ext)s"`
const OutputSubtitleFile string = `-o "subtitle:%(webpage_url_domain)s/%(channel)s/Videos/Subtitles/%(title)s [%(id)s].%(ext)s"`
const OutputThumbnailFile string = `-o "thumbnail:%(webpage_url_domain)s/%(channel)s/Videos/Thumbnails/%(title)s [%(id)s].%(ext)s"`

// --Testing----------------------------------------------------------------------------------------//
const TestURL1 string = `https://www.youtube.com/watch?v=GW2g-5WALrc`
const TestURL2 string = `https://www.youtube.com/watch?v=GW2g-5WALrc&list=PLFKeDWeuu3BZEBcRmolX6BDiFhK-GhCsd`
const TestURL3 string = `https://www.youtube.com/watch?v=-VC4FuG8P6Q`

// --Test Command Playlist--------------------------------------------------------------------------//
const TestCmdPlaylist string = `yt-dlp_x86.exe "https://www.youtube.com/watch?v=5WfiTHiU4x8&list=PLIhvC56v63IKrRHh3gvZZBAGvsvOhwrRF" -P "./" -o "%(webpage_url_domain)s/%(channel)s/%(playlist)s/%(playlist_index)s - %(title)s [%(id)s].%(ext)s" -o "subtitle:%(webpage_url_domain)s/%(channel)s/%(playlist)s/subs/%(playlist_index)s - %(title)s [%(id)s].%(ext)s" -o "thumbnail:%(webpage_url_domain)s/%(channel)s/%(playlist)s/thumbnails/%(playlist_index)s - %(title)s [%(id)s].%(ext)s" -S "res:240" --write-thumbnail --write-auto-subs`

// --Test Command Videos----------------------------------------------------------------------------//
const TestCmdVideo string = `yt-dlp_x86.exe "https://www.youtube.com/watch?v=AaseHnf0k2o" -P "./" -o "%(webpage_url_domain)s/%(channel)s/Videos/%(title)s [%(id)s].%(ext)s" -o "subtitle:%(webpage_url_domain)s/%(channel)s/Videos/Subtitles/%(title)s [%(id)s].%(ext)s" -o "thumbnail:%(webpage_url_domain)s/%(channel)s/Videos/Thumbnails/%(title)s [%(id)s].%(ext)s"  -S "res:240" --write-thumbnail --write-auto-subs`

// command path should be picked from .env
const CommandName string = `yt-dlp_x86.exe`
const Space string = " "

// Parent Directory
const mediaDirectory string = `-P {{MediaDir}}`

func GetMediaDirectory() string {
	mediaDir := strings.ReplaceAll(mediaDirectory, "{{MediaDir}}", c.Config("MEDIA_PATH"))
	return mediaDir
}
