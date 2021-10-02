// import "rss.app/cli"
package cli

import (
	"flag"

	"rss.app/web"
)

const (
	flagVersionHelp = "Show application version"
)

func Init() {
	var (
		flagVersion bool
	)

	flag.BoolVar(&flagVersion, "version", false, flagVersionHelp)

	flag.Parse()

	if flagVersion {
		version()
		return
	}

	version()
	web.Start()
}
