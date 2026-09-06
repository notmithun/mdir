package filesystem

type EntryType uint8

const (
	File EntryType = iota
	Directory
)

type Entry struct {
	Name string
	Type EntryType
	Size int64
}