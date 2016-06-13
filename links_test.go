package viewhelpers

import (
	"html/template"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOutboundLink(t *testing.T) {
	expected := template.HTML(`<a href="http://www.example.org" target="_blank" rel="nofollow noopener noreferrer">www.example.org</a>`)

	got := OutboundLink("www.example.org")

	assert.Equal(t, expected, got)
}
