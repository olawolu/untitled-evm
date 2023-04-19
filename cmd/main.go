package main

import (
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/helicarrierstudio/untitled-evm/evm"
)

func main() {
	codeFile := flag.String("code", "./config/code.json", "code file")
	blockFile := flag.String("block", "./config/block.json", "block file")
	transactionFile := flag.String("transaction", "./config/tx.json", "transaction file")
	stateFile := flag.String("state", "./config/state.json", "state file")

	flag.Parse()
	code := parseCodeFile(*codeFile)
	block := parseBlockFile(*blockFile)
	transaction := parseTransactionFile(*transactionFile)
	state := parseStateFile(*stateFile)

	stack, result, status := evm.Run(code, &transaction, state, &block)

	fmt.Println("--------stack--------")
	for _, item := range stack {
		fmt.Printf("[%v]\n", item.Hex())
	}
	b, err := hex.DecodeString(stack[0].String()[2:])
	if err != nil {
		log.Fatal("Error during hex.DecodeString(): ", err)
	}
	fmt.Println(string(b))

	fmt.Println("result:", result)
	fmt.Println("status:", status)
}

type code struct {
	Bin string
	Asm string
}

func parseCodeFile(file string) []byte {
	args := os.Args
	if len(args) > 1 {
		if os.Args[1] == "code" {
			bin, err := hex.DecodeString(os.Args[2])
			if err != nil {
				log.Fatal("Error during hex.DecodeString(): ", err)
			}
			return bin
		}
	}

	content, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}
	var payload code
	err = json.Unmarshal(content, &payload)
	if err != nil {
		log.Fatal("Error during json.Unmarshal(): ", err)
	}

	bin, err := hex.DecodeString(payload.Bin)
	if err != nil {
		log.Fatal("Error during hex.DecodeString(): ", err)
	}

	return bin
}

func parseBlockFile(file string) evm.Block {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	var payload evm.Block
	err = json.Unmarshal(content, &payload)
	if err != nil {
		log.Fatal("Error during json.Unmarshal(): ", err)
	}
	return payload
}

func parseTransactionFile(file string) evm.Transaction {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	var payload evm.Transaction
	err = json.Unmarshal(content, &payload)
	if err != nil {
		log.Fatal("Error during json.Unmarshal(): ", err)
	}
	return payload
}

func parseStateFile(file string) interface{} {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	var payload interface{}
	err = json.Unmarshal(content, &payload)
	if err != nil {
		log.Fatal("Error during json.Unmarshal(): ", err)
	}
	return payload
}
