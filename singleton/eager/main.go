package main

import (
	"fmt"
	"time"
)

var log *logger = &logger{Timestamp: time.Now()}

type logger struct {
	Timestamp time.Time
}

func main() {
	fmt.Println(log.Timestamp)
}
