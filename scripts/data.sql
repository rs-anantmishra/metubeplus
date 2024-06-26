INSERT INTO tblLiveStatusType (LiveStatus)
VALUES('not_live') UNION 
VALUES('is_live') UNION 
VALUES('is_upcoming') UNION
VALUES('was_live')UNION
VALUES('post_live');

INSERT INTO tblAvailabilityType (Availability)
VALUES('private') UNION 
VALUES('premium_only') UNION 
VALUES('subscriber_only') UNION 
VALUES('needs_auth') UNION 
VALUES('unlisted') UNION 
VALUES('public');

INSERT INTO tblSourceType (Source)
VALUES('Downloaded') UNION 
VALUES('Uploaded') UNION 
VALUES('Local') UNION
VALUES('Metadata');

INSERT INTO tblFileType (File)
VALUES('Audio') UNION 
VALUES('Video') UNION 
VALUES('Thumbnail') UNION 
VALUES('Subtitles');


--Take time to define these properly.
--This should detail user actions that can be reviewed later.
INSERT INTO tblActivityType (File), 1
VALUES('Get Metadata for Audio', 1, 1) UNION 
VALUES('Get Metadata for Video', 1, 1) UNION 
VALUES('Get Metadata for Playlist', 1, 1) UNION 
VALUES('Get Metadata & Download Video', 1, 1) UNION 
VALUES('Get Metadata & Download Audio', 1, 1) UNION 
VALUES('Get Metadata & Download Playlist', 1, 1) UNION
VALUES('Get Metadata & Download Thumbnail', 1, 1) UNION 
VALUES('Get Metadata & Download Subtitles', 1, 1) UNION 
VALUES('Delete Audio', 0, 1) UNION 
VALUES('Delete Video', 0, 1) UNION 
VALUES('Delete Playlist', 0, 1) UNION 
VALUES('Create Local Playlist', 0, 1) UNION 
VALUES('Stream Video', 0, 1) UNION 
VALUES('Stream Audio', 0, 1) UNION 
VALUES('Stream Playlist', 0, 1) UNION 
VALUES('Stream Local Playlist', 0, 1);




