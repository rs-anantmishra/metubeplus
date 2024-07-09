package database

// This file to be replaced with .sql files for each query
const InsertChannelCheck string = `Select YoutubeChannelId From tblChannels Where YoutubeChannelId = ?`
const InsertChannel string = `INSERT INTO tblChannels Select NULL, ?, ?, ?, ?;`

const InsertPlaylistCheck string = `Select YoutubePlaylistID From tblPlaylists WHERE YoutubePlaylistID = ?`
const InsertPlaylist string = `INSERT INTO tblPlaylists Select NULL, ?, ?, ?, ?, ?, ?, ?;`

const InsertDomainCheck string = `Select Id From tblDomains Where Domain = ?`
const InsertDomain string = `INSERT INTO tblDomains Select NULL, ?, ?;`

const InsertFormatCheck string = `Select Id From tblFormats Where Format = ? AND FormatNote = ? AND Resolution = ? AND StreamType = ?`
const InsertFormat string = `INSERT INTO tblFormats Select NULL, ?, ?, ?, ?, ?;`

const InsertMetadataCheck string = `Select Id From tblVideos Where YoutubeVideoId = ?`
const InsertMetadata string = `INSERT INTO tblVideos (
	Title
	,Description
	,DurationSeconds
	,WebpageURL
	,LiveStatus
	,Availability
	,PlaylistVideoIndex
	,IsFileDownloaded
	,ChannelId
	,PlayListId
	,DomainId
	,FormatId
	,YoutubeVideoId
	,WatchCount
	,IsDeleted
	,CreatedDate
)
VALUES (NULL, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);`

const InsertTagsCheck string = `Select Id From tblTags Where Name = ?`
const InsertTags string = `INSERT INTO tblTags Select NULL, ?, ?, ?;`

const InsertCategoriesCheck string = `Select Id From tblCategories Where Name = ?`
const InsertCategories string = `INSERT INTO tblCategories Select NULL, ?, ?, ?;`
