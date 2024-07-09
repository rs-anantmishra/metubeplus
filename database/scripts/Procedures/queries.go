// This file to be replaced with .sql files for each query
package procedures

const InsertChannel string = `IF (Select YoutubeChannelId FROM tblChannels Where YoutubeChannelId = ?)
BEGIN 
	Select Id from tblChannels where YoutubeChannelId = ?;
END
ELSE 
BEGIN 

	INSERT INTO tblChannels
	Select NULL, ?, ?, ?, ?;
	SELECT last_insert_rowid();

END`

const InsertPlaylist string = `IF (Select YoutubePlaylistID FROM tblPlaylists WHERE YoutubePlaylistID = ?)
BEGIN 
	Select Id from tblPlaylists where YoutubePlaylistID = ?;
END
ELSE 
BEGIN 

	INSERT INTO tblPlaylists
	Select NULL, ?, ?, ?, ?, ?, ?, ?;
	SELECT last_insert_rowid();

END`

const InsertDomain string = `IF (Select Id FROM tblDomains Where Domain = ?)
BEGIN 
	Select Id from tblDomains where Domain = ?;
END
ELSE 
BEGIN 

	INSERT INTO tblDomains
	Select NULL, ?, ?;
	SELECT last_insert_rowid();

END`

const InsertFormat string = `IF (Select Id FROM tblFormats Where Format = ? AND FormatNote = ? AND Resolution = ? AND StreamType = ?)
BEGIN 
	Select Id from tblFormats where Format = ? AND FormatNote = ? AND Resolution = ? AND StreamType = ?;
END
ELSE 
BEGIN 

INSERT INTO tblFormats
Select NULL, ?, ?, ?, ?, ?;
SELECT last_insert_rowid();

END`

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

const InsertTags string = `IF (Select Id FROM tblTags Where Name = ?)
BEGIN 
	Select Id FROM tblTags Where Name = ?;
END
ELSE 
BEGIN 

	INSERT INTO tblTags
	Select NULL, ?, ?, ?;
	SELECT last_insert_rowid();

END`

const InsertCategories string = `IF (Select Id FROM tblCategories Where Name = ?)
BEGIN 
	Select Id FROM tblCategories Where Name = ?;
END
ELSE 
BEGIN 

	INSERT INTO tblCategories
	Select NULL, ?, ?, ?;
	SELECT last_insert_rowid();

END`
