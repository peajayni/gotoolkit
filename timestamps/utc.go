package timestamps

import (
	"time"
)

func NewUTCTimestamps() *UTCTimestamps {
	return &UTCTimestamps{}
}

type UTCTimestamps struct{}

func (f *UTCTimestamps) New() time.Time {
	return time.Now().UTC()
}
