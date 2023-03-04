package testlang

import (
	"fmt"
	"strings"

	"github.com/UoY-RoboStar/rtcg/internal/testlang/channel"
	"github.com/UoY-RoboStar/rtcg/internal/testlang/value"
)

// Event is the type of events in a trace.
//
// An event with an empty channel name is considered to be absent, which is only well-formed for fail nodes.
type Event struct {
	Channel channel.Channel `json:"channel,omitempty"` // Channel is the channel on which the event is occurring.
	Value   value.Value     `json:"value,omitempty"`   // Value is the value, if any, carried by this event.
}

// NewEvent is shorthand for constructing an Event with channel ch, direction d, and value v.
func NewEvent(ch string, d channel.Kind, v value.Value) Event {
	return Event{Channel: channel.Channel{Name: ch, Kind: d}, Value: v}
}

// Input is shorthand for constructing an Event with direction TypeIn, channel ch, and value v.
func Input(ch string, v value.Value) Event {
	return NewEvent(ch, channel.KindIn, v)
}

// Output is shorthand for constructing an Event with direction TypeIn, channel ch, and value v.
func Output(ch string, v value.Value) Event {
	return NewEvent(ch, channel.KindOut, v)
}

// Equals checks whether this event equals other.
func (e *Event) Equals(other Event) bool {
	return e.Channel.Equals(other.Channel) && e.Value.Equals(other.Value)
}

func (e *Event) MarshalText() ([]byte, error) {
	chanBytes, err := e.Channel.MarshalText()
	if err != nil {
		return nil, fmt.Errorf("couldn't marshal channel of event: %w", err)
	}

	if e.Value.IsEmpty() {
		return chanBytes, nil
	}

	val, err := e.Value.MarshalText()
	if err != nil {
		return nil, fmt.Errorf("couldn't marshal value of event: %w", err)
	}

	return channel.EventSepJoin(chanBytes, val), nil
}

func (e *Event) String() string {
	ch := e.Channel.String()
	if e.Value.IsEmpty() {
		return ch
	}

	return strings.Join([]string{ch, e.Value.String()}, channel.EventSep)
}

func (e *Event) UnmarshalText(text []byte) error {
	val, err := e.Channel.UnmarshalTextWithRemainder(text)
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
