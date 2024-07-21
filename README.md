# metubeplus

APPLICATION BEHAVIOR:
If you have an entry for the video in SQLite and you are attempting to download the video again, it asks to save the video as - 
1. A newer version and keep both the old and the new.
2. Download again and overwrite the existing video file.
3. Refresh metadata and keep the existing video file as it is.
	a. Show a warning - If the video was part of an existing playlist that the Playlist binding will be removed - keep that binding?


- hidden folders can be implemented.
- Should be able to refresh video information
- by default hide deleted Videos (option to show deleted Videos / faded / not-clickable)
- total space used by Video/playlist/channel
- manually retry download
- manually refresh metadata for videos
- auto update yt-dlp
- auto refresh media metadata 24hours (configurable) // to have the database in sync with fs and ensure the Video is not manually deleted, and if it is then library will reflect the current state of files. Bottom Line, source of truth for application should be the fs and not the app-db.
- file download and download with subtitle options
- Change files location:
	- change download location but keep current stuff as it is.
	- change download location and move all current stuff to this new location.
	- keep the download location the same but move all current files to a new location.
	- I did it manually and now Videos are not available in application (Fix it!)
	- I did it manually and now Videos are not available in application (Remove unavailable videos!)
- edit video description/tags for local media files
- Upload Media to Local Folder
- backup db
- option to download db so that its manually backed up and restorable.
- custom names
- custom command
- url extraction on page / page change pattern with start and end index
- add extracted urls to a note and save it with a name 
- the name is a required field to Upload Video.
- change db to litestream at a later point.
- Save executed commands. log all failed commands and provide options to show it. (Implement a basic activity Viewer)
- add videos through uploading a file with links
- add videos list (; separated values)
- add single video url
- So the player has option to download videos, access local video via http, upload videos, and then another mode will be where Video will be added to the library but metadata only, but for playing it, web will be accessed as plyr also supports this kind of thing.
- Option to download subtitles at a later point should be provided in UI
- Get Subtitles in a different language
- Reasonable implementation of activity log is important.
- Implement litestream instead of sqlite3
- Only Show those tags and categories on UI which have more than 1 videos attached to it.
- Have option to disable APILogs & ActivityLogs
- Parallel Downloads (4 videos at a time.)
- Add and unwatched OR watched badge on Videos like wallhaven does.
- Add Functionality for NetworkPath for Files table later.
- Write Code to download yt-dlp and check for the executable before running a download.
- Rules can be applied to have videos of certain domains moved to hidden folders directly.
- Toralize connections through yt-dlp (bypass blocks) (--proxy socks5://localhost:9150 + [torproxy](https://hub.docker.com/r/dperson/torproxy/) )
- Local playlists can have audios, videos or both from different sources - local, downloaded, metaOnly, online.
- Download at a later time.
- Bookmark favorites.
- Fetch title for queued downloads.

[UI layout is similar to briefkasten]
- landing page shows all Videos
	- Downloader
	- All Videos (Scroll further down)
	- Channels
	- Downloaded
	- Local
///////////////////////////////////////////////////////////


tblVideos
	Id
	Title
	Description
	Duration
	WebpageURL
	IsFileDownloaded (Step1 Get Video Info, Step2 Download it. If Step2 Fails This is useful.)	
	IsDeleted
	VideoFormatId
	AvailabilityId
	PlayListId
	ChannelId
	LiveStatusId
	DomainId	
	VideoFileId	
	ThumbnailFileId
	SubtitlesFileId
	YoutubeVideoId
	CreatedDate

tblChannels
	Id
	Title
	WebpageURL
	YoutubeChannelId
	CreatedDate
	
tblLiveStatusType
	Id
	LiveStatus [One of "not_live", "is_live", "is_upcoming", "was_live", "post_live" (was live, but VOD is not yet processed)]
	
tblAvailabilityType
	Id
	Availability ("private", "premium_only", "subscriber_only", "needs_auth", "unlisted" or "public")

tblPlayList
	Id
	Title
	VideoCount
	ChannelId
	YoutubePlaylistID
	CreatedDate
	
tblDomains
	Id
	Domain (webpage_url_domain, local)
	CreatedDate

//IsUsed updated when video is added or removed.
tblTags
	Id
	Tag
	IsUsed
	CreatedDate

tblVideoTags
	Id
	VideoId
	TagId

tblFiles
	Id
	FileTypeId
	SourceId
	FilePath
	FileName
	Extension
	FileSize
	Resolution
	ParentDirectory
	IsDeleted
	CreatedDate
	
tblSourceType
	Id
	Source (Downloaded, Uploaded, Local, Metadata)

tblFileType
	Id
	File (Audio, Video, Thumbnail, Subtitles)
	
tblFormatType
	Id
	Format
	StreamId 

tblStreamType
	Id
	Stream (Audio, Video)
	
tblVideoFormat
	Id
	VideoId
	FormatId
	


//////////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////

metubeplus
	Config -> .env file
		yt_media_directory_location(s) [; separated]
		local_media_directory_location(s) [; separated]
		backup_db_location [mounted on container, preferably on another disk]
	Downloads
		Audio
			Domain
				Channels
					PlaylistName
						<FileName>	(Playlisted files live inside a Folder)
					<FileName> (Audio files with no playlist are within the Channels Folder)
		Video
			Domain
				Channels
					PlaylistName
						Thumbnails
						Subtitles	
						<FileName>	
					<FileName> (Videos with no playlist are within the Channels Folder)
					Thumbnails
					Subtitles	

		Thumbnails
		Subtitless
	Uploaded
		Audio
		Video	
		Thumbnails
	Local
		Thumbnails
