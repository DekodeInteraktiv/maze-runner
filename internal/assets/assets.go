package assets

import "embed"

// content holds our static web server content.
//go:embed static viewer
var Content embed.FS
