# untiltled-evm

Created to understand how the EVM works.

The primary goal of this project is to able to execute a compiled bytecode of simple stack operations (Ethereum smart contracts are compiled to bytecode which the EVM executes).

To achieve this, some of the opcodes defined in the Ethereum [yellow paper](https://ethereum.github.io/yellowpaper/paper.pdf) will be  implemented.

## Resources

- [Ethereum Yellow Paper](https://ethereum.github.io/yellowpaper/paper.pdf)
- [Playground](https://www.evm.codes/playground)

## How to run

```bash
git clone
cd untiltled-evm
cd examples
go run evm.go
```
