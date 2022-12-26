package log

import "time"

var Logger *logger = &logger{Timestamp: time.Now()}

type logger struct {
	Timestamp time.Time
}
