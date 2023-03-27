package cpp

import (
	"github.com/UoY-RoboStar/rtcg/internal/stm"
	"github.com/UoY-RoboStar/rtcg/internal/testlang/rstype"
)

// Context contains any context derived from the C++ generator config.
//
// This is common to multiple different generators that consider C++ config.
type Context struct {
	Includes      []Include              // Includes contains the user-configured includes.
	ChannelTypes  map[string]ChannelType // ChannelTypes contains the calculated channel types.
	HasConversion bool                   // HasConversion is true if there is a convert.cpp file.
}

// Process processes a config into a Context.
// It expects the unified type-map from the testing suite.
func (c *Config) Process(types stm.TypeMap) Context {
	ctx := Context{
		Includes:      c.Includes,
		ChannelTypes:  make(map[string]ChannelType, len(types)),
		HasConversion: false,
	}

	overrides := c.ChannelMap()

	for cname, ctype := range types {
		override, ok := overrides[cname]

		if ok {
			ctx.HasConversion = true
		}

		ctx.ChannelTypes[cname] = ChannelType{Base: ctype, Override: override}
	}

	return ctx
}

// ChannelType defines a type override for a channel.
type ChannelType struct {
	Base     *rstype.RsType // Base is the inferred RoboStar type for this channel.
	Override string         // Override is any manually specified type for this channel.
}

// HasOverride is true if the channel type has been overridden.
func (t ChannelType) HasOverride() bool {
	return t.Override != ""
}
