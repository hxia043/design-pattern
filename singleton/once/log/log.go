package log

import (
	"sync"
	"time"
)

var oncelogger sync.Once

var log *logger

type logger struct {
	Timestamp time.Time
}

func NewLogger() *logger {
	if log == nil {
		oncelogger.Do(func() {
			log = &logger{Timestamp: time.Now()}
		})

		return log
	}

	return log
}
