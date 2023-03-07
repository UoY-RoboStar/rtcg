package gen

import (
	"encoding/xml"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/UoY-RoboStar/rtcg/internal/gen/cpp"
)

// Config contains configuration for the generator.
type Config struct {
	XMLName   xml.Name     `xml:"rtcg-gen"`            // XMLName sets the name of the Config struct in XML.
	Cpps      []cpp.Config `xml:"cpp"`                 // Cpps contains CppTarget elements.
	Directory string       `xml:"directory,omitempty"` // Directory contains the root directory for generator files.
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

	if config.Directory == "" {
		config.Directory = filepath.Dir(path)
	}

	return &config, nil
}
