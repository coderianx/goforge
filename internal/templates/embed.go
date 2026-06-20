package templates

import "embed"

//go:embed gin/** fiber/** chi/** echo/** gorillamux/** stdlib/**
var FS embed.FS
