// Package cpp contains C++ configuration.
package cpp

import (
	"github.com/UoY-RoboStar/rtcg/internal/gen/config/catkin"
	"github.com/UoY-RoboStar/rtcg/internal/gen/config/makefile"
	"github.com/UoY-RoboStar/rtcg/internal/testlang/channel"
)

// Config contains configuration for a C++ generator.
type Config struct {
	Variant  Variant   `xml:"variant,attr"` // Variant gives the variant of C++ to generate (e.g. ROS).
	Includes []Include `xml:"include"`      // Includes contains custom includes.
	Channels []Channel `xml:"channel"`      // Channels contains channel configuration.

	Makefile *makefile.Makefile `xml:"makefile"` // Makefile, if given, configures a Makefile.
	Catkin   *catkin.Config     `xml:"catkin"`   // Catkin contains Catkin generator configurations.
}

// New constructs a Config programmatically with the given variant and options.
func New(variant Variant, options ...Option) *Config {
	var cfg Config

	cfg.Variant = variant

	pcfg := &cfg

	for _, option := range options {
		option(pcfg)
	}

	return pcfg
}

// Option is a functional option for building a Config.
type Option func(*Config)

// WithChannel binds channel name to type ty in the configuration.
func WithChannel(name, ty string) Option {
	return func(config *Config) {
		config.Channels = append(config.Channels, Channel{Name: name, Type: ty})
	}
}

func (c *Config) getChannels() []channel.Channel {

	//carray := make([]channel.Channel, len(c.Channels))
	var carray []channel.Channel

	for _, chi := range c.Channels {
		ach := new(channel.Channel)
		err := ach.UnmarshalText([]byte(chi.Name))
		if (err == nil) {
			carray = append(carray, *ach)
		}
	}

	return carray
}

// ChannelMap gets the Channels field of this Config as a map from channel names to type overrides.
func (c *Config) ChannelMap() map[string]string {
	cmap := make(map[string]string, len(c.Channels))

	for _, over := range c.Channels {

		// Create object and parse
		chp := new(channel.Channel)
		err := chp.UnmarshalText([]byte(over.Name))

		if err == nil {
			cmap[chp.Name] = over.Type
		}
	}

	return cmap
}

// ChannelTopicMap gets the Channels field of this Config as a map from channel names to topic overrides.
func (c *Config) ChannelTopicMap() map[string]string {
	cmap := make(map[string]string, len(c.Channels))

	for _, over := range c.Channels {

		// Create object and parse
		chp := new(channel.Channel)
		err := chp.UnmarshalText([]byte(over.Name))
		
		if err == nil {
			cmap[chp.Name] = over.Topic
		} else {
			cmap[over.Name] = over.Topic
		}
	}

	return cmap
}

// ChannelIO gets the Channel types


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
	Topic string `xml:"topic,attr"` // Topic is the ROS topic related to that channel.
}
