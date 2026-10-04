// Package assets embeds the YAML content files served by the TUI, so the
// compiled binary is self-contained and needs no files on disk at runtime.
package assets

import "embed"

// Files holds every YAML file in this directory.
//
// Go's //go:embed directive only reaches files next to (or below) the source
// file that declares it. Keeping the YAML here, next to this file, is what lets
// the rest of the app import it as an embedded filesystem instead of relying on
// the process working directory.
//
//go:embed *.yaml
var Files embed.FS
