package web

import (
	"embed"
	"io/fs"
)

//go:embed all:out
var BuildFs embed.FS

// embed the production build of the FE located in the out directory.
func BuildFS() fs.FS {
	build, err := fs.Sub(BuildFs, "out")
	if err != nil {
		panic(err)
	}
	return build
}
