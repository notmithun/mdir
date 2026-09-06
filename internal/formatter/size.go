package formatter

import "fmt"

func FormatSize(bytes int64) string {
	const (
		KB = 1000
		MB = 1000000
		GB = 1000000000
		TB = 1000000000000
	)

	switch {
	case bytes < KB:
		return fmt.Sprintf("%d B", bytes)

	case bytes < MB:
		return fmt.Sprintf("%.3f KB (%d B)", float64(bytes)/KB, bytes)

	case bytes < GB:
		return fmt.Sprintf("%.3f MB (%d B)", float64(bytes)/MB, bytes)

	case bytes < TB:
		return fmt.Sprintf("%.3f GB (%d B)", float64(bytes)/GB, bytes)

	default:
		return fmt.Sprintf("%.3f TB (%d B)", float64(bytes)/TB, bytes)
	}
}
