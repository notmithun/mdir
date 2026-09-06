package formatter

import (
	"fmt"

	"github.com/notmithun/mdir/internal/filesystem"
)

func PrintEntries(entries []filesystem.Entry) {
	var fileCount int
	var directoryCount int
	var totalSize int64

	maxNameLength := 0

	for _, entry := range entries {
		if len(entry.Name) > maxNameLength {
			maxNameLength = len(entry.Name)
		}
	}

	for _, entry := range entries {
		switch entry.Type {
		case filesystem.File:
			size := FormatSize(entry.Size)

			fmt.Printf("%-*s%s\n", maxNameLength+4, entry.Name, size)

			fileCount++
			totalSize += entry.Size
		case filesystem.Directory:
			fmt.Printf("%-*s<DIR>\n", maxNameLength+4, entry.Name)

			directoryCount++
		}
	}

	fmt.Printf("\n\t\t%d Files (%s)\n", fileCount, FormatSize(totalSize))
	fmt.Printf("\t\t%d DIR\n", directoryCount)
}

func PrintTree(entries []filesystem.Entry) {
	for i, entry := range entries {
		last := i == len(entries)-1
		printTreeEntry(entry, "", last)
	}
}

func printTreeEntry(entry filesystem.Entry, prefix string, last bool) {
	connector := "├── "
	if last {
		connector = "└── "
	}

	switch entry.Type {
	case filesystem.File:
		fmt.Printf("%s%s%s\t%s\n",
			prefix,
			connector,
			entry.Name,
			FormatSize(entry.Size),
		)
	case filesystem.Directory:
		fmt.Printf("%s%s%s\t<DIR>\n",
			prefix,
			connector,
			entry.Name,
		)

		childPrefix := prefix + "│   "

		if last {
			childPrefix = prefix + "    "
		}

		for i, child := range entry.Children {
			childLast := i == len(entry.Children)-1

			printTreeEntry(
				child,
				childPrefix,
				childLast,
			)
		}
	}
}
