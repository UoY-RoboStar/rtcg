package cpp

import (
	"embed"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path"
	"path/filepath"
)

//go:embed embed/prelude/*
var prelude embed.FS

const (
	preludeDir   = "rtcg"          // preludeDir is the subdirectory of srcDir where we store the prelude.
	preludeMount = "embed/prelude" // preludeMount is the directory in prelude where the prelude is.
)

func (g *Generator) copyPrelude() error {
	ents, err := fs.ReadDir(prelude, preludeMount)
	if err != nil {
		return fmt.Errorf("couldn't find prelude in embedded files: %w", err)
	}

	for _, e := range ents {
		if err := g.copyPreludeFile(e.Name()); err != nil {
			return err
		}
	}

	return nil
}

func (g *Generator) copyPreludeFile(name string) error {
	src, err := prelude.Open(path.Join(preludeMount, name))
	if err != nil {
		return fmt.Errorf("couldn't open prelude file %q: %w", name, err)
	}

	return writeFile(src, filepath.Join(g.srcBaseDir, preludeDir, name))
}

func (g *Generator) copyLocalFile(name string) error {
	srcPath := filepath.Join(g.dirSet.Input, name)
	dstPath := filepath.Join(g.srcBaseDir, name)

	src, err := os.Open(srcPath)
	if err != nil {
		return fmt.Errorf("couldn't open file %q: %w", srcPath, err)
	}

	return writeFile(src, dstPath)
}

func writeFile(src io.ReadCloser, outFile string) error {
	dst, err := os.Create(outFile)
	if err != nil {
		err = fmt.Errorf("couldn't create file %q: %w", outFile, err)

		return errors.Join(err, src.Close())
	}

	if _, err = io.Copy(dst, src); err != nil {
		err = fmt.Errorf("couldn't copy file %q: %w", outFile, err)
	}

	return errors.Join(err, dst.Close(), src.Close())
}
