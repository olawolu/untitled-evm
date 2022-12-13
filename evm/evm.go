package evm

import "github.com/holiman/uint256"

type EVM struct {
	context     *ExecutionContext
	interpreter *Interpreter
}

func NewEvm(ctx *ExecutionContext) *EVM {
	evm := &EVM{
		context: ctx,
	}
	evm.interpreter = NewInterpreter(evm)
	return evm
}

func (*EVM) execute(code []byte) (stack []uint256.Int, returnData string, success bool) {
	return
}

func Run(code []byte) (stack []uint256.Int, returnData string, success bool) {
	ctx := &ExecutionContext{}
	vm := NewEvm(ctx)
	stack, returnData, success = vm.execute(code)
	return
}
