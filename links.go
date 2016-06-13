package viewhelpers

import (
	"fmt"
	"html/template"
)

// OutboundLink - simple outbound link with nofollow and noreferrer attributes
func OutboundLink(href string) template.HTML {
	return template.HTML(
		fmt.Sprintf(`<a href="http://%s" target="_blank" rel="nofollow noopener noreferrer">%s</a>`, href, href),
	)
}
