package evm

import (
	"fmt"

	"github.com/holiman/uint256"
)

/** Memory is a simple memory implementation
It is not thread safe
*/

type Memory struct {
	data []byte
}

func NewMemory() *Memory {
	return &Memory{
		data: make([]byte, 0),
	}
}

func (m *Memory) Get(offset, size uint64) []byte {
	if len(m.data) > int(offset) {
		return m.data[offset : offset+size]
	}
	return nil
}

func (m *Memory) Set(offset, size uint64, value []byte) {
	if offset+size > uint64(len(m.data)) {
		panic("memory limit reached: requested size is larger than available memory")
	}

	if len(value) > len(m.data) {
		m.data = append(m.data, make([]byte, uint64(len(value)))...)
	}
	copy(m.data[offset:], value[:])
}

func (m *Memory) Set32(offset uint64, val *uint256.Int) {
	// length of store may never be less than offset + size.
	// The store should be resized PRIOR to setting the memory
	fmt.Println("Set32", offset, len(m.data))
	if offset+32 > uint64(len(m.data)) {
		panic("memory limit reached: requested size is larger than available memory")
	}
	// Fill in relevant bits
	b32 := val.Bytes32()
	copy(m.data[offset:], b32[:])
}

func (m *Memory) Resize(size uint64) {
	if uint64(len(m.data)) <= size {
		m.data = append(m.data, make([]byte, size)...)
	}
}
