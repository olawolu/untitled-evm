package evm

import (
	"fmt"
	"math"
	"math/bits"

	"github.com/holiman/uint256"
)

type VM struct {
	Context      *executionContext
	instructions InstructionSet
}

func NewVm(ctx *executionContext) *VM {
	evm := &VM{
		Context:      ctx,
		instructions: loadInstructionSet(),
	}
	evm.Context.vm = evm
	return evm
}

func (vm *VM) Execute(code []byte) ([]uint256.Int, string, bool) {
	var stack []uint256.Int

	for !vm.Context.stopped && vm.Context.pc < uint64(len(code)) {
		// Read the next opcode
		opCode := decodeOp(vm.Context)
		op := vm.instructions[opCode]

		var memorySize uint64
		if op.memorySize != nil {
			memSize, overflow := op.memorySize(vm.Context.stack)
			if overflow {
				fmt.Println("memory size overflow")
				return stack, "", false
			}

			if memorySize, overflow = SafeMul(toWordSize(memSize), 32); overflow {
				fmt.Println("memory size overflow")
				return stack, "", false
			}

			if memorySize > 0 {
				vm.Context.memory.Resize(memorySize)
			}
		}

		stack = op.execute(vm.Context)

		// vm.Context.pc++
	}

	if vm.Context.pc >= uint64(len(code)) && vm.Context.stopped {
		vm.Context.done = true
	}

	return stack, vm.Context.returnData, vm.Context.done
}

func Run(code []byte, t *Transaction, state interface{}, block *Block) (stack []uint256.Int, returnData string, success bool) {
	ctx := NewExecutionContext(code, t, state, block)
	vm := NewVm(ctx)
	stack, returnData, success = vm.Execute(code)
	return
}

func decodeOp(ctx *executionContext) OpCode {
	op := ctx.ReadCode(1)
	return OpCode(op[0])
}

// SafeMul returns x*y and checks for overflow.
func SafeMul(x, y uint64) (uint64, bool) {
	hi, lo := bits.Mul64(x, y)
	return lo, hi != 0
}

func toWordSize(size uint64) uint64 {
	if size > math.MaxUint64-31 {
		return math.MaxUint64/32 + 1
	}

	return (size + 31) / 32
}
