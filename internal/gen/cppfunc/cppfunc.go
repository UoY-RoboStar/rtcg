// Package cppfunc contains functions used in C++ code generation.
package cppfunc

import (
	"fmt"
	"strings"
	"text/template"

	"github.com/UoY-RoboStar/rtcg/internal/testlang"
)

// Funcs adds the C++ function map to base.
func Funcs(base *template.Template) *template.Template {
	return base.Funcs(template.FuncMap{
		"cppEnumField":   EnumField,
		"cppOutcomeEnum": OutcomeEnum,
		"cppStateEntry":  StateEntry,
		"cppStateEnum":   StateEnum,
		"cppTestEnum":    TestEnum,
	})
}

// StateEntry gets the name of the entry method for the state with the given id.
func StateEntry(id testlang.NodeID) string {
	return string(id) + "Entry"
}

// OutcomeEnum gets a reference to the enum member for outcome.
func OutcomeEnum(outcome testlang.Outcome) string {
	return enumName(outcomeEnumName, outcome)
}

// StateEnum gets a reference to the enum member for the state with the given id.
func StateEnum(id testlang.NodeID) string {
	return enumName(stateEnumName, id)
}

// TestEnum gets a reference to the enum member for the test called name.
func TestEnum(name string) string {
	return enumName(testEnumName, name)
}

// EnumField massages variant to become suitable as an enum field name.
func EnumField(variant any) string {
	return strings.ToUpper(fmt.Sprint(variant))
}

func enumName(name string, variant any) string {
	return name + "::" + EnumField(variant)
}

const (
	stateEnumName   = "State"   // stateEnumName is the name in the C++ code for the state enum.
	testEnumName    = "Test"    // testEnumName is the name in the C++ code for the test enum.
	outcomeEnumName = "Outcome" // outcomeEnumName is the name in the C++ code for the outcome enum.
)
