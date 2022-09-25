package static

import (
	"embed"
)

// Embed static files
//go:embed web/*
var EmbedStaticWeb embed.FS

