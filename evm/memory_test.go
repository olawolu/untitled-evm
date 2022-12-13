package evm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMemory_Get(t *testing.T) {
	offset := 0
	size := 32
	mem := NewMemory()
	mem.Resize(uint64(offset + size))

	mem.Set(uint64(offset), uint64(size), []byte{0x01, 0x02})
	assert.Equal(t, []byte{0x01, 0x02}, mem.Get(0, 2))
}

func TestMemory_Set(t *testing.T) {
	offset := 0
	size := 32
	mem := NewMemory()
	mem.Resize(uint64(offset+size))

	mem.Set(uint64(offset), uint64(size), []byte{0x01, 0x02})
	assert.Equal(t, []byte{0x01, 0x02}, mem.Get(0, 2))
}
