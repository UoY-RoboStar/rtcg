package gencommon

import "path/filepath"

// DirSet is a set of directories associated with a generator.
type DirSet struct {
	Input  string // inputDir is the input directory of the generator.
	Output string // outputDir is the output directory of the generator.
}

// Subdir gets a DirSet formed by appending subdir to the input and output paths of this DirSet.
func (d DirSet) Subdir(subdir string) DirSet {
	return DirSet{
		Input:  filepath.Join(d.Input, subdir),
		Output: filepath.Join(d.Output, subdir),
	}
}

// OutputPath appends file to the output directory of this directory.
func (d DirSet) OutputPath(file string) string {
	return filepath.Join(d.Output, file)
}
