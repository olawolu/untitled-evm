package evm

type executionContext struct {
	pc      uint64
	stopped bool
	code    []byte
	stack   *Stack
	memory  *Memory
}

func NewExecutionContext(code []byte) *executionContext {
	return &executionContext{
		pc:     0,
		code:   code,
		stack:  NewStack(),
		memory: NewMemory(),
	}
}

func (ctx *executionContext) ReadCode(n uint64) []byte {
	value := ctx.code[ctx.pc : ctx.pc+n]
	ctx.pc += n
	return value
}

func (ctx *executionContext) Stop() {
	ctx.stopped = true
}

func (ctx *executionContext) Counter() uint64 {
	return ctx.pc
}