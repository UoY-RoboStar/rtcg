package gen

import (
	"encoding/xml"
	"errors"
	"fmt"
	"os"
)

// Config contains configuration for the generator.
type Config struct {
	XMLName   xml.Name         `xml:"rtcg-gen"` // XMLName sets the name of the Config struct in XML.
	Cpps      []CppTarget      `xml:"cpp"`      // Cpps contains CppTarget elements.
	Makefiles []MakefileTarget `xml:"makefile"` // Makefiles contains MakefileTarget elements.
}

// LoadConfig loads a generator config at path.
func LoadConfig(path string) (*Config, error) {
	var config Config

	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("couldn't open config file %q: %w", path, err)
	}

	if err := xml.NewDecoder(file).Decode(&config); err != nil {
		err = fmt.Errorf("couldn't decode config file %q: %w", path, err)

		return nil, errors.Join(err, file.Close())
	}

	if err := file.Close(); err != nil {
		return nil, fmt.Errorf("couldn't close config file %q: %w", path, err)
	}

	return &config, nil
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
