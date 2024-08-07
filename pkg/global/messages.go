package global

import (
	"sync"

	"github.com/gofiber/fiber/v2/log"
)

var onceDownloadStatus sync.Once
var onceActiveItem sync.Once

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
	dsQueue    []DownloadStatus
	activeItem []DownloadStatus
)

func NewActiveItem() []DownloadStatus {
	onceActiveItem.Do(func() { // <-- atomic, does not allow repeating
		activeItem = make([]DownloadStatus, 1) // <-- thread safe
		log.Info("During ActiveItem", activeItem)
	})

	return activeItem
}

func NewDownloadStatus() []DownloadStatus {
	onceDownloadStatus.Do(func() { // <-- atomic, does not allow repeating
		dsQueue = make([]DownloadStatus, 0) // <-- thread safe
		log.Info("During DownloadStatus", dsQueue)
	})

	return dsQueue
}

const (
	Queued      = iota
	Downloading = iota
	Completed   = iota
)
