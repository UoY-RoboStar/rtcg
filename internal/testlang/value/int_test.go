package value_test

import (
	"reflect"
	"testing"

	"github.com/UoY-RoboStar/rtcg/internal/testlang/rstype"
	"github.com/UoY-RoboStar/rtcg/internal/testlang/value"
)

// TestIntValue_Type tests the Type method of IntValue.
func TestIntValue_Type(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		i    value.IntValue
		want *rstype.RsType
	}{
		{name: "positive", i: value.IntValue(4), want: rstype.Arithmos(rstype.ArithmosDomainNat)},
		{name: "zero", i: value.IntValue(0), want: rstype.Arithmos(rstype.ArithmosDomainNat)},
		{name: "negative", i: value.IntValue(-4), want: rstype.Arithmos(rstype.ArithmosDomainInt)},
	}
	for _, tt := range tests {
		test := tt

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := test.i.Type(); !reflect.DeepEqual(got, test.want) {
				t.Errorf("Type() = %v, want %v", got, test.want)
			}
		})
	}
}
