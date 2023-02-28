package gen

import "encoding/xml"

// Config contains serialisable configuration for the generator.
type Config struct {
	XMLName   xml.Name         `xml:"rtcg-gen"` // XMLName sets the name of the Config struct in XML.
	Cpps      []CppTarget      `xml:"cpp"`      // Cpps contains CppTarget elements.
	Makefiles []MakefileTarget `xml:"makefile"` // Makefiles contains MakefileTarget elements.
}

// CppTarget configures a C++ generator.
type CppTarget struct {
	Variant  string    `xml:"variant,attr"` // Variant gives the variant of C++ to generate (e.g. ROS).
	Includes []Include `xml:"include"`      // Includes contains custom includes.
}

// MakefileTarget configures a Makefile.
type MakefileTarget struct{}

// Include captures a custom include header.
type Include struct {
	Src      string `xml:"src,attr"`    // Src is the source to include.
	IsSystem bool   `xml:"system,attr"` // IsSystem, if true, generates a `<>` include.
}
