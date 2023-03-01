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

	dst, err := os.Create(filepath.Join(g.outputDir, srcDir, preludeDir, name))
	if err != nil {
		err = fmt.Errorf("couldn't create prelude file %q: %w", name, err)

		return errors.Join(err, src.Close())
	}

	_, err = io.Copy(dst, src)
	if err != nil {
		err = fmt.Errorf("couldn't copy prelude file %q: %w", name, err)
	}

	return errors.Join(err, dst.Close(), src.Close())
}
