package evm

import (
	"fmt"

	"github.com/holiman/uint256"
)

func invalidOp(ctx *executionContext) []uint256.Int {
	fmt.Println("invalid opcode", ctx.code)
	return ctx.stack.data
}

func stopOp(ctx *executionContext) []uint256.Int {
	ctx.stopped = true
	return ctx.stack.data
}

func push1Op(ctx *executionContext) []uint256.Int {
	var value uint256.Int
	
	if ctx.pc < uint64(len(ctx.code)) {
		data := ctx.code[ctx.pc : ctx.pc+1]
		value.SetBytes(data)
		ctx.stack.Push(value)
	}
	return ctx.stack.data
}
