package database

// This file to be replaced with .sql files for each query
const InsertChannelCheck string = `Select Id From tblChannels Where YoutubeChannelId = ?`
const InsertChannel string = `INSERT INTO tblChannels Select NULL, ?, ?, ?, ?, ?;`

const InsertPlaylistCheck string = `Select Id From tblPlaylists WHERE YoutubePlaylistId = ?`
const InsertPlaylist string = `INSERT INTO tblPlaylists Select NULL, ?, ?, ?, ?, ?, ?, ?, ?, ?;`

const InsertPlaylistVideosCheck string = `Select Id From tblPlaylistVideoFiles WHERE PlaylistId = ? AND VideoId = ?`
const InsertPlaylistVideos string = `INSERT INTO tblPlaylistVideoFiles(Id, VideoId, PlaylistId, PlaylistVideoIndex, CreatedDate) Select NULL, ?, ?, ?, ?;`

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
	,OriginalURL
	,WebpageURL
	,LiveStatus
	,Availability
	,YoutubeViewCount
	,LikeCount
	,DislikeCount
	,License
	,AgeLimit
	,PlayableInEmbed
	,UploadDate
	,ReleaseTimestamp
	,ModifiedTimestamp
	,IsFileDownloaded
	,FileId
	,ChannelId
	,DomainId
	,FormatId
	,YoutubeVideoId
	,WatchCount
	,IsDeleted
	,CreatedDate
)
VALUES (NULL, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);`

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

const GetNetworkVideoURLById string = `Select WebpageURL from tblVideos Where Id = ?`
const GetVideoInformationById string = `Select V.Title, P.Title, C.Name as 'Channel', D.Domain, P.Title as 'PlaylistTitle', YoutubeVideoId, V.WebpageURL
										FROM tblVideos V 
										INNER JOIN tblChannels C ON C.Id = V.ChannelId 
										INNER JOIN tblDomains D ON D.Id = V.DomainId
										INNER JOIN tblPlaylistVideoFiles PVF ON PVF.VideoId = V.Id
										INNER JOIN tblPlaylists P ON P.Id = PVF.PlaylistId
										WHERE V.Id = ?;`

const UpdateVideoFileFields string = `UPDATE tblVideos SET IsFileDownloaded = ?, FileId = ? WHERE Id = ?;`
const UpdatePVFFileId string = `UPDATE tblPlaylistVideoFiles SET FileId = ? WHERE VideoId = ?`

//below may not be used
const UpdatePVFThumbnailFileId string = `Select F.Id
										 From tblFiles F
										 INNER JOIN tblPlaylistVideoFiles PVF ON PVF.VideoId = F.VideoId
										 WHERE PVF.PlaylistId = ? AND F.FileType = 'Thumbnail'
										 AND F.VideoId = ?`

const GetAllVideos_Info string = `Select DISTINCT V.Id, V.Title, V.Description, V.DurationSeconds, V.OriginalURL, V.WebpageURL, V.IsFileDownloaded, V.IsDeleted, C.Name, V.LiveStatus, D.Domain, V.LikeCount, V.YoutubeViewCount as 'ViewsCount', V.WatchCount, V.UploadDate, V.Availability, F.Format, V.YoutubeVideoId, V.CreatedDate
								  FROM tblVideos V
								  INNER JOIN tblChannels C ON V.ChannelId = C.Id
								  INNER JOIN tblPlaylistVideoFiles PVF ON V.Id = PVF.VideoId
								  INNER JOIN tblPlaylists P ON PVF.PlaylistId = P.Id
								  INNER JOIN tblDomains D ON V.DomainId = D.Id
								  INNER JOIN tblFormats F ON V.FormatId = F.Id
								  INNER JOIN tblFiles FIThumbnail ON (V.Id = FIThumbnail.VideoId AND FIThumbnail.FileType = 'Thumbnail')
								  INNER JOIN tblFiles FIVideo ON (V.Id = FIVideo.VideoId AND FIVideo.FileType = 'Video') ORDER BY V.Id ASC;`

const GetVideoTags_AllVideos string = `Select V.Id, T.Name, T.IsUsed, T.CreatedDate
							 FROM tblVideos V
							 INNER JOIN tblVideoFileTags VFT ON V.Id = VFT.VideoId
							 INNER JOIN tblTags T ON T.Id = VFT.TagId;`

const GetVideoCategories_AllVideos string = `Select V.Id, C.Name, C.IsUsed, C.CreatedDate
											 FROM tblVideos V
											 INNER JOIN tblVideoFileCategories VFC ON V.Id = VFC.VideoId
											 INNER JOIN tblCategories C ON C.Id = VFC.CategoryId;`

const GetVideoFiles_AllVideos string = `Select V.Id, F.FileType, F.FileSize, F.Extension, F.FilePath, F.FileName
										FROM tblVideos V
										INNER JOIN tblFiles F ON V.Id = F.VideoId;`

const GetQueuedVideoDetailsById string = `Select V.Id, V.Title, C.Name as 'Channel', V.Description, V.DurationSeconds as 'Duration', V.WebpageURL, FIThumbnail.FilePath || '\' || FIThumbnail.FileName as 'Thumbnail'
										  FROM tblVideos V
										  INNER JOIN tblChannels C ON V.ChannelId = C.Id
										  INNER JOIN tblPlaylists P ON V.PlaylistId = P.Id
										  INNER JOIN tblFiles FIThumbnail ON (V.Id = FIThumbnail.VideoId AND FIThumbnail.FileType = 'Thumbnail')
										  WHERE V.Id = ?`
