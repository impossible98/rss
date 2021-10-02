// import "rss.app/cli"
package cli

import (
	"rss.app/constant"
	"rss.app/logger"
)

func version() {
	logger.Info(`%s Version: %s`, constant.App, constant.Version)
}
