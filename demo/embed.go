package demo

import (
	"embed"
)

//go:embed openapi.yaml
var OpenApiFile embed.FS
