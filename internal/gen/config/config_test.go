package config_test

import (
	"encoding/xml"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/UoY-RoboStar/rtcg/internal/gen/config"
	"github.com/UoY-RoboStar/rtcg/internal/gen/config/catkin"
	"github.com/UoY-RoboStar/rtcg/internal/gen/config/cpp"
	"github.com/UoY-RoboStar/rtcg/internal/gen/config/makefile"
)

// TestLoadConfig tests Load on sample configs.
func TestLoadConfig(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		want    config.Config
		wantErr error
	}{
		{
			name: "valid1",
			want: config.Config{
				XMLName: xml.Name{Space: "", Local: "rtcg-gen"},
				Cpps: []cpp.Config{
					{
						Variant:  cpp.VariantAnimate,
						Makefile: &makefile.Makefile{},
						Catkin:   nil,
					},
					{
						Variant: cpp.VariantRos,
						Includes: []cpp.Include{
							{Src: "std_msgs/Float32.h", IsSystem: false},
							{Src: "sensor_msgs/BatteryState.h", IsSystem: false},
						},
						Makefile: nil,
						Catkin: &catkin.Config{Package: &catkin.Package{
							XMLName:          xml.Name{Space: "", Local: "package"},
							Format:           2,
							Name:             "bmon-${NAME}",
							Version:          "",
							Description:      "",
							Maintainer:       catkin.Maintainer{Email: "", Name: ""},
							License:          "",
							BuildtoolDepends: nil,
							Depends:          nil,
						}},
					},
				},
				Directory: filepath.Clean("testdata"),
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		test := tt

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got, err := config.Load(filepath.Join("testdata", test.name+".xml"))
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
