// Program rtcg-gen performs test code generation given JSON input.
package main

import (
	"os"
	"rtcg/internal/cli"
	"rtcg/internal/gen"
)

func main() {
	cli.HandleError(run(), "TEMPLATE-DIR INPUT-FILE")
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
		return nil, cli.ErrBadArgs
	}
	return &gen.Generator{
		TemplateDir: args[1],
		InputFile:   args[2],
	}, nil
}
