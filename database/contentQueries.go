package database

const GetVideoSearchChannelTitle string = `Select DISTINCT V.Id 'VideoId', V.Title, C.Name as 'Channel'
FROM tblVideos V
INNER JOIN tblChannels C ON V.ChannelId = C.Id
INNER JOIN tblPlaylistVideoFiles PVF ON V.Id = PVF.VideoId
INNER JOIN tblPlaylists P ON PVF.PlaylistId = P.Id
INNER JOIN tblDomains D ON V.DomainId = D.Id
INNER JOIN tblFormats F ON V.FormatId = F.Id
INNER JOIN tblFiles FIThumbnail ON (V.Id = FIThumbnail.VideoId AND FIThumbnail.FileType = 'Thumbnail')
INNER JOIN tblFiles FIVideo ON (V.Id = FIVideo.VideoId AND FIVideo.FileType = 'Video') ORDER BY V.Id ASC;`

// Playlists Page comes from here
const GetAllPlaylists string = `SELECT DISTINCT P.Id, P.Title, P.PlaylistUploader, P.ItemCount, P.YoutubePlaylistId, F.FilePath || '\' || F.FileName as 'ThumbnailURL'
FROM tblPlaylists P
INNER JOIN tblPlaylistVideoFiles PVF ON PVF.PlaylistId = P.Id
INNER JOIN tblVideos V ON V.Id = PVF.VideoId
INNER JOIN tblFiles F ON F.VideoId = PVF.VideoId
WHERE P.Id > 0 
	AND PVF.PlaylistVideoIndex = 1 
	AND F.FileType = 'Thumbnail';`
