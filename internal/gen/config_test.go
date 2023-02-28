package gen_test

import (
	"encoding/xml"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/UoY-RoboStar/rtcg/internal/gen"
)

// TestLoadConfig tests LoadConfig on sample configs.
func TestLoadConfig(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		want    gen.Config
		wantErr error
	}{
		{
			name: "valid1",
			want: gen.Config{
				XMLName: xml.Name{
					Space: "",
					Local: "rtcg-gen",
				},
				Cpps: []gen.CppTarget{
					{
						Variant: "animate",
					},
					{
						Variant: "ros",
						Includes: []gen.Include{
							{Src: "std_msgs/Float32.h", IsSystem: false},
							{Src: "sensor_msgs/BatteryState.h", IsSystem: false},
						},
					},
				},
				Makefiles: []gen.MakefileTarget{{}},
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		test := tt

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got, err := gen.LoadConfig(filepath.Join("testdata", test.name+".xml"))
			if !reflect.DeepEqual(err, test.wantErr) {
				t.Errorf("got error %v, wanted %v", err, test.wantErr)

				return
			}
			if !reflect.DeepEqual(*got, test.want) {
				t.Errorf("got config %v, wanted %v", got, test.want)
			}
		})
	}
}
