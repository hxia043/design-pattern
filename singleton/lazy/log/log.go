package log

import (
	"sync"
	"time"
)

var Log *logger
var mu sync.Mutex

type logger struct {
	Timestamp time.Time
}

func NewLogger(log *logger) *logger {
	if log == nil {
		mu.Lock()
		defer mu.Unlock()

		if log == nil {
			log = &logger{Timestamp: time.Now()}
			return log
		}
	}

	return log
}
