package catkin_test

import (
	"encoding/xml"
	"reflect"
	"testing"

	"github.com/UoY-RoboStar/rtcg/internal/gen/config/catkin"
)

func TestPackage_Autofill(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		input catkin.Package
		want  catkin.Package
	}{
		"empty": {
			input: catkin.Package{
				XMLName:          xml.Name{Space: "", Local: ""},
				Format:           0,
				Name:             "",
				Version:          "",
				Description:      "",
				Maintainer:       catkin.Maintainer{Email: "", Name: ""},
				License:          "",
				BuildtoolDepends: nil,
				Depends:          nil,
			},
			want: catkin.Package{
				XMLName:          xml.Name{Space: "", Local: ""},
				Format:           2,
				Name:             "${NAME}",
				Version:          "0.1.0",
				Description:      "Package for ${NAME}",
				Maintainer:       catkin.Maintainer{Email: "email@example.com", Name: "${NAME} Maintainers"},
				License:          "Proprietary",
				BuildtoolDepends: nil,
				Depends:          nil,
			},
		},
	}

	for name, tt := range tests {
		name := name
		tt := tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := tt.input
			got.Autofill()

			cmp(t, got, tt.want)
		})
	}
}

func TestPackage_Expand(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		input catkin.Package
		want  catkin.Package
	}{
		"empty": {
			input: catkin.Package{
				XMLName:          xml.Name{Space: "", Local: ""},
				Format:           0,
				Name:             "",
				Version:          "",
				Description:      "",
				Maintainer:       catkin.Maintainer{Email: "", Name: ""},
				License:          "",
				BuildtoolDepends: nil,
				Depends:          nil,
			},
			want: catkin.Package{
				XMLName:          xml.Name{Space: "", Local: ""},
				Format:           0,
				Name:             "",
				Version:          "",
				Description:      "",
				Maintainer:       catkin.Maintainer{Email: "", Name: ""},
				License:          "",
				BuildtoolDepends: nil,
				Depends:          nil,
			},
		},
		"minimal": {
			input: catkin.Package{
				XMLName:          xml.Name{Space: "", Local: ""},
				Format:           2,
				Name:             "${NAME}",
				Version:          "0.1.0",
				Description:      "Package for ${NAME}",
				Maintainer:       catkin.Maintainer{Email: "email@example.com", Name: "${NAME} Maintainers"},
				License:          "Proprietary",
				BuildtoolDepends: nil,
				Depends:          nil,
			},
			want: catkin.Package{
				XMLName:          xml.Name{Space: "", Local: ""},
				Format:           2,
				Name:             "foobar",
				Version:          "0.1.0",
				Description:      "Package for foobar",
				Maintainer:       catkin.Maintainer{Email: "email@example.com", Name: "foobar Maintainers"},
				License:          "Proprietary",
				BuildtoolDepends: nil,
				Depends:          nil,
			},
		},
	}

	for name, tt := range tests {
		name := name
		tt := tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := tt.input
			got.Expand("foobar")

			cmp(t, got, tt.want)
		})
	}
}

// TODO: Write

func cmp(t *testing.T, got, want catkin.Package) {
	t.Helper()

	for _, field := range []Field{
		FFormat{},
		FFormat{},
		FName{},
		FVersion{},
		FDescription{},
		FMaintainerEmail{},
		FMaintainerName{},
		FLicense{},
		FBuildtoolDepends{},
		FDepends{},
	} {
		gotF := field.Select(&got)
		wantF := field.Select(&want)

		if !reflect.DeepEqual(gotF, wantF) {
			t.Errorf("got %s %v, want %v", field.Name(), gotF, wantF)
		}
	}
}

type Field interface {
	Name() string
	Select(p *catkin.Package) any
}

type FFormat struct{}

func (FFormat) Name() string {
	return "format"
}

func (FFormat) Select(p *catkin.Package) any {
	return p.Format
}

type FName struct{}

func (FName) Name() string {
	return "name"
}

func (FName) Select(p *catkin.Package) any {
	return p.Name
}

type FVersion struct{}

func (FVersion) Name() string {
	return "version"
}

func (FVersion) Select(p *catkin.Package) any {
	return p.Version
}

type FDescription struct{}

func (FDescription) Name() string {
	return "description"
}

func (FDescription) Select(p *catkin.Package) any {
	return p.Description
}

type FMaintainerEmail struct{}

func (FMaintainerEmail) Name() string {
	return "maintainer email"
}

func (FMaintainerEmail) Select(p *catkin.Package) any {
	return p.Maintainer.Email
}

type FMaintainerName struct{}

func (FMaintainerName) Name() string {
	return "maintainer name"
}

func (FMaintainerName) Select(p *catkin.Package) any {
	return p.Maintainer.Name
}

type FLicense struct{}

func (FLicense) Name() string {
	return "license"
}

func (FLicense) Select(p *catkin.Package) any {
	return p.License
}

type FBuildtoolDepends struct{}

func (FBuildtoolDepends) Name() string {
	return "buildtool depends"
}

func (FBuildtoolDepends) Select(p *catkin.Package) any {
	return p.BuildtoolDepends
}

type FDepends struct{}

func (FDepends) Name() string {
	return "depends"
}

func (FDepends) Select(p *catkin.Package) any {
	return p.Depends
}
