package asm

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// the first 16 memory addresses are reserved; global currentVariableAddress
// is used to keep track of where we are.
const FirstVariableAddress = 16

var currentVariableAddress Address = FirstVariableAddress

// Return True if this string represents an A instruction
func isAInstruction(i Instruction) bool {
	return regexp.MustCompile(`^@.+$`).MatchString(string(i))
}

// Assemble a single A instruction into binary
func (i AInstruction) Assemble(st SymbolTable) ([]Code, error) {

	inst := strings.Trim(string(i), "@")

	// if it's a number then print it in binary
	if regexp.MustCompile(`^[0-9]+$`).MatchString(inst) {
		num, err := strconv.Atoi(inst)
		if err != nil {
			return []Code{}, fmt.Errorf("unable to assemble A instruction: %v", i)
		}
		return []Code{Code(num)}, nil
	}

	// If the symbol doea not resolve, then claim another variable slot.
	addr, ok := st[Symbol(inst)]
	if !ok {
		st[Symbol(inst)] = currentVariableAddress
		addr = currentVariableAddress
		currentVariableAddress++
	}
	return []Code{Code(addr)}, nil
}

// FIXME
func (i AInstruction) Resolve(s SymbolTable, a Address) Address {
	return a + 1
}
