package transition

import (
	"fmt"
	"strings"

	"github.com/UoY-RoboStar/rtcg/internal/testlang"
	"github.com/UoY-RoboStar/rtcg/internal/testlang/channel"
)

// BaseSet is the structure common to all transition sets.
type BaseSet struct {
	Channel channel.Channel `json:"channel"` // Channel is the channel at the head of this set.
}

// IsForChannel checks whether b is for channel.
func (b BaseSet) IsForChannel(channel channel.Channel) bool {
	return b.Channel.Equals(channel)
}

// Set is a set of transitions for a given channel.
type Set struct {
	BaseSet
	Transitions []Transition `json:"transitions,omitempty"` // Transitions is the list of transitions.
}

// NewSet constructs a transition set with the given channel and transitions.
func NewSet(channel channel.Channel, transitions ...Transition) Set {
	return Set{BaseSet: BaseSet{Channel: channel}, Transitions: transitions}
}

// AddToSets adds the transition under channel into sets.
//
// If there is an existing set in sets for channel, it acquires the new transition in-place, and we return sets.
// Otherwise, we return the result of appending a new set with that channel and transition onto sets.
func AddToSets(sets []Set, channel channel.Channel, transition Transition) []Set {
	// Try merging onto an existing set.
	for i := range sets {
		if ts := &sets[i]; ts.IsForChannel(channel) {
			ts.Transitions = append(ts.Transitions, transition)

			return sets
		}
	}
	// No transition set with this channel exists yet.
	return append(sets, NewSet(channel, transition))
}

func (t Set) String() string {
	tstrs := make([]string, len(t.Transitions))
	for i, v := range t.Transitions {
		tstrs[i] = v.String()
	}

	return fmt.Sprintf("%s:{%s}", &t.Channel, strings.Join(tstrs, ", "))
}

// StateMap is a map from state IDs to transition lists.
type StateMap map[testlang.NodeID][]Transition

// AggregateSet is an aggregate of Set instances on states.
type AggregateSet struct {
	BaseSet

	States StateMap `json:"states,omitempty"` // States maps states to transition lists.
}

// NewAggregateSet constructs an aggregate set with the given channel and state-transition map.
func NewAggregateSet(channel channel.Channel, states map[testlang.NodeID][]Transition) AggregateSet {
	return AggregateSet{BaseSet: BaseSet{Channel: channel}, States: states}
}

// AddToAggregateSets adds a transition-set set into aggs under the state with ID stateID.
//
// If there is an existing map in aggs for channel, it acquires the transitions of set in-place, and we return aggs.
// Otherwise, we return the result of appending a new aggregate with that transition set onto aggs.
func AddToAggregateSets(aggs []AggregateSet, stateID testlang.NodeID, set Set) []AggregateSet {
	// Try merging onto an existing set.
	for i := range aggs {
		if ts := &aggs[i]; ts.IsForChannel(set.Channel) {
			ts.States[stateID] = append(ts.States[stateID], set.Transitions...)

			return aggs
		}
	}
	// No transition set with this channel exists yet.
	return append(aggs, NewAggregateSet(set.Channel, map[testlang.NodeID][]Transition{stateID: set.Transitions}))
}
