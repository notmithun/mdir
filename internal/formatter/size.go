package formatter

import "fmt"

func FormatSize(bytes int64) string {
	if bytes < 1000 {
		return fmt.Sprintf("%d B", bytes)
	}

	kb := float64(bytes) / 1000
	return fmt.Sprintf("%.3f KB (%d B)", kb, bytes)
}
