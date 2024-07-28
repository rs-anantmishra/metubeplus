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
