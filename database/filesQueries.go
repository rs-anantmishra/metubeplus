package database

const GetStorageUsedInfo string = `Select SUM(FileSize) as 'Filesize' from tblFiles WHERE FileType = 'Video';`