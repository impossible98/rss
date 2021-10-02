// import "rss.app/constant"
package constant

import (
	"bytes"
	"os"
)

func readSecretFile(filename, fallback string) string {
	data, err := os.ReadFile(filename)
	if err != nil {
		return fallback
	}

	value := string(bytes.TrimSpace(data))
	if value == "" {
		return fallback
	}

	return value
}

var (
	App     = "RSS"
	Version = readSecretFile("VERSION", "1.0.0")
)
