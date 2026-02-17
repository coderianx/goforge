package templates

import "embed"

//go:embed gin/** fiber/** chi/**
var FS embed.FS
