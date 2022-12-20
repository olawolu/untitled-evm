package evm

import (
	"github.com/holiman/uint256"
)

type (
	InstructionSet map[OpCode]Instruction
	executionFunc  func(ctx *executionContext) []uint256.Int
	memFunc        func(*Stack) (size uint64, overflow bool)
)

type Instruction struct {
	execute    executionFunc
	memorySize memFunc
}

func loadInstructionSet() InstructionSet {
	iSet := InstructionSet{
		STOP: {
			execute: stopOp,
		},
		ADD: {
			execute: addOp,
		},
		MUL: {
			execute: mulOp,
		},
		SUB: {
			execute: subOp,
		},
		DIV: {
			execute: divOp,
		},
		SDIV: {
			execute: sdivOp,
		},
		MOD: {
			execute: modOp,
		},
		SMOD: {
			execute: smodOp,
		},
		LT: {
			execute: ltOp,
		},
		GT: {
			execute: gtOp,
		},
		SLT: {
			execute: sltOp,
		},
		SGT: {
			execute: sgtOp,
		},
		EQ: {
			execute: eqOp,
		},
		ISZERO: {
			execute: iszeroOp,
		},
		AND: {
			execute: andOp,
		},
		OR: {
			execute: orOp,
		},
		XOR: {
			execute: xorOp,
		},
		NOT: {
			execute: notOp,
		},
		BYTE: {
			execute: byteOp,
		},
		POP: {
			execute: popOp,
		},
		PUSH1: {
			execute: push1Op,
		},
		PUSH2: {
			execute: makePush(2),
		},
		PUSH3: {
			execute: makePush(3),
		},
		PUSH4: {
			execute: makePush(4),
		},
		PUSH13: {
			execute: makePush(13),
		},
		PUSH20: {
			execute: makePush(20),
		},
		PUSH32: {
			execute: makePush(32),
		},
		DUP1: {
			execute: dup1Op,
		},
		DUP3: {
			execute: dup3Op,
		},
		SWAP1: {
			execute: swap1Op,
		},
		SWAP3: {
			execute: swap3Op,
		},
		JUMP: {
			execute: jumpOp,
		},
		JUMPI: {
			execute: jumpiOp,
		},
		JUMPDEST: {
			execute: jumpDestOp,
		},
		PC: {
			execute: pcOp,
		},
		MSTORE: {
			execute:    mstoreOp,
			memorySize: memoryMStore,
		},
		MLOAD: {
			execute:    mloadOp,
			memorySize: memoryMLoad,
		},
		MSTORE8: {
			execute:    mstore8Op,
			memorySize: memoryMStore8,
		},
		MSIZE: {
			execute: msizeOp,
		},
		SHA3: {
			execute:    sha3Op,
			memorySize: memorySha3,
		},
		ADDRESS: {
			execute: addressOp,
		},
		CALLER: {
			execute: callerOp,
		},
		BALANCE: {
			execute: balanceOp,
		},
		ORIGIN: {
			execute: originOp,
		},
		COINBASE: {
			execute: coinbaseOp,
		},
		TIMESTAMP: {
			execute: timestampOp,
		},
		NUMBER: {
			execute: numberOp,
		},
		DIFFICULTY: {
			execute: difficultyOp,
		},
		GASLIMIT: {
			execute: gaslimitOp,
		},
		GASPRICE: {
			execute: gaspriceOp,
		},
		CHAINID: {
			execute: chainidOp,
		},
		CALLVALUE: {
			execute: callvalueOp,
		},
		CALLDATALOAD: {
			execute: calldataloadOp,
		},
		CALLDATASIZE: {
			execute: calldatasizeOp,
		},
		CALLDATACOPY: {
			execute:    calldatacopyOp,
			memorySize: memoryCallDataCopy,
		},
		CODESIZE: {
			execute: codesizeOp,
		},
		CODECOPY: {
			execute:    codecopyOp,
			memorySize: memoryCodeCopy,
		},
		EXTCODESIZE: {
			execute: extcodesizeOp,
		},
		EXTCODECOPY: {
			execute:    extcodecopyOp,
			memorySize: memoryExtCodeCopy,
		},
		SELFBALANCE: {
			execute: selfbalanceOp,
		},
		SSTORE: {
			execute: sstoreOp,
		},
		SLOAD: {
			execute: sloadOp,
		},
		RETURN: {
			execute:    returnOp,
			memorySize: memoryReturn,
		},
		REVERT: {
			execute:    revertOp,
			memorySize: memoryRevert,
		},
		CALL: {
			execute:    callOp,
			memorySize: memoryCall,
		},
		CREATE: {
			execute:    createOp,
			memorySize: memoryCreate,
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
