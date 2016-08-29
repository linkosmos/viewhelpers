package viewhelpers

import (
	"fmt"
	"html/template"
)

// -
var (
	StylesheetTagPath = "/assets/stylesheets"
)

// StylesheetTag - generates StylesheetTag HEAD element
func StylesheetTag(href string) template.HTML {
	return template.HTML(
		fmt.Sprintf(`<link type="text/css" rel="stylesheet" href="%s/%s">`,
			StylesheetTagPath,
			href,
		),
	)
}
