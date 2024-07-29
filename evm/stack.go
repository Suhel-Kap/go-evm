package evm

import (
	"fmt"
	"math/big"
)

type Stack struct {
	stack []*big.Int
}

// NewStack creates a new stack with the given maximum size.
func NewStack() *Stack {
	return &Stack{
		stack: make([]*big.Int, 0, MAXIMUM_STACK_SIZE),
	}
}

// Push pushes a value onto the stack.
func (stack *Stack) Push(value *big.Int) {
	stack.stack = append(stack.stack, value)
}

// Pop pops a value from the stack.
func (s *Stack) Pop() *big.Int {
	if len(s.stack) == 0 {
		return nil
	}

	value := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return value
}

func (s *Stack) Print() {
	for i := len(s.stack) - 1; i >= 0; i-- {
		if i == len(s.stack)-1 {
			fmt.Printf("%d <- top\n", s.stack[i])
		} else if i == 0 {
			fmt.Printf("%d <- bottom\n", s.stack[i])
		} else {
			fmt.Printf("%d\n", s.stack[i])
		}
	}
}
