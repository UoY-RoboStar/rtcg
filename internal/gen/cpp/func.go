package cpp

import (
	"fmt"
	"strings"
	"text/template"

	"github.com/UoY-RoboStar/rtcg/internal/strmanip"
	"github.com/UoY-RoboStar/rtcg/internal/testlang"
	"github.com/UoY-RoboStar/rtcg/internal/testlang/rstype"
	"github.com/UoY-RoboStar/rtcg/internal/testlang/value"
)

// Funcs gets the C++ function map.
func Funcs() template.FuncMap {
	return template.FuncMap{
		"cppCallbackName":     CallbackName,
		"cppChannelMsgType":   ChannelMsgType,
		"cppChannelValueType": ChannelValueType,
		"cppConvertFrom":      ConvertFrom,
		"cppConvertTo":        ConvertTo,
		"cppEnumField":        EnumField,
		"cppOutcomeEnum":      OutcomeEnum,
		"cppStateEntry":       StateEntry,
		"cppStateEnum":        StateEnum,
		"cppTestEnum":         TestEnum,
		"cppType":             StdType,
		"cppValue":            Value,
	}
}

// CallbackName gets the name of the callback for the channel with name cha.
func CallbackName(cha string) string {
	return cha + "Callback"
}

// ChannelMsgType gets the name of the defined message type for the channel with name cha.
// This is usually a pointer to the value type.
func ChannelMsgType(cha string) string {
	return strmanip.UpcaseFirst(cha) + "Msg"
}

// ChannelValueType gets the name of the defined message type for the channel with name cha.
// This is usually a pointer to the value type.
func ChannelValueType(cha string) string {
	return strmanip.UpcaseFirst(cha) + "Val"
}

// ConvertTo gets the to-conversion function name for the channel with name cha.
func ConvertTo(cha string) string {
	return convert("to", cha)
}

// ConvertFrom gets the from-conversion function name for channel with name cha.
func ConvertFrom(cha string) string {
	return convert("from", cha)
}

func convert(dir, cha string) string {
	return dir + strmanip.UpcaseFirst(cha)
}

// StateEntry gets the name of the entry method for the state with the given id.
func StateEntry(id testlang.NodeID) string {
	return "enter" + strmanip.UpcaseFirst(string(id))
}

// OutcomeEnum gets a reference to the enum member for outcome.
func OutcomeEnum(outcome testlang.Outcome) string {
	return rtcgEnumName(outcomeEnumName, outcome)
}

// StateEnum gets a reference to the enum member for the state with the given id.
func StateEnum(id testlang.NodeID) string {
	return localEnumName(stateEnumName, id)
}

// TestEnum gets a reference to the enum member for the test called name.
func TestEnum(name string) string {
	return localEnumName(testEnumName, name)
}

// EnumField massages variant to become suitable as an enum field name.
func EnumField(variant any) string {
	return strings.ToUpper(fmt.Sprint(variant))
}

// StdType gets the standard C++ type for the given RoboStar type.
func StdType(rsType rstype.RsType) string {
	switch {
	case rsType.IsEnum():
		return "std::string"
	case rsType.IsNat():
		return "unsigned int"
	case rsType.IsInt():
		return "int"
	case rsType.IsReal():
		return "double"
	default:
		return "void *"
	}
}

// Value gets a C++ encoding of a value.
func Value(val value.Value) string {
	// TODO: do this without type introspection.
	typ := val.Type()

	if typ.IsEnum() {
		return fmt.Sprintf("\"%s\"", val.StringValue())
	}

	return val.StringValue()
}

func rtcgEnumName(name string, variant any) string {
	return "rtcg::" + localEnumName(name, variant)
}

func localEnumName(name string, variant any) string {
	return name + "::" + EnumField(variant)
}

const (
	stateEnumName   = "State"   // stateEnumName is the name in the C++ code for the state enum.
	testEnumName    = "Test"    // testEnumName is the name in the C++ code for the test enum.
	outcomeEnumName = "Outcome" // outcomeEnumName is the name in the C++ code for the outcome enum.
)
