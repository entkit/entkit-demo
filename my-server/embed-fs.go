package main

import "embed"

//go:embed refine-project
var RefineProjectFS embed.FS

//go:embed other-refine-project
var OtherRefineProjectFS embed.FS
