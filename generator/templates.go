package generator

import _ "embed"

//go:embed templates/manifest.plist
var manifestTemplate []byte

//go:embed templates/index.html
var indexTemplate []byte
