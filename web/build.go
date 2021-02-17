package web

import (
	"embed"
	"io/fs"
	"net/http"
	"os"
	"path"
)

//go:embed build/*
var Assets embed.FS

type fsFunc func(name string) (fs.File, error)

func (f fsFunc) Open(name string) (fs.File, error) {
	return f(name)
}

func AssetHandler(prefix, root string) http.Handler {
	handler := fsFunc(func(name string) (fs.File, error) {

		// wrap the embedded file system so that requests to open the file
		// are prefixed with our root, which is the build dir
		assetPath := path.Join(root, name)

		f, err := Assets.Open(assetPath)
		if os.IsNotExist(err) {
			// not found errors suggest we are seeing a request for a
			// Javascript internal route, so serve index.html
			return Assets.Open("build/index.html")
		}

		return f, err
	})

	// wrap the http.Handler to strip the /web/ prefix from our file reads
	return http.StripPrefix(prefix, http.FileServer(http.FS(handler)))
}
