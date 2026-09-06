package filesystem

import (
	"os"
	"path/filepath"

	"github.com/notmithun/mdir/internal/platform"
)

type ScanOptions struct {
	All       bool
	Recursive bool
}

type Scanner struct{}

func NewScanner() *Scanner {
	return &Scanner{}
}

func (s *Scanner) Scan(path string, options ScanOptions) ([]Entry, error) {
	return s.scanDirectory(path, options)
}

func (s *Scanner) scanDirectory(path string, options ScanOptions) ([]Entry, error) {
	dirEntries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	entries := make([]Entry, 0, len(dirEntries))

	for _, dirEntry := range dirEntries {
		if !options.All && platform.IsHidden(dirEntry) {
			continue
		}

		entry := Entry{
			Name: dirEntry.Name(),
		}

		if dirEntry.IsDir() {
			entry.Type = Directory
			if options.Recursive {
				childPath := filepath.Join(path, dirEntry.Name())

				children, err := s.scanDirectory(childPath, options)
				if err != nil {
					continue
				}

				entry.Children = children
			}
		} else {
			entry.Type = File

			info, err := dirEntry.Info()
			if err != nil {
				return nil, err
			}

			entry.Size = info.Size()
		}

		entries = append(entries, entry)
	}
	return entries, nil
}
