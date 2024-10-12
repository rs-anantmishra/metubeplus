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
