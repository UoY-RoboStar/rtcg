// Program rtcg-gen performs test code generation given a JSON representation of a state machine suite.
package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/UoY-RoboStar/rtcg/internal/gen/config"

	"github.com/UoY-RoboStar/rtcg/internal/cli"
	"github.com/UoY-RoboStar/rtcg/internal/gen"
	"github.com/UoY-RoboStar/rtcg/internal/stm"
)

const numAnonymousArgs = 2

func main() {
	cli.HandleError(run(), usage)
}

func usage() {
	cli.FlagUsage("CONFIG-FILE [STM-FILE]")
}

func run() error {
	a, err := parseArgs()
	if err != nil {
		return err
	}

	return a.run()
}

func parseArgs() (*genAction, error) {
	var (
		action genAction
		err    error
	)

	flag.BoolVar(&action.clean, "clean", false, "remove output directory before generating")
	flag.StringVar(&action.outputDir, "output", "out", "output directory")
	flag.Usage = usage
	flag.Parse()

	argv := flag.Args()

	action.inputFile, err = cli.ParseFileArgument(argv, numAnonymousArgs)
	if err != nil {
		return nil, fmt.Errorf("couldn't parse args: %w", err)
	}

	// Parse this after action.inputFile so that we get an error if argv is empty
	action.configFile = argv[0]

	return &action, nil
}

type genAction struct {
	clean      bool   // clean makes the generator remove the output directory before generating.
	outputDir  string // outputDir is the output directory.
	configFile string
	inputFile  string
}

func (a *genAction) run() error {
	a.outputDir = filepath.Clean(a.outputDir)
	a.configFile = filepath.Clean(a.configFile)
	a.inputFile = filepath.Clean(a.inputFile)

	stms, err := a.readSuite()
	if err != nil {
		return fmt.Errorf("couldn't read state machine suite: %w", err)
	}

	if err := a.maybeClean(); err != nil {
		return err
	}

	return a.generate(stms)
}

func (a *genAction) maybeClean() error {
	if !a.clean {
		return nil
	}

	if a.outputDir == "." {
		return ErrCleanDot
	}

	if err := os.RemoveAll(a.outputDir); err != nil {
		return fmt.Errorf("couldn't clean %q: %w", a.outputDir, err)
	}

	return nil
}

// ErrCleanDot occurs if we try to clean "." (current directory).
//
// We don't support this because it would then change or confuse the current working directory.
var ErrCleanDot = errors.New("can't clean current directory")

func (a *genAction) readSuite() (stm.Suite, error) {
	file, err := cli.OpenFileOrStdin(a.inputFile)
	if err != nil {
		return nil, fmt.Errorf("couldn't open input: %w", err)
	}

	suite, err := stm.ReadSuite(file)

	return suite, errors.Join(err, file.Close())
}

func (a *genAction) generate(stms stm.Suite) error {
	cfg, err := config.Load(a.configFile)
	if err != nil {
		return fmt.Errorf("couldn't get config for generator: %w", err)
	}

	g, err := gen.New(*cfg, a.outputDir) // for now
	if err != nil {
		return fmt.Errorf("couldn't create generator: %w", err)
	}

	if err := g.Generate(stms); err != nil {
		return fmt.Errorf("couldn't generate: %w", err)
	}

	return nil
}
