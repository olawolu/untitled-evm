package evm

import (
	"encoding/hex"
	"fmt"
	"os"
	"strconv"

	"github.com/blocktree/openwallet/crypto"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/holiman/uint256"
)

func invalidOp(ctx *executionContext) []uint256.Int {
	fmt.Println("invalid opcode", ctx.code[ctx.pc-1])
	os.Exit(1)
	return ctx.stack.data
}

func stopOp(ctx *executionContext) []uint256.Int {
	ctx.stopped = true
	return ctx.stack.data
}

func push1Op(ctx *executionContext) []uint256.Int {
	var value uint256.Int

	if ctx.pc < uint64(len(ctx.code)) {
		data := ctx.ReadCode(1)
		value.SetBytes(data)
		ctx.stack.Push(value)
	}
	return ctx.stack.data
}

func makePush(n uint64) func(ctx *executionContext) []uint256.Int {
	var value uint256.Int
	return func(ctx *executionContext) []uint256.Int {
		data := ctx.ReadCode(n)
		value.SetBytes(data)
		ctx.stack.Push(value)
		return ctx.stack.data
	}
}

func popOp(ctx *executionContext) []uint256.Int {
	ctx.stack.Pop()
	return ctx.stack.data
}

func addOp(ctx *executionContext) []uint256.Int {
	var result uint256.Int
	x, y := ctx.stack.Pop(), ctx.stack.Pop()
	result.AddOverflow(&x, &y)
	ctx.stack.Push(result)
	return ctx.stack.data
}

func mulOp(ctx *executionContext) []uint256.Int {
	var result uint256.Int
	x, y := ctx.stack.Pop(), ctx.stack.Pop()
	result.MulOverflow(&x, &y)
	ctx.stack.Push(result)
	return ctx.stack.data
}

func subOp(ctx *executionContext) []uint256.Int {
	var result uint256.Int
	x, y := ctx.stack.Pop(), ctx.stack.Pop()
	result.SubOverflow(&x, &y)
	ctx.stack.Push(result)
	return ctx.stack.data
}

func divOp(ctx *executionContext) []uint256.Int {
	var result uint256.Int
	x, y := ctx.stack.Pop(), ctx.stack.Pop()
	result.Div(&x, &y)
	ctx.stack.Push(result)
	return ctx.stack.data
}

func sdivOp(ctx *executionContext) []uint256.Int {
	var result uint256.Int
	x, y := ctx.stack.Pop(), ctx.stack.Pop()
	result.SDiv(&x, &y)
	ctx.stack.Push(result)
	return ctx.stack.data
}

func modOp(ctx *executionContext) []uint256.Int {
	var result, a, b uint256.Int

	a = ctx.stack.Pop()
	b = ctx.stack.Pop()
	result.Mod(&a, &b)
	ctx.stack.Push(result)
	return ctx.stack.data
}

func smodOp(ctx *executionContext) []uint256.Int {
	var result, a, b uint256.Int

	a = ctx.stack.Pop()
	b = ctx.stack.Pop()
	result.SMod(&a, &b)
	ctx.stack.Push(result)
	return ctx.stack.data
}

func ltOp(ctx *executionContext) []uint256.Int {
	var result, a, b uint256.Int

	a = ctx.stack.Pop()
	b = ctx.stack.Pop()
	if a.Lt(&b) {
		result = *new(uint256.Int).SetBytes([]byte{1})
	} else {
		result = *new(uint256.Int).SetBytes([]byte{0})
	}
	ctx.stack.Push(result)
	return ctx.stack.data
}

func sltOp(ctx *executionContext) []uint256.Int {
	var result, a, b uint256.Int

	a = ctx.stack.Pop()
	b = ctx.stack.Pop()

	if a.Slt(&b) {
		result = *new(uint256.Int).SetBytes([]byte{1})
	} else {
		result = *new(uint256.Int).SetBytes([]byte{0})
	}
	ctx.stack.Push(result)
	return ctx.stack.data
}

func gtOp(ctx *executionContext) []uint256.Int {
	var result, a, b uint256.Int

	a = ctx.stack.Pop()
	b = ctx.stack.Pop()
	if a.Gt(&b) {
		result = *new(uint256.Int).SetBytes([]byte{1})
	} else {
		result = *new(uint256.Int).SetBytes([]byte{0})
	}
	ctx.stack.Push(result)
	return ctx.stack.data
}

