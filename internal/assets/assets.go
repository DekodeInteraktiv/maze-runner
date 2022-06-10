package assets

import "embed"

// content holds our static web server content.
//go:embed static viewer controller leaderboard register
var Content embed.FS
