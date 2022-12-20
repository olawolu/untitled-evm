package evm

import (
	"github.com/holiman/uint256"
)

func memoryMLoad(stack *Stack) (uint64, bool) {
	return calcMemSize64WithUint(stack.Back(0), 32)
}

func memoryMStore8(stack *Stack) (uint64, bool) {
	return calcMemSize64WithUint(stack.Back(0), 1)
}

func memoryMStore(stack *Stack) (uint64, bool) {
	return calcMemSize64WithUint(stack.Back(0), 32)
}

func memorySha3(stack *Stack) (uint64, bool) {
	return calcMemSize64WithUint(stack.Back(0), stack.Back(1).Uint64())
}

func memoryCallDataCopy(stack *Stack) (uint64, bool) {
	return calcMemSize64WithUint(stack.Back(0), stack.Back(2).Uint64())
}

func memoryCodeCopy(stack *Stack) (uint64, bool) {
	return calcMemSize64WithUint(stack.Back(0), stack.Back(2).Uint64())
}

func memoryExtCodeCopy(stack *Stack) (uint64, bool) {
	return calcMemSize64WithUint(stack.Back(1), stack.Back(3).Uint64())
}

func memoryReturn(stack *Stack) (uint64, bool) {
	return calcMemSize64WithUint(stack.Back(0), stack.Back(1).Uint64())
}

func memoryRevert(stack *Stack) (uint64, bool) {
	return calcMemSize64WithUint(stack.Back(0), stack.Back(1).Uint64())
}

func memoryCall(stack *Stack) (uint64, bool) {
	return calcMemSize64WithUint(stack.Back(3), stack.Back(4).Uint64())
}

func memoryCreate(stack *Stack) (uint64, bool) {
	return calcMemSize64WithUint(stack.Back(1), stack.Back(2).Uint64())
}

// calcMemSize64WithUint calculates the required memory size, and returns
// the size and whether the result overflowed uint64
// Identical to calcMemSize64, but length is a uint64
func calcMemSize64WithUint(off *uint256.Int, length64 uint64) (uint64, bool) {
	// if length is zero, memsize is always zero, regardless of offset
	if length64 == 0 {
		return 0, false
	}
	// Check that offset doesn't overflow
	offset64, overflow := off.Uint64WithOverflow()
	if overflow {
		return 0, true
	}
	val := offset64 + length64
	// if value < either of it's parts, then it overflowed
	return val, val < offset64
}