func sgtOp(ctx *executionContext) []uint256.Int {
	var result, a, b uint256.Int

	a = ctx.stack.Pop()
	b = ctx.stack.Pop()
	if a.Sgt(&b) {
		result = *new(uint256.Int).SetBytes([]byte{1})
	} else {
		result = *new(uint256.Int).SetBytes([]byte{0})
	}
	ctx.stack.Push(result)
	return ctx.stack.data
}

func eqOp(ctx *executionContext) []uint256.Int {
	var result, a, b uint256.Int

	a = ctx.stack.Pop()
	b = ctx.stack.Pop()
	if a.Eq(&b) {
		result = *new(uint256.Int).SetBytes([]byte{1})
	} else {
		result = *new(uint256.Int).SetBytes([]byte{0})
	}
	ctx.stack.Push(result)

	return ctx.stack.data
}

func iszeroOp(ctx *executionContext) []uint256.Int {
	var result, a uint256.Int

	a = ctx.stack.Pop()
	if a.IsZero() {
		result = *new(uint256.Int).SetBytes([]byte{1})
	} else {
		result = *new(uint256.Int).SetBytes([]byte{0})
	}
	ctx.stack.Push(result)
	return ctx.stack.data
}

func andOp(ctx *executionContext) []uint256.Int {
	var result, a, b uint256.Int

	a = ctx.stack.Pop()
	b = ctx.stack.Pop()
	result.And(&a, &b)
	ctx.stack.Push(result)
	return ctx.stack.data
}

func orOp(ctx *executionContext) []uint256.Int {
	var result, a, b uint256.Int

	a = ctx.stack.Pop()
	b = ctx.stack.Pop()
	result.Or(&a, &b)
	ctx.stack.Push(result)
	return ctx.stack.data
}

func xorOp(ctx *executionContext) []uint256.Int {
	var result, a, b uint256.Int

	a = ctx.stack.Pop()
	b = ctx.stack.Pop()
	result.Xor(&a, &b)
	ctx.stack.Push(result)
	return ctx.stack.data
}

func notOp(ctx *executionContext) []uint256.Int {
	var result, a uint256.Int

	a = ctx.stack.Pop()
	result.Not(&a)
	ctx.stack.Push(result)
	return ctx.stack.data
}

func byteOp(ctx *executionContext) []uint256.Int {
	var result, a, b uint256.Int

	a = ctx.stack.Pop()
	b = ctx.stack.Pop()
	result = *b.Byte(&a)
	ctx.stack.Push(result)
	ctx.pc++
	return ctx.stack.data
}

func dup1Op(ctx *executionContext) []uint256.Int {
	result := ctx.stack.Peek()
	ctx.stack.Push(*result)
	return ctx.stack.data
}

func dup3Op(ctx *executionContext) []uint256.Int {
	result := ctx.stack.PeekN(2)
	ctx.stack.Push(*result)
	return ctx.stack.data
}

func swap1Op(ctx *executionContext) []uint256.Int {
	ctx.stack.Swap(1)
	return ctx.stack.data
}

func swap3Op(ctx *executionContext) []uint256.Int {
	ctx.stack.Swap(3)
	return ctx.stack.data
}

func jumpOp(ctx *executionContext) []uint256.Int {
	ctx.stack.Pop()
	ctx.pc += 2
	return ctx.stack.data
}

func jumpiOp(ctx *executionContext) []uint256.Int {
	pos, cond := ctx.stack.Pop(), ctx.stack.Pop()
	if !cond.IsZero() {
		if !validJumpDest(ctx.code, &pos) {
			return ctx.stack.data
		}
		ctx.pc = pos.Uint64()
	}
	return ctx.stack.data
}

func validJumpDest(code []byte, dest *uint256.Int) bool {
	udest, overflow := dest.Uint64WithOverflow()
	if overflow || udest >= uint64(len(code)) {
		return false
	}
	if OpCode(code[udest]) != JUMPDEST {
		return false
	}
	return true
}

func jumpDestOp(ctx *executionContext) []uint256.Int {
	return ctx.stack.data
}

func pcOp(ctx *executionContext) []uint256.Int {
	pointer := ctx.pc - 1
	ctx.stack.Push(*new(uint256.Int).SetUint64(pointer))
	return ctx.stack.data
}

func mstoreOp(ctx *executionContext) []uint256.Int {
	_offset, _val := ctx.stack.Pop(), ctx.stack.Pop()
	offset := _offset.Uint64()

	ctx.memory.Set32(offset, &_val)

	return ctx.stack.data
}
func mloadOp(ctx *executionContext) []uint256.Int {
	var result uint256.Int
	offset := ctx.stack.Pop()
	content := ctx.memory.Get(offset.Uint64(), 32)
	result.SetBytes(content)
	ctx.stack.Push(result)

	return ctx.stack.data
}

