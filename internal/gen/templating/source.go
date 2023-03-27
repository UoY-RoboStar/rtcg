package templating

import "io/fs"

// Source names a source of template files.
type Source struct {
	Name string // Name is the name of the source.
	Src  fs.FS  // Src is the source filesystem.
}
