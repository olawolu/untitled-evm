package evm

type ExecutionContext struct{}

type Interpreter struct {
	evm *EVM
}

func NewInterpreter(evm *EVM) *Interpreter {
	return &Interpreter{
		evm: evm,
	}
}
