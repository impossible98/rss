// import "rss.app/web"
package web

import (
	"embed"
	"io/fs"
	"net/http"

	"rss.app/config"
	"rss.app/logger"
)

//go:embed public
var public embed.FS

func Start(port string) {
	address := config.Address

	logger.Info("RSS is serving on: http://%s:%s", address, port)
	http.Handle("/", http.FileServer(getFileSystem()))

	err := http.ListenAndServe(":"+port, nil)

	if err != nil {
		logger.Fatal("ListenAndServe fail:", err)
	}
}

func getFileSystem() http.FileSystem {
	fsys, err := fs.Sub(public, "public")
	if err != nil {
		panic(err)
	}

	return http.FS(fsys)
}
