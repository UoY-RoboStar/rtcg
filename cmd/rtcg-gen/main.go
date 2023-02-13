// Program rtcg-gen performs test code generation given a JSON representation of a state machine suite.
package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/UoY-RoboStar/rtcg/internal/cli"
	"github.com/UoY-RoboStar/rtcg/internal/gen"
	"github.com/UoY-RoboStar/rtcg/internal/stm"
)

const (
	usage            = "TEMPLATE-DIR [STM-FILE]"
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
	stms, err := a.readSuite()
	if err != nil {
		return fmt.Errorf("couldn't read state machine suite: %w", err)
	}

	return a.generate(stms)
}

func (a *genAction) readSuite() (stm.Suite, error) {
	file, err := cli.OpenFileOrStdin(a.inputFile)
	if err != nil {
		return nil, fmt.Errorf("couldn't open input: %w", err)
	}

	suite, err := stm.ReadSuite(file)

	return suite, errors.Join(err, file.Close())
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
