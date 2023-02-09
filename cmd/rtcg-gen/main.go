// Program rtcg-gen performs test code generation given JSON input.
package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/UoY-RoboStar/rtcg/internal/cli"
	"github.com/UoY-RoboStar/rtcg/internal/gen"
	"github.com/UoY-RoboStar/rtcg/internal/stm"
	"github.com/UoY-RoboStar/rtcg/internal/testlang"
)

const (
	usage            = "TEMPLATE-DIR [INPUT-FILE]"
	numAnonymousArgs = 2
)

func main() {
	cli.HandleError(run(), usage)
}

func run() error {
	a, err := parseArgs(os.Args)
	if err != nil {
		return err
	}

	return a.run()
}

func parseArgs(argv []string) (*genAction, error) {
	var (
		action genAction
		err    error
	)

	action.templateDir = argv[1]

	action.inputFile, err = cli.ParseFileArgument(argv, numAnonymousArgs+1)
	if err != nil {
		return nil, fmt.Errorf("couldn't parse args: %w", err)
	}

	return &action, nil
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
	file, err := cli.OpenFileOrStdin(a.inputFile)
	if err != nil {
		return nil, fmt.Errorf("couldn't open input: %w", err)
	}

	suite, err := testlang.ReadSuite(file)

	return suite, errors.Join(err, file.Close())
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

	if err := g.Generate(stms); err != nil {
		return fmt.Errorf("couldn't generate: %w", err)
	}

	return nil
}
