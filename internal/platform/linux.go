//go:build linux

package platform

import (
	"os"
	"strings"
)

func IsHidden(entry os.DirEntry) bool {
	return strings.HasPrefix(entry.Name(), ".")
}
