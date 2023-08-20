package migrations

import (
	"embed"
)

//go:embed *.sql
var EmbedMigratonFS embed.FS
