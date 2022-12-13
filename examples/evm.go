package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"

	"github.com/helicarrierstudio/untitled-evm/evm"
	"github.com/holiman/uint256"
)

type code struct {
	Bin string
	Asm string
}

type expect struct {
	Stack   []string
	Success bool
	Return  string
}

type TestCase struct {
	Name   string
	Code   code
	Expect expect
}

func main() {
	content, err := ioutil.ReadFile("./evm.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	var payload []TestCase
	err = json.Unmarshal(content, &payload)
	if err != nil {
		log.Fatal("Error during json.Unmarshal(): ", err)
	}

	for index, test := range payload {
		fmt.Printf("Test #%v of %v: %v\n", index+1, len(payload), test.Name)

		bin, err := hex.DecodeString(test.Code.Bin)
		if err != nil {
			log.Fatal("Error during hex.DecodeString(): ", err)
		}

		var expectedStack []uint256.Int
		var expectedReturn string
		var expectedSuccess bool
		var in = new(uint256.Int)
		for _, s := range test.Expect.Stack {
			i, ok := new(big.Int).SetString(s, 0)
			if !ok {
				log.Fatal("Error during big.Int.SetString(): ", err)
			}
			in.SetFromBig(i)
			expectedStack = append(expectedStack, *in)
		}

		expectedReturn = test.Expect.Return
		expectedSuccess = test.Expect.Success

		// Note: as the test cases get more complex, you'll need to modify this
		// to pass down more arguments to the evm function and return more than
		// just the stack.
		stack, returnData, success := evm.Run(bin)

		match := len(stack) == len(expectedStack)
		if match {
			for i, s := range stack {
				match = match && (s.Cmp(&expectedStack[i]) == 0)
			}
		}

		if !match {
			fmt.Printf("Instructions: \n%v\n", test.Code.Asm)
			fmt.Printf("Expected: %v\n", toStrings(expectedStack))
			fmt.Printf("Got: %v\n\n", toStrings(stack))
			fmt.Printf("Progress: %v/%v\n\n", index, len(payload))
			log.Fatal("Stack mismatch")
		}

		if returnData != expectedReturn {
			fmt.Printf("Instructions: \n%v\n", test.Code.Asm)
			fmt.Printf("Expected: %v\n", expectedReturn)
			fmt.Printf("Got: %v\n\n", returnData)
			fmt.Printf("Progress: %v/%v\n\n", index, len(payload))
			log.Fatal("Return data mismatch")
		}

		if success != expectedSuccess {
			fmt.Printf("Instructions: \n%v\n", test.Code.Asm)
			fmt.Printf("Expected: %v\n", expectedSuccess)
			fmt.Printf("Got: %v\n\n", success)
			fmt.Printf("Progress: %v/%v\n\n", index, len(payload))
			log.Fatal("Success mismatch")
		}
	}
}

func toStrings(stack []uint256.Int) []string {
	var strings []string
	for _, s := range stack {
		strings = append(strings, s.String())
	}
	return strings
}
