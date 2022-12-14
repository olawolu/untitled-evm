package evm

import "github.com/holiman/uint256"

type (
	InstructionSet map[OpCode]Instruction
	executionFunc  func(ctx *executionContext) []uint256.Int
)

type Instruction struct {
	execute executionFunc
}

func loadInstructionSet() InstructionSet {
	iSet := InstructionSet{
		STOP:       {
			execute: stopOp,
		},
		ADD:        {},
		MUL:        {},
		SUB:        {},
		DIV:        {},
		SDIV:       {},
		MOD:        {},
		SMOD:       {},
		ADDMOD:     {},
		MULMOD:     {},
		EXP:        {},
		SIGNEXTEND: {},
		PUSH1:      {
			execute: push1Op,
		},
	}
	return validateInstructionSet(iSet)
}

func validateInstructionSet(iSet InstructionSet) InstructionSet {
	for i := 0; i < 256; i++ {
		if _, ok := iSet[OpCode(i)]; !ok {
			iSet[OpCode(i)] = Instruction{
				execute: invalidOp,
			}
		}
	}
	return iSet
}
