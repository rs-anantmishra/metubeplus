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
	,FileId
	,ChannelId
	,PlayListId
	,DomainId
	,FormatId
	,YoutubeVideoId
	,WatchCount
	,IsDeleted
	,CreatedDate
)
VALUES (NULL, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);`

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
const InsertFile string = `INSERT INTO tblFiles SELECT NULL, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?`

//Files with duplicate names to be overwritten or renamed(Choice thru UI).
const InsertMediaFileCheck string = `SELECT Id From tblFiles WHERE FileType = ? AND VideoId = ? AND FileName = ?`

const GetNetworkVideoURLById string = `Select WebpageURL from tblVideos Where Id = ?`
const GetVideoInformationById string = `Select V.Title, V.PlayListId, C.Name as 'Channel', D.Domain, P.Title as 'PlaylistTitle', YoutubeVideoId
										FROM tblVideos V 
										INNER JOIN tblChannels C ON C.Id = V.ChannelId 
										INNER JOIN tblDomains D ON D.Id = V.DomainId
										INNER JOIN tblPlaylists P ON P.Id = V.PlayListId
										WHERE V.Id = ?;`

const UpdateVideoFileFields string = `UPDATE tblVideos SET IsFileDownloaded = ?, FileId = ? WHERE Id = ?;`

const GetAllVideos_Info string = `Select V.Id, V.Title, V.Description, V.DurationSeconds, V.WebpageURL, V.IsFileDownloaded, V.IsDeleted, C.Name, P.Title, V.LiveStatus, D.Domain, V.Availability, F.Format, V.YoutubeVideoId, V.CreatedDate,  FI.FilePath || '\' || FI.FileName as 'ThumbnailFilePath'
								  FROM tblVideos V
								  INNER JOIN tblChannels C ON V.ChannelId = C.Id
								  INNER JOIN tblPlaylists P ON V.PlaylistId = P.Id
								  INNER JOIN tblDomains D ON V.DomainId = D.Id
								  INNER JOIN tblFormats F ON V.FormatId = F.Id
								  INNER JOIN tblFiles FI ON (V.Id = FI.VideoId AND FI.FileType = 'Thumbnail') LIMIT 25;`
