package main

import (
	"fmt"
	"sync"
	"time"
)

var log *logger
var mu sync.Mutex

type logger struct {
	timestamp time.Time
}

func NewLogger(log *logger) *logger {
	if log == nil {
		mu.Lock()
		defer mu.Unlock()

		if log == nil {
			log = &logger{timestamp: time.Now()}
			return log
		}
	}

	return log
}

func main() {
	logger := NewLogger(log)
	fmt.Println(logger.timestamp)
}
