//go:generate go-bindata -pkg static -ignore .../.DS_Store -o files.go files/...

package static

import (
	"log"
	"net/http"
	"os"

	"github.com/elazarl/go-bindata-assetfs"
)

// all static/ files embedded as a Go library
func FileSystemHandler() http.Handler {
	var h http.Handler
	if info, err := os.Stat("static/files/"); err == nil && info.IsDir() {
		log.Printf("serving local static/files")
		h = http.FileServer(http.Dir("static/files/"))
	} else {
		log.Printf("serving embedded static files")
		h = http.FileServer(&assetfs.AssetFS{
			Asset:     Asset,
			AssetInfo: AssetInfo,
			AssetDir:  AssetDir,
			Prefix:    "files",
		})
	}
	return h
}
