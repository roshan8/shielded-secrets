package app

import "embed"

//go:embed fe/build/**
var Frontend embed.FS
