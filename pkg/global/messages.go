package global

import (
	"sync"
)

var once sync.Once

// -- todo --------------------------------------------------
// handle case where Video is already added - Skip download.
// If Video is downloaded and is part of playlist then,
// add mapping for playlist in db but skip download.
// -- todo --------------------------------------------------

// Singletons

type DownloadStatus struct {
	VideoId       int    //Local VideoId
	VideoURL      string //Network Video Id
	StatusMessage string //completion percentage is in here
	State         int    //0 = downloading, 1 = queued, 2 = Downloaded
}

var (
	dsQueue []*DownloadStatus
)

func NewDownloadStatus() []*DownloadStatus {
	once.Do(func() { // <-- atomic, does not allow repeating
		dsQueue = make([]*DownloadStatus, 0) // <-- thread safe
	})
	return dsQueue
}

const (
	Queued      = iota
	Downloading = iota
	Completed   = iota
)
