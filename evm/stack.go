package evm

import "github.com/holiman/uint256"

/** Stack is a simple stack implementation
 It is not thread safe
 data is stored in a slice as LIFO
 e.g. [1, 2, 3] is treated as
 		|3|
 		|2|
 		|1|
 means 3 is the on top of the stack at index 0
*/
type Stack struct {
	depth int
	data  []uint256.Int
}

func NewStack() *Stack {
	return &Stack{
		depth: 0,
		data:  make([]uint256.Int, 0),
	}
}

func (s *Stack) Push(v uint256.Int) {
	if s.depth > 0 {
		s.data = append([]uint256.Int{v}, s.data...)
	} else {
		s.data = append(s.data, v)
	}
	s.depth++
}

func (s *Stack) Pop() uint256.Int {
	v := s.data[0]
	s.data = s.data[1:]
	s.depth--
	return v
}

func (s *Stack) Swap(n int) {
	s.data[0], s.data[n] = s.data[n], s.data[0]
}

func (s *Stack) Peek() *uint256.Int {
	return &s.data[0]
}

func (s *Stack) PeekN(n int) *uint256.Int {
	return &s.data[n]
}

func (s *Stack) Back(n int) *uint256.Int {
	return &s.data[n]
}

func (s *Stack) Depth() int {
	return s.depth
}

func (s *Stack) Data() []uint256.Int {
	return s.data
}
