package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

var oncelogger sync.Once
var onceenv sync.Once

var log *logger

type logger struct {
	timestamp time.Time
}

func NewLogger() *logger {
	if log == nil {
		oncelogger.Do(func() {
			log = &logger{timestamp: time.Now()}
		})

		return log
	}

	return log
}

func getEnv() {
	if os.Getenv("TEST") == "" {
		onceenv.Do(func() {
			if err := os.Setenv("TEST", "hxia"); err != nil {
				fmt.Println(err)
			}
		})
	}

	fmt.Println(os.Getenv("TEST"))
}

func main() {
	l := NewLogger()
	fmt.Println(l.timestamp.UTC())

	for i := 0; i < 10; i++ {
		go getEnv()
	}

	time.Sleep(1 * time.Second)
}
