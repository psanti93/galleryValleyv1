package templates

import "embed"

// takes anything that matches *.gohtml pattern and puts it in the binary

//go:embed *.gohtml
var FS embed.FS
