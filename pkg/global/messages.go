package global

import (
	"sync"
)

var onceDownloadStatus sync.Once
var onceActiveItem sync.Once
var onceQueueAlive sync.Once

// -- todo --------------------------------------------------
// handle case where Video is already added - Skip download.
// If Video is downloaded and is part of playlist being downloaded then,
// video will be added twice.
// -- todo --------------------------------------------------

// Singletons

type DownloadStatus struct {
	VideoId       int    //Local VideoId
	VideoURL      string //Network Video Id
	StatusMessage string //completion percentage is in here
	State         int    //0 = downloading, 1 = queued, 2 = Downloaded
	// VideoTitle    string //VideoTitle
}

var (
	dsQueue    []DownloadStatus
	activeItem []DownloadStatus
	queueAlive []int
)

func NewActiveItem() []DownloadStatus {
	onceActiveItem.Do(func() { // <-- atomic, does not allow repeating
		activeItem = make([]DownloadStatus, 1) // <-- thread safe
	})

	return activeItem
}

func NewDownloadStatus() []DownloadStatus {
	onceDownloadStatus.Do(func() { // <-- atomic, does not allow repeating
		dsQueue = make([]DownloadStatus, 0) // <-- thread safe
	})

	return dsQueue
}

func NewQueueAlive() []int {
	onceQueueAlive.Do(func() { // <-- atomic, does not allow repeating
		queueAlive = make([]int, 1) // <-- thread safe
	})

	return queueAlive
}

const (
	Queued      = iota
	Downloading = iota
	Completed   = iota
)
