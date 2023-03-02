package value

import (
	"github.com/UoY-RoboStar/rtcg/internal/testlang/rstype"
	"reflect"
	"testing"
)

func TestIntValue_Type(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		i    IntValue
		want *rstype.RsType
	}{
		{name: "positive", i: IntValue(4), want: rstype.Arithmos(rstype.ArithmosDomainNat)},
		{name: "zero", i: IntValue(0), want: rstype.Arithmos(rstype.ArithmosDomainNat)},
		{name: "negative", i: IntValue(-4), want: rstype.Arithmos(rstype.ArithmosDomainInt)},
	}
	for _, tt := range tests {
		test := tt

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := tt.i.Type(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Type() = %v, want %v", got, tt.want)
			}
		})
	}
}
