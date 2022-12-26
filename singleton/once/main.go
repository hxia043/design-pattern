package main

import (
	"fmt"
	"once/log"
	"os"
	"sync"
	"time"
)

var onceenv sync.Once

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
	l := log.NewLogger()
	fmt.Println(l.Timestamp.UTC())

	for i := 0; i < 10; i++ {
		go getEnv()
	}

	time.Sleep(1 * time.Second)
}
