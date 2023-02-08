package testlang

import (
	"bytes"
	"fmt"
	"strings"
)

// Channel is the type of directional channels.
type Channel struct {
	Name      string `json:"name"`      // Name is the name of the channel.
	Direction InOut  `json:"direction"` // Direction is the direction of the channel.
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

func (c *Channel) MarshalText() (text []byte, err error) {
	name := []byte(c.Name)
	// The only valid Channel with an empty name is an empty one, which marshals to an empty string.
	if len(name) == 0 {
		return []byte{}, nil
	}

	direction, err := c.Direction.MarshalText()
	if err != nil {
		return nil, fmt.Errorf("couldn't marshal direction of channel: %w", err)
	}

	return bytes.Join([][]byte{name, direction}, EventSep), nil
}

func (c *Channel) UnmarshalText(text []byte) error {
	_, err := c.unmarshalTextWithRemainder(text)
	return err
}

func (c *Channel) String() string {
	if c.IsEmpty() {
		return "(empty channel)"
	}
	return strings.Join([]string{c.Name, c.Direction.String()}, ".")
}

func (c *Channel) unmarshalTextWithRemainder(text []byte) ([]byte, error) {
	fields := bytes.SplitN(text, EventSep, 3)
	numFields := len(fields)
	if numFields < 2 || 3 < numFields {
		return nil, BadEventFieldCountError{Got: numFields}
	}

	c.Name = string(bytes.TrimSpace(fields[0]))

	if err := c.Direction.UnmarshalText(fields[1]); err != nil {
		return nil, fmt.Errorf("couldn't unmarshal direction of event: %w", err)
	}

	if numFields == 2 {
		return nil, nil
	}
	return fields[2], nil
}
