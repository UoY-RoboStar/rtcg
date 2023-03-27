// Package subgen contains infrastructure for sub-generators.
package subgen

import "github.com/UoY-RoboStar/rtcg/internal/stm"

// Subgenerator captures the idea of a test code sub-generator.
type Subgenerator interface {
	// Name gets the name of this item.
	Name() string

	// OnSuite specialises this generator for a particular suite.
	// This allows for pre-computing of contexts and other data that is needed for a particular suite.
	OnSuite(suite *stm.Suite) OnSuite
}

// OnSuite is a sub-generator configured against a particular test suite.
type OnSuite interface {
	// Parent gets the parent sub-generator.
	Parent() Subgenerator

	// Dirs gets the list of directories to make for the given test suite.
	Dirs() []string

	// Generate generates code for suite.
	Generate() error
}
