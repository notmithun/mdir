//go:build windows

package platform

import (
	"os"
	"syscall"
)

func IsHidden(entry os.DirEntry) bool {
	info, err := entry.Info()
	if err != nil {
		return false
	}

	data, ok := info.Sys().(*syscall.Win32FileAttributeData)
	if !ok {
		return false
	}

	return data.FileAttributes&syscall.FILE_ATTRIBUTE_HIDDEN != 0
}
