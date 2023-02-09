package testlang

import (
	"bytes"
	"fmt"
	"strings"
)

// Channel is the type of directional channels.
type Channel struct {
	Name      string    `json:"name"`      // Name is the name of the channel.
	Direction Direction `json:"direction"` // Direction is the direction of the channel.
}

// IsEmpty gets whether the channel is considered empty.
//
// Emptiness of a channel is solely dependent on its name.
func (c *Channel) IsEmpty() bool {
	return c.Name == ""
}

// Equals gets whether this channel equals another channel.
func (c *Channel) Equals(other Channel) bool {
	return c.Name == other.Name && c.Direction == other.Direction
}

func (c *Channel) MarshalText() ([]byte, error) {
	name := []byte(c.Name)
	// The only valid Channel with an empty name is an empty one, which marshals to an empty string.
	if len(name) == 0 {
		return []byte{}, nil
	}

	direction, err := c.Direction.MarshalText()
	if err != nil {
		return nil, fmt.Errorf("couldn't marshal direction of channel: %w", err)
	}

	return eventSepJoin(name, direction), nil
}

func (c *Channel) UnmarshalText(text []byte) error {
	_, err := c.unmarshalTextWithRemainder(text)

	return err
}

func (c *Channel) String() string {
	if c.IsEmpty() {
		return "(empty channel)"
	}

	return strings.Join([]string{c.Name, c.Direction.String()}, eventSep)
}

const (
	// eventSep is the separator used for events and channels.
	eventSep = "."
	// numChannelParts is the number of parts a channel can theoretically be split into.
	numChannelParts = 2
	// numEventParts is the number of parts an event can theoretically be split into.
	numEventParts = numChannelParts + 1
)

func eventSepJoin(items ...[]byte) []byte {
	return bytes.Join(items, []byte(eventSep))
}

func (c *Channel) unmarshalTextWithRemainder(text []byte) ([]byte, error) {
	fields := bytes.SplitN(text, []byte(eventSep), numEventParts)

	numFields := len(fields)
	if numFields < numChannelParts || numEventParts < numFields {
		return nil, BadEventFieldCountError{Got: numFields}
	}

	c.Name = string(bytes.TrimSpace(fields[0]))

	if err := c.Direction.UnmarshalText(fields[1]); err != nil {
		return nil, fmt.Errorf("couldn't unmarshal direction of event: %w", err)
	}

	if numFields == numChannelParts {
		return nil, nil
	}

	return fields[numEventParts-1], nil
}
