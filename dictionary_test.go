package viewhelpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDict(t *testing.T) {
	val1 := "value1"
	val2 := []string{"some", "values"}
	dict, err := Dict("key1", "value1", "key2", val2)

	assert.Nil(t, err)

	got, _ := dict["key1"]
	assert.Equal(t, val1, got)

	got, _ = dict["key2"]
	assert.Equal(t, val2, got)
}
