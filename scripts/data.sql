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
VALUES('Local');

INSERT INTO tblFileType (File)
VALUES('Audio') UNION 
VALUES('Video') UNION 
VALUES('Thumbnail') UNION 
VALUES('Subtitles');

-- INSERT INTO tblStreamType (Stream)
-- VALUES('Audio') UNION 
-- VALUES('Video');
