package evm

import (
	"testing"

	"github.com/holiman/uint256"
	"github.com/stretchr/testify/assert"
)

func TestStack_NewStack(t *testing.T) {
	stack := NewStack()

	assert.Equal(t, 0, stack.Depth())
	assert.Equal(t, 0, len(stack.Data()))
}

func TestStack_Push(t *testing.T) {
	stack := NewStack()
	stack.Push(*uint256.NewInt(1))
	stack.Push(*uint256.NewInt(2))
	stack.Push(*uint256.NewInt(3))

	assert.Equal(t, uint256.NewInt(3), stack.Peek())
}

func TestStack_Pop(t *testing.T) {
	stack := NewStack()
	stack.Push(*uint256.NewInt(1))
	stack.Push(*uint256.NewInt(2))
	stack.Push(*uint256.NewInt(3))

	v := stack.Pop()

	assert.Equal(t, uint256.NewInt(3), &v)
}

func TestStack_Swap(t *testing.T) {
	stack := NewStack()
	stack.Push(*uint256.NewInt(1))
	stack.Push(*uint256.NewInt(2))
	stack.Push(*uint256.NewInt(3))

	stack.Swap(1)

	assert.Equal(t, uint256.NewInt(2), stack.Peek())
}

func TestStack_Peek(t *testing.T) {
	stack := NewStack()
	stack.Push(*uint256.NewInt(1))
	stack.Push(*uint256.NewInt(2))
	stack.Push(*uint256.NewInt(3))

	assert.Equal(t, uint256.NewInt(3), stack.Peek())
}

func TestStack_PeekN(t *testing.T) {
	stack := NewStack()
	stack.Push(*uint256.NewInt(1))
	stack.Push(*uint256.NewInt(2))
	stack.Push(*uint256.NewInt(3))

	assert.Equal(t, uint256.NewInt(2), stack.PeekN(1))
}

func TestStack_Back(t *testing.T) {
	stack := NewStack()
	stack.Push(*uint256.NewInt(1))
	stack.Push(*uint256.NewInt(2))
	stack.Push(*uint256.NewInt(3))

	assert.Equal(t, uint256.NewInt(1), stack.Back(2))
}
