package global

import "sync"

var once sync.Once

// type global
//a simple single string
type messages []string

var (
	instance messages
)

func NewMessage() messages {
	once.Do(func() { // <-- atomic, does not allow repeating
		instance = make(messages, 1) // <-- thread safe
	})
	return instance
}

//isInvoked checks if process is already triggered
type isInvoked []bool

var (
	invokedInstance isInvoked
)

func NewIsInvoked() isInvoked {
	once.Do(func() { // <-- atomic, does not allow repeating
		invokedInstance = make(isInvoked, 1) // <-- thread safe
	})
	return invokedInstance
}

//handle case where Video is already added - Skip download.
//If Video is downloaded and is part of playlist then,
//add mapping for playlist in db but skip download.
type DownloadStatus struct {
	VideoId       int    //Local VideoId
	StatusMessage string //completion percentage is in here
	State         string //0 = downloading, 1 = queued, 2 = Downloaded
}

var (
	dsInstance []DownloadStatus
)

func NewDownloadStatus() []DownloadStatus {
	once.Do(func() { // <-- atomic, does not allow repeating
		dsInstance = make([]DownloadStatus, 0) // <-- thread safe
	})
	return dsInstance
}
