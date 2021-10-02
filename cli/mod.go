// import "rss.app/cli"
package cli

import (
	"flag"

	"rss.app/web"
)

const (
	flagPortHelp    = "Show applicaion port"
	flagVersionHelp = "Show application version"
)

func Init() {
	var (
		flagPort    string
		flagVersion bool
	)

	flag.StringVar(&flagPort, "port", "8000", flagPortHelp)
	flag.BoolVar(&flagVersion, "version", false, flagVersionHelp)

	flag.Parse()

	if flagVersion {
		version()
		return
	}

	version()
	web.Start(flagPort)
}
