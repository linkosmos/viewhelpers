package viewhelpers

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestLoadtime(t *testing.T) {
	ti := time.Now().Add(-3 * time.Second)
	assert.Equal(t, 3000, Loadtime(ti))
}
