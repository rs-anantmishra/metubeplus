package videos

// --Testing----------------------------------------------------------------------------------------//
const TestURL1 string = `https://www.youtube.com/watch?v=GW2g-5WALrc`
const TestURL2 string = `https://www.youtube.com/watch?v=GW2g-5WALrc&list=PLFKeDWeuu3BZEBcRmolX6BDiFhK-GhCsd`
const TestURL3 string = `https://www.youtube.com/watch?v=-VC4FuG8P6Q`

// --Options---------------------------------------------------------------------------------------//
const ShowProgress string = `--progress`
const ProgressDelta string = `--progress-delta 1` //seconds
const Filepath string = `--print after_move:filepath`
const Channel string = `--print before_dl:channel`              //`--print after_move:channel`
const Title string = `--print before_dl:title`                  //`--print after_move:title`
const Description string = `--print before_dl:description`      //`--print after_move:description`
const Extension string = `--print before_dl:ext`                //`--print after_move:ext`
const Duration string = `--print before_dl:duration`            //`--print after_move:duration`
const URLDomain string = `--print before_dl:webpage_url_domain` //`--print after_move:webpage_url_domain`
const OriginalURL string = `--print before_dl:original_url`     //`--print after_move:original_url`
const PlaylistTitle string = `--print after_move:playlist_title`
const PlaylistCount string = `--print after_move:playlist_count`
const PlaylistIndex string = `--print after_move:playlist_index`
const AutoSubtitles string = `--write-auto-subs`

// --Extras----------------------------------------------------------------------------------------//
const WriteThumbnail string = `--write-thumbnail`
const YTFormatString string = `--print before_dl:format`
const SkipDownload string = `--skip-download`
const InfoJSON string = `--write-info-json`

// --Command---------------------------------------------------------------------------------------//
//command path should be picked from .env
const CommandName string = `yt-dlp_x86.exe`
const Space string = " "
const OutputFileNameSwitch string = `-o`
const OutputFileName string = `%(upload_date)s\%(title)s [%(id)s].%(ext)s`
