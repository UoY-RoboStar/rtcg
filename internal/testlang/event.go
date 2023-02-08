package testlang

import (
	"bytes"
	"fmt"
	"strings"
)

// Event is the type of events in a trace.
//
// An event with an empty channel name is considered to be absent, which is only well-formed for fail nodes.
type Event struct {
	Channel Channel `json:"channel,omitempty"` // Channel is the channel on which the event is occurring.
	Value   Value   `json:"value,omitempty"`   // Value is the value, if any, carried by this event.
}

// NewEvent is shorthand for constructing an Event with channel ch, direction d, and value v.
func NewEvent(ch string, d InOut, v Value) Event {
	return Event{Channel: Channel{Name: ch, Direction: d}, Value: v}
}

// Input is shorthand for constructing an Event with direction In, channel ch, and value v.
func Input(ch string, v Value) Event {
	return NewEvent(ch, In, v)
}

// Output is shorthand for constructing an Event with direction In, channel ch, and value v.
func Output(ch string, v Value) Event {
	return NewEvent(ch, Out, v)
}

func (e *Event) MarshalText() (text []byte, err error) {
	channel, err := e.Channel.MarshalText()
	if err != nil {
		return nil, err
	}

	fields := [][]byte{channel}

	if e.Value.IsPresent() {
		value, err := e.Value.MarshalText()
		if err != nil {
			return nil, fmt.Errorf("couldn't marshal value of event: %w", err)
		}
		fields = append(fields, value)
	}

	return bytes.Join(fields, EventSep), nil
}

func (e *Event) String() string {
	ch := e.Channel.String()
	if e.Value.IsEmpty() {
		return ch
	}
	return strings.Join([]string{ch, e.Value.String()}, ".")
}

func (e *Event) UnmarshalText(text []byte) error {
	val, err := e.Channel.unmarshalTextWithRemainder(text)
	if err != nil {
		return fmt.Errorf("couldn't unmarshal channel of event: %w", err)
	}

	if val == nil {
		// No value in this event.
		return nil
	}

	if err := e.Value.UnmarshalText(val); err != nil {
		return fmt.Errorf("couldn't unmarshal value of event: %w", err)
	}
	return nil
}

// EventSep is the separator used for event fields.
var EventSep = []byte(".")

// BadEventFieldCountError is an error type arising when the number of '.' delimited fields in an Event is not valid.
type BadEventFieldCountError struct {
	Got int // Got is the number of fields we got.
}

func (b BadEventFieldCountError) Error() string {
	return fmt.Sprintf("an event must have three %q delimited fields, got %d", string(EventSep), b.Got)
}
