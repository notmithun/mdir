package main

import (
	"fmt"

	"github.com/notmithun/mdir/internal/filesystem"
	"github.com/notmithun/mdir/internal/formatter"
)

func main() {
	scanner := filesystem.NewScanner()

	entries, err := scanner.Scan(".")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	formatter.PrintEntries(entries)
}
