# metubeplus

- hidden folders can be implemented.
- Should be able to refresh video information
- by default hide deleted Videos (option to show deleted Videos / faded / not-clickable)
- total space used by Video/playlist/channel
- manually retry download
- manually refresh metadata for videos
- auto update yt-dlp
- auto refresh media metadata 24hours (configurable)
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
	DomainId (determines its local, downloaded)
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
	Source (Downloaded, Uploaded, Local)

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
	Config
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
						<FileName>	(Playlist files live inside a Folder)
					<FileName> (Videos with no playlist are within the Channels Folder)
		Thumbnails
	Uploaded
		Audio
		Video	
		Thumbnails
	Local
		Thumbnails
