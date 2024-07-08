package procedures

const UpsertChannel string = `IF EXISTS`
const InsertMetadata string = `INSERT INTO tblVideos (Title, Description, DurationSeconds, WebpageURL, PlaylistVideoIndex, IsFileDownloaded, ChannelId, PlayListId, LiveStatusId, DomainId, AvailabilityId, FormatId, YoutubeVideoId, IsDeleted, CreatedDate) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);`
