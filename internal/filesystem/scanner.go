package filesystem

import "os"

type Scanner struct{}

func NewScanner() *Scanner {
	return &Scanner{}
}

func (s *Scanner) Scan(path string) ([]Entry, error) {
	dirEntries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	entries := make([]Entry, 0, len(dirEntries))

	for _, dirEntry := range dirEntries {
		entry := Entry{
			Name: dirEntry.Name(),
		}

		if dirEntry.IsDir() {
			entry.Type = Directory
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