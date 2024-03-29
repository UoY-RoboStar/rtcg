package stm

import (
	"fmt"

	"github.com/UoY-RoboStar/rtcg/internal/structure"
	"github.com/UoY-RoboStar/rtcg/internal/testlang"
	"github.com/UoY-RoboStar/rtcg/internal/testlang/rstype"
	"github.com/UoY-RoboStar/rtcg/internal/validate"
)

// Builder builds state machines from tests.
type Builder struct {
	nodeNum uint64                          // nodeNum is a monotonically increasing counter for naming unnamed nodes.
	stack   structure.Stack[*testlang.Node] // stack is a stack used for in-order test traversal.
	stm     Stm                             // stm is the state machine currently being built.
}

// BuildSuite builds a test suite tests into a map of state machines.
func (b *Builder) BuildSuite(tests validate.Suite) (*Suite, error) {
	var (
		suite Suite
		err   error
	)

	if suite.Tests, err = b.BuildMany(tests); err != nil {
		return nil, err
	}

	if suite.Types, err = unifyTypes(suite.Tests); err != nil {
		return nil, err
	}

	return &suite, nil
}

// unifyTypes returns a map of unified inferred types for each channel across each test in s.
func unifyTypes(tests map[string]*Stm) (TypeMap, error) {
	tmap := make(TypeMap)

	for _, test := range tests {
		for cName, cType := range test.Types {
			uType, err := rstype.Unify(cType, tmap[cName])
			if err != nil {
				return nil, fmt.Errorf("couldn't reconcile types for channel %s: %w", cName, err)
			}

			tmap[cName] = uType
		}
	}

	return tmap, nil
}

// BuildMany builds a STM for each test in the validated suite s.
func (b *Builder) BuildMany(s validate.Suite) (map[string]*Stm, error) {
	suite := make(map[string]*Stm, len(s))

	for name, test := range s {
		m, err := b.Build(test)
		if err != nil {
			return nil, fmt.Errorf("building %s: %w", name, err)
		}

		suite[name] = &m
	}

	return suite, nil
}

// Build builds a single state machine from the given validated test.
func (b *Builder) Build(test *validate.Test) (Stm, error) {
	b.nodeNum = 0
	b.stm = Stm{States: []*State{}, Tests: structure.NewSet[string](), Types: map[string]*rstype.RsType{}}

	testRoot := test.Root()
	testRoot.ID = "initial"

	b.stack.Clear()
	b.stack.Push(testRoot)

	for !b.stack.IsEmpty() {
		node := b.stack.Pop()

		if err := b.processNode(node); err != nil {
			return b.stm, err
		}
	}

	return b.stm, nil
}

func (b *Builder) processNode(node *testlang.Node) error {
	if node.Outcome == testlang.OutcomeFail {
		// We don't emit failing states.
		return nil
	}

	sn := b.buildState(node)
	b.stm.States = append(b.stm.States, sn)

	if err := b.inferNodeEventType(node); err != nil {
		return err
	}

	b.pushNext(node)

	return nil
}

func (b *Builder) pushNext(node *testlang.Node) {
	for i := range node.Next {
		b.stack.Push(&node.Next[i])
	}
}

func (b *Builder) inferNodeEventType(node *testlang.Node) error {
	if node.Event == nil {
		return nil
	}

	return b.inferEventType(node.Event)
}

func (b *Builder) inferEventType(event *testlang.Event) error {
	var err error

	chanName := event.Channel.Name

	if b.stm.Types[chanName], err = rstype.Unify(b.stm.Types[chanName], event.Value.Type()); err != nil {
		return fmt.Errorf("incompatible type information for %s: %w", chanName, err)
	}

	return nil
}

func (b *Builder) buildState(node *testlang.Node) *State {
	result := NewState(node.ID)

	for i := range node.Next {
		np := &node.Next[i]

		b.ensureNodeID(np)
		b.stm.Tests.Add(np.Tests...)
		result.AddOutgoingNode(np)
	}

	return result
}

func (b *Builder) ensureNodeID(n *testlang.Node) {
	if n.ID != "" {
		return
	}

	n.ID = testlang.NodeID(fmt.Sprintf("step%d", b.nodeNum))
	b.nodeNum++
}
