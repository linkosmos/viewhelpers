package viewhelpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLen(t *testing.T) {
	input := []string{"2", "3"}

	assert.Equal(t, len(input), Len(input))

	input2 := "invalid"
	assert.Equal(t, len(input2), Len(input2))
}
