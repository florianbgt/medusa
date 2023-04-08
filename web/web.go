package web

import (
	"embed"
	"io/fs"
)

//go:embed all:dist
var BuildFs embed.FS

// embed the production build of the FE located in the out directory.
func BuildFS() fs.FS {
	build, err := fs.Sub(BuildFs, "dist")
	if err != nil {
		panic(err)
	}
	return build
}
