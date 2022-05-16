package assets

import "embed"

// content holds our static web server content.
//go:embed static
var Content embed.FS
