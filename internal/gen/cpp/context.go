package cpp

import (
	"github.com/UoY-RoboStar/rtcg/internal/gen/gencommon"
	"github.com/UoY-RoboStar/rtcg/internal/stm"
	"github.com/UoY-RoboStar/rtcg/internal/testlang/rstype"
)

// Context is the context passed into the C++ code generator.
type Context struct {
	gencommon.Context

	Includes     []Include              // Includes contains the user-configured includes.
	ChannelTypes map[string]ChannelType // ChannelTypes contains the calculated channel types.
}

// NewContext creates a new template context from a named state machine.
func NewContext(name string, machine *stm.Stm, config Config) Context {
	return Context{
		Context:      gencommon.NewContext(name, machine),
		Includes:     config.Includes,
		ChannelTypes: channelTypes(config, machine),
	}
}

func channelTypes(config Config, machine *stm.Stm) map[string]ChannelType {
	overrides := make(map[string]ChannelType, len(config.Channels))

	for _, ccfg := range config.Channels {
		if ccfg.Type != "" {
			overrides[ccfg.Name] = ChannelType{Base: machine.Types[ccfg.Name], Override: ccfg.Type}
		}
	}

	return overrides
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
