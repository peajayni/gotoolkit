package timestamps

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestUTCTimestampsSmoke(t *testing.T) {
	ts := NewUTCTimestamps()
	now := ts.New()
	assert.Equal(t, time.UTC, now.Location())
}
