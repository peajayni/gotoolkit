package timestamps

import (
	"time"
)

func NewUTCTimestamps() *UTCTimestamps {
	return &UTCTimestamps{}
}

type UTCTimestamps struct{}

func (f *UTCTimestamps) Make() time.Time {
	return time.Now().UTC()
}
