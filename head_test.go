package viewhelpers

import (
	"html/template"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStylesheetTag(t *testing.T) {
	expected := template.HTML(`<link type="text/css" rel="stylesheet" href="/assets/stylesheets/test.css">`)
	assert.Equal(t, expected, StylesheetTag("test.css"))
}
