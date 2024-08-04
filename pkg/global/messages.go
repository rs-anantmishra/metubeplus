package global

import (
	"sync"

	"github.com/gofiber/fiber/v2/log"
)

var onceDS sync.Once
var onceAI sync.Once

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
	onceAI.Do(func() { // <-- atomic, does not allow repeating
		log.Info("DURING ONCE 1", activeItem)
		activeItem = make([]DownloadStatus, 0) // <-- thread safe
		log.Info("DURING ONCE 2", activeItem)
	})

	return activeItem
}

func NewDownloadStatus() []DownloadStatus {
	onceDS.Do(func() { // <-- atomic, does not allow repeating
		log.Info("DURING ONCE 1 Active ITEM", dsQueue)
		dsQueue = make([]DownloadStatus, 0) // <-- thread safe
		log.Info("DURING ONCE 2 Active ITEM", dsQueue)
	})

	return dsQueue
}

const (
	Queued      = iota
	Downloading = iota
	Completed   = iota
)
