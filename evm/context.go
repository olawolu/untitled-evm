package evm

import (
	"encoding/hex"
	"fmt"
	"strconv"

	"github.com/holiman/uint256"
)

type Block struct {
	Coinbase   string
	Timestamp  string
	Number     string
	Difficulty string
	GasLimit   string
	ChainId    string
}

type Transaction struct {
	To       string
	From     string
	Origin   string
	GasPrice string
	Value    string
	Data     string
}

type executionContext struct {
	pc          uint64
	vm          *VM
	done        bool
	code        []byte
	block       *Block
	stack       *Stack
	state       interface{}
	memory      *Memory
	storage     *Storage
	stopped     bool
	returnData  string
	transaction *Transaction
}

func NewExecutionContext(code []byte, t *Transaction, state interface{}, block *Block) *executionContext {
	return &executionContext{
		pc:          0,
		code:        code,
		stack:       NewStack(),
		state:       state,
		block:       block,
		memory:      NewMemory(),
		storage:     NewStorage(),
		transaction: t,
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

func (ctx *executionContext) GetBalance(address string) *uint256.Int {
	if ctx.state == nil {
		return uint256.NewInt(0)
	}
	obj, _ := ctx.state.(map[string]interface{})

	balanceI, _ := obj[address].(map[string]interface{})

	balance := balanceI["balance"].(string)

	uBalance, err := strconv.ParseUint(balance, 10, 64)
	if err != nil {
		fmt.Println("Error", err)
	}

	uBalance256 := new(uint256.Int).SetUint64(uBalance)

	return uBalance256
}

func (ctx *executionContext) GetCodeSize(address string) uint64 {
	if ctx.state == nil {
		return 0
	}
	obj, _ := ctx.state.(map[string]interface{})

	codeI, _ := obj[address].(map[string]interface{})

	code := codeI["code"].(map[string]interface{})

	bin := code["bin"].(string)

	binBytes, err := hex.DecodeString(bin)
	if err != nil {
		fmt.Println("Error", err)
	}

	return uint64(len(binBytes))
}

func (ctx *executionContext) GetCode(address string) []byte {
	if ctx.state == nil {
		return []byte{}
	}

	obj, _ := ctx.state.(map[string]interface{})

	codeI, _ := obj[address].(map[string]interface{})

	code := codeI["code"].(map[string]interface{})

	bin := code["bin"].(string)

	binBytes, err := hex.DecodeString(bin)
	if err != nil {
		fmt.Println("Error", err)
	}

	return binBytes
}
