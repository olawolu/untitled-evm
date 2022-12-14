package evm_test

import (
	"testing"

	"github.com/helicarrierstudio/untitled-evm/evm"
	"github.com/stretchr/testify/assert"
)

func Test_executionContext_ReadCode(t *testing.T) {
	code := []byte{0x01, 0x02, 0x03, 0x04, 0x05}
	ctx := evm.NewExecutionContext(code)
	result := ctx.ReadCode(2)
	assert.Equal(t, []byte{0x01, 0x02}, result)
	assert.Equal(t, uint64(2), ctx.Counter())
}
