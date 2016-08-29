package viewhelpers

import (
	"fmt"
	"html/template"
)

// -
var (
	StylesheetTagPath = "/assets/stylesheets"
	JavascriptTagPath = "/assets/javascripts"
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

// JavascriptTag - generates <script> element
func JavascriptTag(src string) template.HTML {
	return template.HTML(
		fmt.Sprintf(`<script src="%s/%s" async></script>`,
			JavascriptTagPath,
			src,
		),
	)
}