func mstore8Op(ctx *executionContext) []uint256.Int {
	offset, val := ctx.stack.Pop(), ctx.stack.Pop()
	ctx.memory.data[offset.Uint64()] = byte(val.Uint64())
	return ctx.stack.data
}

func msizeOp(ctx *executionContext) []uint256.Int {
	ctx.stack.Push(*new(uint256.Int).SetUint64(uint64(len(ctx.memory.data))))
	return ctx.stack.data
}

func sha3Op(ctx *executionContext) []uint256.Int {
	var result uint256.Int
	offset, size := ctx.stack.Pop(), ctx.stack.Pop()
	content := ctx.memory.Get(offset.Uint64(), size.Uint64())
	result.SetBytes(crypto.Keccak256(content))
	ctx.stack.Push(result)
	return ctx.stack.data
}

func addressOp(ctx *executionContext) []uint256.Int {
	val, err := uint256.FromHex(ctx.transaction.To)
	if err != nil {
		fmt.Println("Error", err)
	}

	ctx.stack.Push(*val)

	return ctx.stack.data
}

func callerOp(ctx *executionContext) []uint256.Int {
	val, err := uint256.FromHex(ctx.transaction.From)
	if err != nil {
		fmt.Println("Error", err)
	}
	ctx.stack.Push(*val)
	return ctx.stack.data
}

func balanceOp(ctx *executionContext) []uint256.Int {
	address := ctx.stack.Pop()
	balance := ctx.GetBalance(address.String())
	ctx.stack.Push(*balance)

	fmt.Println("stack", ctx.stack.data)
	return ctx.stack.data
}

func originOp(ctx *executionContext) []uint256.Int {
	val, err := uint256.FromHex(ctx.transaction.Origin)
	if err != nil {
		fmt.Println("Error", err)
	}
	ctx.stack.Push(*val)
	return ctx.stack.data
}

func coinbaseOp(ctx *executionContext) []uint256.Int {
	val, err := uint256.FromHex(ctx.block.Coinbase)
	if err != nil {
		fmt.Println("Error", err)
	}
	ctx.stack.Push(*val)
	return ctx.stack.data
}

func timestampOp(ctx *executionContext) []uint256.Int {
	timeStampUint, err := strconv.ParseUint(ctx.block.Timestamp, 10, 64)
	if err != nil {
		fmt.Println("Error", err)
	}
	ctx.stack.Push(*new(uint256.Int).SetUint64(timeStampUint))
	return ctx.stack.data
}

func numberOp(ctx *executionContext) []uint256.Int {
	numberUint, err := strconv.ParseUint(ctx.block.Number, 10, 64)
	if err != nil {
		fmt.Println("Error", err)
	}
	ctx.stack.Push(*new(uint256.Int).SetUint64(numberUint))
	return ctx.stack.data
}

func difficultyOp(ctx *executionContext) []uint256.Int {
	difficultyUint, err := uint256.FromHex(ctx.block.Difficulty)
	if err != nil {
		fmt.Println("Error", err)
	}

	ctx.stack.Push(*difficultyUint)
	return ctx.stack.data
}

func gaslimitOp(ctx *executionContext) []uint256.Int {
	gasLimitUint, err := uint256.FromHex(ctx.block.GasLimit)
	if err != nil {
		fmt.Println("Error", err)
	}

	ctx.stack.Push(*gasLimitUint)
	return ctx.stack.data
}

func gaspriceOp(ctx *executionContext) []uint256.Int {
	gasPriceUint, err := strconv.ParseUint(ctx.transaction.GasPrice, 10, 64)
	if err != nil {
		fmt.Println("Error", err)
	}
	ctx.stack.Push(*new(uint256.Int).SetUint64(gasPriceUint))
	return ctx.stack.data
}

func chainidOp(ctx *executionContext) []uint256.Int {
	chainIdUint, err := strconv.ParseUint(ctx.block.ChainId, 10, 64)
	if err != nil {
		fmt.Println("Error", err)
	}
	ctx.stack.Push(*new(uint256.Int).SetUint64(chainIdUint))
	return ctx.stack.data
}

func callvalueOp(ctx *executionContext) []uint256.Int {
	callValueUint, err := strconv.ParseUint(ctx.transaction.Value, 10, 64)
	if err != nil {
		fmt.Println("Error", err)
	}
	ctx.stack.Push(*new(uint256.Int).SetUint64(callValueUint))
	return ctx.stack.data
}

