package ids

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXidsSmoke(t *testing.T) {
	ids := NewXids()
	id := ids.New()
	assert.True(t, len(id) > 0)
}
