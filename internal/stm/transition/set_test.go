package transition_test

import (
	"github.com/UoY-RoboStar/rtcg/internal/testlang/value"
	"reflect"
	"testing"

	"github.com/UoY-RoboStar/rtcg/internal/stm/transition"
	"github.com/UoY-RoboStar/rtcg/internal/testlang"
	"github.com/UoY-RoboStar/rtcg/internal/testlang/channel"
)

// TestAddToAggregateSets performs a test run of transition.AddToAggregateSets.
func TestAddToAggregateSets(t *testing.T) {
	t.Parallel()

	channel1 := channel.In("foo")
	channel2 := channel.Out("bar")

	tr1 := transition.Transition{Value: value.Int(42), Next: "y"}
	tr2 := transition.Transition{Value: value.Int(56), Next: "z"}
	tr3 := transition.Transition{Value: value.Enum("Out"), Next: "z"}
	tr4 := transition.Transition{Value: value.Int(72), Next: ""}

	var aggs []transition.AggregateSet

	aggs = transition.AddToAggregateSets(aggs, "x", transition.NewSet(channel1, tr1))
	aggs = transition.AddToAggregateSets(aggs, "x", transition.NewSet(channel1, tr2))
	aggs = transition.AddToAggregateSets(aggs, "y", transition.NewSet(channel1, tr3))
	aggs = transition.AddToAggregateSets(aggs, "y", transition.NewSet(channel2, tr4))

	if len(aggs) != 2 {
		t.Fatalf("should have created 2 aggregate sets, got %d (%v)", len(aggs), aggs)
	}

	checkAggregateSet(t, aggs[0], channel1, map[testlang.NodeID][]transition.Transition{"x": {tr1, tr2}, "y": {tr3}})
	checkAggregateSet(t, aggs[1], channel2, map[testlang.NodeID][]transition.Transition{"y": {tr4}})
}

func checkAggregateSet(t *testing.T, set transition.AggregateSet, channel channel.Channel, smap transition.StateMap) {
	t.Helper()

	if !set.IsForChannel(channel) {
		t.Errorf("wrong channel: got %q, want %q", set.Channel, channel)
	}

	if !reflect.DeepEqual(set.States, smap) {
		t.Errorf("wrong state map: got %v, want %v", set.States, smap)
	}
}
