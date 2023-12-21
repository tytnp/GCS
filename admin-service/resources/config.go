package resources

import "embed"

//go:embed *.yaml
var ConfigFs embed.FS
