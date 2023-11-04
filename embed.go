package main

import "embed"

//go:embed fe/build/**
var Frontend embed.FS
