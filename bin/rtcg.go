package main

import (
	"errors"
	"fmt"
	"os"
	"test-code-gen/internal/gen"
)

func main() {
	if err := run(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		if errors.Is(err, ErrBadArgs) {
			usage()
		}
	}
}

func run() error {
	g, err := parseArgs(os.Args)
	if err != nil {
		return err
	}

	return g.Run()
}

func parseArgs(args []string) (*gen.Generator, error) {
	if len(args) != 3 {
		return nil, ErrBadArgs
	}
	return &gen.Generator{
		TemplateDir: args[1],
		InputFile:   args[2],
	}, nil
}

func usage() {
	_, _ = fmt.Fprintf(os.Stderr, "usage: %s TEMPLATE-DIR INPUT-FILE", os.Args[0])
}

var ErrBadArgs = errors.New("invalid arguments supplied")