func calldataloadOp(ctx *executionContext) []uint256.Int {
	x := ctx.stack.Peek()

	if o, overflow := x.Uint64WithOverflow(); !overflow {
		b, err := hex.DecodeString(ctx.transaction.Data)
		if err != nil {
			fmt.Println("Error", err)
		}
		dt := getData(b, o, 32)
		x.SetBytes(dt)
	} else {
		x.Clear()
	}

	return ctx.stack.data
}

func getData(data []byte, offset, size uint64) []byte {
	dataLen := uint64(len(data))
	if offset > dataLen {
		offset = dataLen
	}

	end := offset + size
	if end > dataLen {
		end = dataLen
	}

	return rightPadData(data[offset:end], int(size))

}

func rightPadData(data []byte, size int) []byte {
	if size < len(data) {
		return data
	}
	padded := make([]byte, size)
	copy(padded, data)
	return padded
}

func calldatasizeOp(ctx *executionContext) []uint256.Int {
	data, err := hex.DecodeString(ctx.transaction.Data)
	if err != nil {
		fmt.Println("Error", err)
	}
	dataLen := uint64(len(data))
	ctx.stack.Push(*new(uint256.Int).SetUint64(dataLen))
	return ctx.stack.data
}

func calldatacopyOp(ctx *executionContext) []uint256.Int {
	mOffset := ctx.stack.Pop()
	offset := ctx.stack.Pop()
	mSize := ctx.stack.Pop()

	if o, overflow := offset.Uint64WithOverflow(); !overflow {
		b, err := hex.DecodeString(ctx.transaction.Data)
		if err != nil {
			fmt.Println("Error", err)
		}
		ctx.memory.Set(mOffset.Uint64(), mSize.Uint64(), getData(b, o, mSize.Uint64()))
	} else {
		mOffset.Clear()
	}

	return ctx.stack.data
}

func codesizeOp(ctx *executionContext) []uint256.Int {
	codeSize := uint64(len(ctx.code))
	ctx.stack.Push(*new(uint256.Int).SetUint64(codeSize))
	return ctx.stack.data
}

func codecopyOp(ctx *executionContext) []uint256.Int {
	mOffset := ctx.stack.Pop()
	offset := ctx.stack.Pop()
	mSize := ctx.stack.Pop()

	if o, overflow := offset.Uint64WithOverflow(); !overflow {
		ctx.memory.Set(mOffset.Uint64(), mSize.Uint64(), getData(ctx.code, o, mSize.Uint64()))
	} else {
		mOffset.Clear()
	}

	return ctx.stack.data
}

func extcodesizeOp(ctx *executionContext) []uint256.Int {
	address := ctx.stack.Pop()
	addressHex := address.String()

	codeSize := ctx.GetCodeSize(addressHex)
	ctx.stack.Push(*new(uint256.Int).SetUint64(codeSize))
	return ctx.stack.data
}

func extcodecopyOp(ctx *executionContext) []uint256.Int {
	address := ctx.stack.Pop()
	mOffset := ctx.stack.Pop()
	offset := ctx.stack.Pop()
	mSize := ctx.stack.Pop()

	addressHex := address.String()
	code := ctx.GetCode(addressHex)

	if o, overflow := offset.Uint64WithOverflow(); !overflow {
		ctx.memory.Set(mOffset.Uint64(), mSize.Uint64(), getData(code, o, mSize.Uint64()))
	} else {
		mOffset.Clear()
	}

	return ctx.stack.data
}

func selfbalanceOp(ctx *executionContext) []uint256.Int {
	balance := ctx.GetBalance(ctx.transaction.To)
	ctx.stack.Push(*balance)
	return ctx.stack.data
}

func sstoreOp(ctx *executionContext) []uint256.Int {
	key := ctx.stack.Pop()
	value := ctx.stack.Pop()

	keyBytes := key.Bytes32()
	valueBytes := value.Bytes()

	ctx.storage.Set(keyBytes, valueBytes)
	return ctx.stack.data
}

func sloadOp(ctx *executionContext) []uint256.Int {
	var result uint256.Int
	key := ctx.stack.Pop()

	keyBytes := key.Bytes32()
	valueBytes := ctx.storage.Get(keyBytes)

	result.SetBytes(valueBytes)
	ctx.stack.Push(result)
	return ctx.stack.data
}

