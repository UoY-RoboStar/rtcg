package testlang

import (
	"bytes"
	"fmt"
)

// Event is the type of events in a trace.
//
// An event with an empty channel name is considered to be absent, which is only well-formed for fail nodes.
type Event struct {
	Channel   string `json:"channel,omitempty"`   // Channel is the name of the channel on which the event is occurring.
	Direction InOut  `json:"direction,omitempty"` // InOut is the direction of the event.
	Value     Value  `json:"value,omitempty"`     // Value is the value, if any, carried by this event.
}

func (e *Event) MarshalText() (text []byte, err error) {
	channel := []byte(e.Channel)

	direction, err := e.Direction.MarshalText()
	if err != nil {
		return nil, fmt.Errorf("couldn't marshal direction of event: %w", err)
	}

	value, err := e.Value.MarshalText()
	if err != nil {
		return nil, fmt.Errorf("couldn't marshal value of event: %w", err)
	}

	fields := [][]byte{channel, direction, value}
	return bytes.Join(fields, EventSep), nil
}

func (e *Event) UnmarshalText(text []byte) error {
	fields := bytes.Split(text, EventSep)
	numFields := len(fields)
	if numFields != 3 {
		return BadEventFieldCountError{Got: numFields}
	}

	e.Channel = string(fields[0])
	if err := e.Direction.UnmarshalText(fields[1]); err != nil {
		return fmt.Errorf("couldn't unmarshal direction of event: %w", err)
	}
	if err := e.Value.UnmarshalText(fields[2]); err != nil {
		return fmt.Errorf("couldn't unmarshal value of event: %w", err)
	}
	return nil
}

var EventSep = []byte(".")

// BadEventFieldCountError is an error type arising when the number of '.' delimited fields in an Event is not valid.
type BadEventFieldCountError struct {
	Got int // Got is the number of fields we got.
}

func (b BadEventFieldCountError) Error() string {
	return fmt.Sprintf("an event must have three %q delimited fields, got %d", string(EventSep), b.Got)
}
