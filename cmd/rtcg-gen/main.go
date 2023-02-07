// Program rtcg-gen performs test code generation given JSON input.
package main

import (
	"os"
	"rtcg/internal/cli"
	"rtcg/internal/gen"
)

func main() {
	cli.HandleError(run(), "TEMPLATE-DIR [INPUT-FILE]")
}

func run() error {
	g, err := parseArgs(os.Args)
	if err != nil {
		return err
	}

	return g.Run()
}

func parseArgs(args []string) (*gen.Generator, error) {
	var g gen.Generator

	if len(args) == 2 {
		g.InputFile = "-" // stdin
	} else if len(args) == 3 {
		g.InputFile = args[2]
	} else {
		return nil, cli.ErrBadArgs
	}

	g.TemplateDir = args[1]
	return &g, nil
}