func returnOp(ctx *executionContext) []uint256.Int {
	var returnData []byte
	offset := ctx.stack.Pop()
	size := ctx.stack.Pop()

	if o, overflow := offset.Uint64WithOverflow(); !overflow {
		if s, overflow := size.Uint64WithOverflow(); !overflow {
			returnData = ctx.memory.Get(o, s)
		}
	}

	trf := hex.EncodeToString(returnData)

	ctx.returnData = trf
	ctx.done = true
	return ctx.stack.data
}

func revertOp(ctx *executionContext) []uint256.Int {
	var returnData []byte
	offset := ctx.stack.Pop()
	size := ctx.stack.Pop()

	if o, overflow := offset.Uint64WithOverflow(); !overflow {
		if s, overflow := size.Uint64WithOverflow(); !overflow {
			returnData = ctx.memory.Get(o, s)
		}
	}

	ctx.returnData = hex.EncodeToString(returnData)
	return ctx.stack.data
}

func callOp(ctx *executionContext) []uint256.Int {
	_ = ctx.stack.Pop()
	to := ctx.stack.Pop()
	value := ctx.stack.Pop()
	inOffset := ctx.stack.Pop()
	inSize := ctx.stack.Pop()
	outOffset := ctx.stack.Pop()
	outSize := ctx.stack.Pop()
	fmt.Println("outsize", outSize)
	toHex := to.String()
	_ = value.String()
	_ = inOffset.String()
	_ = inSize.String()
	_ = outOffset.String()
	_ = outSize.String()

	stateObj, _ := ctx.state.(map[string]interface{})
	addrObj, _ := stateObj[toHex].(map[string]interface{})
	codeObj, _ := addrObj["code"].(map[string]interface{})
	codeStr, _ := codeObj["bin"].(string)

	fromAddress := ctx.transaction.To

	bin, err := hex.DecodeString(codeStr)
	if err != nil {
		panic(err)
	}
	vm := ctx.vm

	// pause the current context and pass execution to a new subcontext
	_ = ctx
	pc := uint64(0)

	newCtx := &executionContext{
		pc:      pc,
		memory:  NewMemory(),
		stack:   NewStack(),
		storage: NewStorage(),
		code:    bin,
		state:   ctx.state,
		transaction: &Transaction{
			From: fromAddress,
		},
	}

	vm.Context = newCtx

	_, _, success := vm.Execute(bin)
	switch success {
	case true:
		ctx.stack.Push(*uint256.NewInt(1))
	case false:
		ctx.stack.Push(*uint256.NewInt(0))
	}

	// resume the parent context
	memSize := len(vm.Context.memory.data)
	ctx.memory.Resize(uint64(memSize))
	ctx.memory = vm.Context.memory
	vm.Context = ctx
	return ctx.stack.data
}

func createOp(ctx *executionContext) []uint256.Int {
	value := ctx.stack.Pop()
	offset := ctx.stack.Pop()
	size := ctx.stack.Pop()

	senderAddress := ctx.transaction.To

	rlpA, err := rlp.EncodeToBytes([]string{senderAddress, "0x0"})
	if err != nil {
		panic(err)
	}

	bin := ctx.memory.Get(offset.Uint64(), size.Uint64())

	addr, err := uint256.FromHex(senderAddress)
	if err != nil {
		panic(err)
	}

	if len(bin) == 0 {
		ctx.state = map[string]interface{}{
			senderAddress: map[string]interface{}{
				"balance": strconv.FormatUint(value.Uint64(), 10),
			},
		}
		ctx.stack.Push(*addr)
		return ctx.stack.data
	}

	address := crypto.Keccak256(rlpA)[12:]
	newAddressStr := hex.EncodeToString(address)

	newTx := &Transaction{
		From:  senderAddress,
		To:    newAddressStr,
		Value: value.String(),
	}

	vm := ctx.vm

	newCtx := &executionContext{
		pc:          0,
		memory:      NewMemory(),
		stack:       NewStack(),
		storage:     NewStorage(),
		code:        bin,
		state:       ctx.state,
		transaction: newTx,
	}

	vm.Context = newCtx

	_, returnValue, _ := vm.Execute(bin)
	returnUint256, err := uint256.FromHex(fmt.Sprintf("%v%v", "0x", returnValue))
	if err != nil {
		panic(err)
	}

	ctx.stack.Push(*addr)
	ctx.memory.Set(offset.Uint64(), size.Uint64(), returnUint256.Bytes())
	ctx.state = map[string]interface{}{
		senderAddress: map[string]interface{}{
			"code": map[string]interface{}{
				"bin": returnValue,
			},
		},
	}
	vm.Context = ctx
	fmt.Println(ctx.stack.data)
	return ctx.stack.data
}
