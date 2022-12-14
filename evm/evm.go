package evm

import "github.com/holiman/uint256"

type VM struct {
	context      *executionContext
	instructions InstructionSet
}

func NewVm(ctx *executionContext) *VM {
	evm := &VM{
		context:      ctx,
		instructions: loadInstructionSet(),
	}
	return evm
}

func (vm *VM) Execute(code []byte) ([]uint256.Int, string, bool) {
	var stack []uint256.Int

	for !vm.context.stopped && vm.context.pc < uint64(len(code)) {
		// Read the next opcode
		opCode := decodeOp(vm.context)
		op := vm.instructions[opCode]

		stack = op.execute(vm.context)
		vm.context.pc++
	}

	return stack, "", vm.context.stopped
}

func Run(code []byte) (stack []uint256.Int, returnData string, success bool) {
	ctx := NewExecutionContext(code)
	vm := NewVm(ctx)
	stack, returnData, success = vm.Execute(code)
	return
}

func decodeOp(ctx *executionContext) OpCode {
	op := ctx.ReadCode(1)
	return OpCode(op[0])
}