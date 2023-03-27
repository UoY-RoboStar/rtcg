// Package catkin contains configuration for the Catkin generator.
package catkin

// Config provides configuration for the Catkin generator.
type Config struct {
	Package *Package `xml:"package,omitempty"` // Package contains package.xml overrides.
}
