// Program rtcg-gen performs test code generation given JSON input.
package main

import (
	"errors"
	"fmt"
	"os"
	"rtcg/internal/cli"
	"rtcg/internal/gen"
	"rtcg/internal/stm"
	"rtcg/internal/testlang"
)

func main() {
	cli.HandleError(run(), "TEMPLATE-DIR [INPUT-FILE]")
}

func run() error {
	a, err := parseArgs(os.Args)
	if err != nil {
		return err
	}
	return a.run()
}

func parseArgs(argv []string) (*genAction, error) {
	var a genAction

	argc := len(argv)
	if argc == 2 {
		a.inputFile = "-" // stdin
	} else if argc == 3 {
		a.inputFile = argv[2]
	} else {
		return nil, cli.ErrBadArgs
	}

	a.templateDir = argv[1]
	return &a, nil
}

type genAction struct {
	templateDir string
	inputFile   string
}

func (a *genAction) run() error {
	s, err := a.readSuite()
	if err != nil {
		return fmt.Errorf("couldn't read test suite: %w", err)
	}

	stms := a.buildStms(s)
	// TODO: consider split STM building and generation?

	return a.generate(stms)
}

func (a *genAction) readSuite() (testlang.Suite, error) {
	f, err := cli.OpenFileOrStdin(a.inputFile)
	if err != nil {
		return nil, err
	}
	suite, err := testlang.ReadSuite(f)
	return suite, errors.Join(err, f.Close())
}

func (a *genAction) buildStms(tests testlang.Suite) stm.Suite {
	var bs stm.Builder
	return bs.BuildSuite(tests)
}

func (a *genAction) generate(stms stm.Suite) error {
	tmpls := os.DirFS(a.templateDir)
	g, err := gen.New(tmpls, "out") // for now
	if err != nil {
		return fmt.Errorf("couldn't create generator: %w", err)
	}
	return g.Generate(stms)
}
