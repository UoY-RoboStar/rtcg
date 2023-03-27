package templating

// File holds information about a templated file.
type File struct {
	Dir  string // Dir is the destination directory of this file, within the test directory.
	Name string // Name is the filename of this file.
	Desc string // Desc is a human-readable description for this file.
	Glob string // Glob is the slash-delimited glob of source templates for this file.
}
