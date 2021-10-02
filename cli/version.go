// import "rss.app/cli"
package cli

import (
	"fmt"

	"rss.app/constant"
)

func version() {
	fmt.Printf("%s v%s\n", constant.App, constant.Version)
}
