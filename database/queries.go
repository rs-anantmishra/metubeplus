package database

// This file to be replaced with .sql files for each query
const InsertChannelCheck string = `Select Id From tblChannels Where YoutubeChannelId = ?`
const InsertChannel string = `INSERT INTO tblChannels Select NULL, ?, ?, ?, ?;`

const InsertPlaylistCheck string = `Select Id From tblPlaylists WHERE YoutubePlaylistID = ?`
const InsertPlaylist string = `INSERT INTO tblPlaylists Select NULL, ?, ?, ?, ?, ?, ?, ?;`

const InsertDomainCheck string = `Select Id From tblDomains Where Domain = ?`
const InsertDomain string = `INSERT INTO tblDomains Select NULL, ?, ?;`

const InsertFormatCheck string = `Select Id From tblFormats Where Format = ?`
const InsertFormat string = `INSERT INTO tblFormats Select NULL, ?, ?, ?, ?, ?;`

const InsertMetadataCheck string = `Select Id From tblVideos Where YoutubeVideoId = ?`
const InsertMetadata string = `INSERT INTO tblVideos (
	Id
	,Title
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

const InsertVideoFileTagsCheck string = `SELECT Id From tblVideoFileTags Where TagId = ? AND VideoId = ?`
const InsertVideoFileTags string = `INSERT INTO tblVideoFileTags SELECT NULL, ?, ?, ?, ?`

const InsertVideoFileCategoriesCheck string = `SELECT Id From tblVideoFileCategories Where CategoryId = ? AND VideoId = ?`
const InsertVideoFileCategories string = `INSERT INTO tblVideoFileCategories SELECT NULL, ?, ?, ?, ?`

const InsertThumbnailFileCheck string = `SELECT Id From tblFiles WHERE FileType = ? AND VideoId = ?`
const InsertSubsFileCheck string = `SELECT Id From tblFiles WHERE FileType = ? AND VideoId = ? AND FileName = ?`
const InsertFile string = `INSERT INTO tblFiles SELECT NULL, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?`

//Files with duplicate names to be overwritten or renamed(Choice thru UI).
const InsertMediaFileCheck string = `SELECT Id From tblFiles WHERE FileType = ? AND VideoId = ? AND FileName = ?`

const GetNetworkVideoIdByVideoId string = `Select WebpageURL from tblVideos Where Id = ?`
