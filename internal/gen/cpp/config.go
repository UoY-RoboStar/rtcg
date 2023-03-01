package cpp

import "github.com/UoY-RoboStar/rtcg/internal/gen/makefile"

// Config contains configuration for a C++ generator.
type Config struct {
	Variant  Variant          `xml:"variant,attr"` // Variant gives the variant of C++ to generate (e.g. ROS).
	Includes []Include        `xml:"include"`      // Includes contains custom includes.
	Makefile *makefile.Config `xml:"makefile"`     // Makefile, if given, configures a Makefile.
}

// Include captures a custom include header.
type Include struct {
	Src      string `xml:"src,attr"`    // Src is the source to include.
	IsSystem bool   `xml:"system,attr"` // IsSystem, if true, generates a `<>` include.
}
