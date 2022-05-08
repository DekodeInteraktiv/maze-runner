package assets

import "embed"

// content holds our static web server content.
//go:embed all:static
var Content embed.FS
