package videos

import (
	"strings"

	c "github.com/rs-anantmishra/metubeplus/config"
)

// --Options---------------------------------------------------------------------------------------//
const Separator string = `*MeTube+*`
const ShowProgress string = `--progress`
const ProgressDelta string = `--progress-delta 1`              //seconds
const Filepath string = `--print %(.{filepath})s`              //`--print after_move:filepath`
const Channel string = `--print %(.{channel})s`                //`--print before_dl:channel`
const Title string = `--print %(.{title})s`                    //`--print before_dl:title`
const Description string = `--print %(.{description})s`        //`--print before_dl:description`
const Extension string = `--print %(.{ext})s`                  //`--print before_dl:ext`
const Duration string = `--print %(.{duration})s`              //`--print before_dl:duration`
const URLDomain string = `--print %(.{webpage_url_domain})s`   //`--print before_dl:webpage_url_domain`
const OriginalURL string = `--print %(.{original_url})s`       //`--print before_dl:original_url`
const PlaylistTitle string = `--print %(.{playlist_title})s`   //`--print before_dl:playlist_title`
const PlaylistCount string = `--print %(.{playlist_count})s`   //`--print before_dl:playlist_count`
const PlaylistIndex string = `--print %(.{playlist_index})s`   //`--print before_dl:playlist_index`
const Tags string = `--print %(.{tags})s`                      //`--print before_dl:tags`
const YTFormatString string = `--print %(.{format})s`          //`--print before_dl:format`
const FileSizeApprox string = `--print %(.{filesize_approx})s` //`--print before_dl:filesize_approx`
const FormatNote string = `--print %(.{format_note})s`         //`--print before_dl:format_note`
const Resolution string = `--print %(.{resolution})s`          //`--print before_dl:resolution`
const Categories string = `--print %(.{categories})s`          //`--print before_dl:categories`
const WriteSubtitles string = `--write-auto-subs`

// --Extras----------------------------------------------------------------------------------------//
const WriteThumbnail string = `--write-thumbnail`
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

// Output Parsing: Warnings & Errors
const WARNING string = `WARNING:`
const ERROR string = `ERROR:`

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
