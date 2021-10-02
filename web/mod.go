// import "rss.app/web"
package web

import (
	"net/http"

	"rss.app/config"
	"rss.app/logger"
)

func Start() {
	address := config.Address
	port := config.Port
	mux := http.NewServeMux()

	rh := http.RedirectHandler("https://www.bing.com", 307)
	mux.Handle("/", rh)

	logger.Info("RSS is serving on: http://%s:%s", address, port)
	http.ListenAndServe(":"+port, mux)
}
