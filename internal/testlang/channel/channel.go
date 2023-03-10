// Package channel describes communication channels in the testing language.
//
// These are effectively the CSP encoding of the RoboChart communications model.
package channel

import (
	"bytes"
	"fmt"
	"strings"
)

// Channel is the type of directional channels.
type Channel struct {
	Name string `json:"name"` // Name is the name of the channel.
	Kind Kind   `json:"kind"` // Kind is the direction of the channel.
}

// New is shorthand for constructing a Channel with the given name and kind.
func New(name string, kind Kind) Channel {
	return Channel{Name: name, Kind: kind}
}

// In constructs an input Channel.
func In(name string) Channel {
	return New(name, KindIn)
}

// IsIn gets whether this Channel is an input.
func (c *Channel) IsIn() bool {
	return c.Kind == KindIn
}

// Out constructs an output Channel.
func Out(name string) Channel {
	return New(name, KindOut)
}

// IsOut gets whether this Channel is an output.
//
// Note that this predicate does not count channels with KindCall as outputs.
func (c *Channel) IsOut() bool {
	return c.Kind == KindOut
}

// Call constructs an operation call Channel.
func Call(name string) Channel {
	return New(name, KindCall)
}

// IsCall gets whether this Channel is a call.
func (c *Channel) IsCall() bool {
	return c.Kind == KindCall
}

// IsEmpty gets whether the channel is considered empty.
//
// Emptiness of a channel is solely dependent on its name.
func (c *Channel) IsEmpty() bool {
	return c.Name == ""
}

// Equals gets whether this channel equals another channel.
func (c *Channel) Equals(other Channel) bool {
	return c.Name == other.Name && c.Kind == other.Kind
}

func (c *Channel) MarshalText() ([]byte, error) {
	name := []byte(c.Name)
	// The only valid Channel with an empty name is an empty one, which marshals to an empty string.
	if len(name) == 0 {
		return []byte{}, nil
	}

	kind, err := c.Kind.MarshalText()
	if err != nil {
		return nil, fmt.Errorf("couldn't marshal kind of channel: %w", err)
	}

	return EventSepJoin(name, kind), nil
}

func (c *Channel) UnmarshalText(text []byte) error {
	_, err := c.UnmarshalTextWithRemainder(text)

	return err
}

func (c *Channel) String() string {
	if c.IsEmpty() {
		return "(empty channel)"
	}

	return strings.Join([]string{c.Name, c.Kind.String()}, EventSep)
}

const (
	// EventSep is the separator used for events and channels.
	EventSep = "."
	// numChannelParts is the number of parts a channel can theoretically be split into.
	numChannelParts = 2
	// numEventParts is the number of parts an event can theoretically be split into.
	numEventParts = numChannelParts + 1
)

func EventSepJoin(items ...[]byte) []byte {
	return bytes.Join(items, []byte(EventSep))
}

func (c *Channel) UnmarshalTextWithRemainder(text []byte) ([]byte, error) {
	fields := bytes.SplitN(text, []byte(EventSep), numEventParts)

	numFields := len(fields)
	if numFields < numChannelParts || numEventParts < numFields {
		return nil, BadFieldCountError{Got: numFields}
	}

	c.Name = string(bytes.TrimSpace(fields[0]))

	kindBytes := bytes.TrimSpace(fields[1])
	if err := c.Kind.UnmarshalText(kindBytes); err != nil {
		return nil, fmt.Errorf("couldn't unmarshal direction of event: %w", err)
	}

	if numFields == numChannelParts {
		return nil, nil
	}

	return fields[numEventParts-1], nil
}

// BadFieldCountError is the error when the number of '.' delimited fields in an Event or Channel is not valid.
type BadFieldCountError struct {
	Got int // Got is the number of fields we got.
}

func (b BadFieldCountError) Error() string {
	return fmt.Sprintf("needed two or three %q delimited fields, got %d", EventSep, b.Got)
}
