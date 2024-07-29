// Package evm is an **incomplete** implementation of the Ethereum Virtual
// Machine for the "EVM From Scratch" course:
// https://github.com/w1nt3r-eth/evm-from-scratch
//
// To work on EVM From Scratch In Go:
//
// - Install Golang: https://golang.org/doc/install
// - Go to the `go` directory: `cd go`
// - Edit `evm.go` (this file!), see TODO below
// - Run `go test ./...` to run the tests
package evm

import (
	"math/big"
)

// Run runs the EVM code and returns the stack and a success indicator.
func Evm(code []byte) ([]*big.Int, bool) {
	stack := NewStack()
	pc := 0

	stack.Push(big.NewInt(1))
	stack.Push(big.NewInt(2))
	stack.Push(big.NewInt(3))

	stack.Print()

	stack.Pop()
	stack.Pop()
	stack.Print()

	for pc < len(code) {
		op := code[pc]
		pc++

		// TODO: Implement the EVM here!
		_ = op // delete this; it's only here to make the compiler think you're already using `op`
	}

	return stack.stack, true
}
