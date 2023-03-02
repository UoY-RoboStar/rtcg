package cpp

import "github.com/UoY-RoboStar/rtcg/internal/gen/makefile"

// Config contains configuration for a C++ generator.
type Config struct {
	Variant  Variant          `xml:"variant,attr"` // Variant gives the variant of C++ to generate (e.g. ROS).
	Includes []Include        `xml:"include"`      // Includes contains custom includes.
	Makefile *makefile.Config `xml:"makefile"`     // Makefile, if given, configures a Makefile.
	Channels []Channel        `xml:"channel"`      // Channels contains channel configuration.
}

// Include captures a custom include header.
type Include struct {
	Src      string `xml:"src,attr"`    // Src is the source to include.
	IsSystem bool   `xml:"system,attr"` // IsSystem, if true, generates a `<>` include.
}

// Channel provides configuration for a channel.
//
// If the type of the channel is overridden, the test will generate calls into conversion
// functions to switch from the native C++ encoding of the channel's RoboStar type to and from
// the custom type.
type Channel struct {
	Name string `xml:"name,attr"` // Name is the name of the affected channel.
	Type string `xml:"type,attr"` // Type is the C++ value type of the channel.
}
