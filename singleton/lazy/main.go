package main

import (
	"fmt"
	"lazy/log"
)

func main() {
	log := log.NewLogger(log.Log)
	fmt.Println(log.Timestamp)
}
