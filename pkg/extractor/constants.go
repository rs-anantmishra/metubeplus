package extractor

import (
	"strings"

	c "github.com/rs-anantmishra/metubeplus/config"
)

// --Options---------------------------------------------------------------------------------------//
const Separator string = `*MeTube+*`
const ShowProgress string = `--progress`
const ProgressDelta string = `--progress-delta 1` //seconds
const YoutubeVideoId string = `--print %(.{id})s`
const Availability string = `--print %(.{availability})s`
const LiveStatus string = `--print %(.{live_status})s`
const Filepath string = `--print %(.{filepath})s`
const Channel string = `--print %(.{channel})s`
const Title string = `--print %(.{title})s`
const Description string = `--print %(.{description})s`
const Extension string = `--print %(.{ext})s`
const Duration string = `--print %(.{duration})s`
const URLDomain string = `--print %(.{webpage_url_domain})s`
const OriginalURL string = `--print %(.{original_url})s`
const PlaylistTitle string = `--print %(.{playlist_title})s`
const PlaylistCount string = `--print %(.{playlist_count})s`
const PlaylistIndex string = `--print %(.{playlist_index})s`
const Tags string = `--print %(.{tags})s`
const YTFormatString string = `--print %(.{format})s`
const FileSizeApprox string = `--print %(.{filesize_approx})s`
const FormatNote string = `--print %(.{format_note})s`
const Resolution string = `--print %(.{resolution})s`
const Categories string = `--print %(.{categories})s`
const ChannelId string = `--print %(.{channel_id})s`
const ChannelURL string = `--print %(.{channel_url})s`
const PlaylistId string = `--print %(.{playlist_id})s`
const WriteSubtitles string = `--write-auto-subs`

// --Options Plaintext-----------------------------------------------------------------------------//
const Plaintext_YoutubeVideoId string = `--print after_move:id`
const Plaintext_Availability string = `--print after_move:availability`
const Plaintext_LiveStatus string = `--print after_move:live_status`
const Plaintext_Filepath string = `--print after_move:filepath`
const Plaintext_Channel string = `--print before_dl:"Channel: %(channel)s"`             //Changed for patching data fields. `--print before_dl:channel`
const Plaintext_Title string = `--print before_dl:"Title: %(title)s"`                   //Changed for patching data fields. `--print before_dl:title`
const Plaintext_Description string = `--print before_dl:"Description: %(description)s"` //Changed for patching data fields. `--print before_dl:description`
const Plaintext_Tags string = `--print before_dl:"Tags: %(tags)s"`                      //Changed for patching data fields. `--print before_dl:tags`
const Plaintext_Categories string = `--print before_dl:"Categories: %(categories)s"`    //Changed for patching data fields. `--print before_dl:categories`
const Plaintext_Extension string = `--print before_dl:ext`
const Plaintext_Duration string = `--print before_dl:duration`
const Plaintext_URLDomain string = `--print before_dl:webpage_url_domain`
const Plaintext_OriginalURL string = `--print before_dl:original_url`
const Plaintext_PlaylistTitle string = `--print before_dl:playlist_title`
const Plaintext_PlaylistCount string = `--print before_dl:playlist_count`
const Plaintext_PlaylistIndex string = `--print before_dl:playlist_index`
const Plaintext_YTFormatString string = `--print before_dl:format`
const Plaintext_FileSizeApprox string = `--print before_dl:filesize_approx`
const Plaintext_FormatNote string = `--print before_dl:format_note`
const Plaintext_Resolution string = `--print before_dl:resolution`
const Plaintext_ChannelId string = `--print before_dl:channel_id`
const Plaintext_ChannelURL string = `--print before_dl:channel_url`
const Plaintext_PlaylistId string = `--print before_dl:playlist_id`

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
const OutputPlaylistThumbnailFile string = `-o "%(webpage_url_domain)s/%(channel)s/%(playlist)s/Thumbnails/%(playlist_index)s - %(title)s [%(id)s].%(ext)s"`

// Videos: Videos, Subtitles, Thumbnails
const OutputVideoFile string = `-o "%(webpage_url_domain)s/%(channel)s/Videos/%(title)s [%(id)s].%(ext)s"`
const OutputSubtitleFile string = `-o "subtitle:%(webpage_url_domain)s/%(channel)s/Videos/Subtitles/%(title)s [%(id)s].%(ext)s"`
const OutputThumbnailFile string = `-o "thumbnail:%(webpage_url_domain)s/%(channel)s/Videos/Thumbnails/%(title)s [%(id)s].%(ext)s"`

// Output Parsing: Warnings & Errors
const WARNING string = `WARNING:`
const ERROR string = `ERROR:`
const ANSWER_START string = `{`

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
