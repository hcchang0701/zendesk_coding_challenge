package fsm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockTransition struct{}

var _ ITransitionFunctions = (*mockTransition)(nil)

func (m *mockTransition) list() error {
	return nil
}

func (m *mockTransition) next() error {
	return nil
}

func (m *mockTransition) prev() error {
	return nil
}

func (m *mockTransition) selc(int) bool {
	return true
}

func (m *mockTransition) back() {}

func (m *mockTransition) quit() {}

func TestStateTransitions(t *testing.T) {
	tt = new(mockTransition)

	state = initial
	execute("quit")
	assert.Equal(t, quitNow, state)

	state = initial
	execute("list")
	execute("quit")
	assert.Equal(t, quitNow, state)

	state = initial
	execute("list")
	execute("selc")
	execute("quit")
	assert.Equal(t, quitNow, state)

	state = initial
	execute("list")
	assert.Equal(t, listAll, state)

	execute("next")
	assert.Equal(t, listAll, state)

	execute("prev")
	assert.Equal(t, listAll, state)

	execute("selc")
	assert.Equal(t, viewOne, state)

	execute("back")
	assert.Equal(t, listAll, state)
}
