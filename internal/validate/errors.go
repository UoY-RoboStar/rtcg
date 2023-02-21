package validate

import "errors"

var (
	// ErrBadOutcome occurs when a non-root node has an out-of-range outcome.
	ErrBadOutcome = errors.New("non-root node has an invalid outcome set")

	// ErrFailHasEvent occurs when a failing node has an event.
	ErrFailHasEvent = errors.New("failing node should not have an event")

	// ErrFailHasNextNodes occurs when a failing node has next nodes.
	ErrFailHasNextNodes = errors.New("failing node should not have next tests")

	// ErrNeedOneNode occurs when a node doesn't have one next node, but should.
	ErrNeedOneNode = errors.New("expected this node to have exactly one next node")

	// ErrNoEvent occurs when a non-failing non-root node has no event.
	ErrNoEvent = errors.New("non-root, non-failing node should have an event set")

	// ErrNoNextNodes occurs when a non-failing node has no onwards nodes.
	ErrNoNextNodes = errors.New("non-failing node should have at least one next node")

	// ErrNoOutcome occurs when a non-root node has no outcome.
	ErrNoOutcome = errors.New("non-root node should have an outcome set")

	// ErrNoTests occurs when a node has no tests set.
	ErrNoTests = errors.New("node should belong to at least one test")

	// ErrRootHasEvent occurs when a root node has an event set.
	ErrRootHasEvent = errors.New("root should not have an event set")

	// ErrRootHasOutcome occurs when a root node has an outcome set.
	ErrRootHasOutcome = errors.New("root should not have an outcome set")
)
