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

var currentVariableAddress int = FirstVariableAddress

// Return True if this string represents an A instruction
func isAInstruction(txt string) bool {
	return regexp.MustCompile(`^@.+$`).MatchString(txt)
}

// Construct an A Instruction object.
func NewAInstruction(txt string) *AInstruction {
	i := AInstruction(txt)
	return &i
}

// Assemble a single A instruction into binary
func (i *AInstruction) Assemble(symbols map[string]int) (string, error) {

	inst := strings.Trim(string(*i), "@")

	// if it's a number then print it in binary
	if regexp.MustCompile(`^[0-9]+$`).MatchString(inst) {
		num, err := strconv.Atoi(inst)
		if err != nil {
			return ``, fmt.Errorf("unable to assemble A instruction: %v", i)
		}
		return fmt.Sprintf("0%015b", num), nil
	}

	// If the symbol doea not resolve, then claim another variable slot.
	addr, ok := symbols[inst]
	if !ok {
		symbols[inst] = currentVariableAddress
		addr = currentVariableAddress
		currentVariableAddress++
	}
	return fmt.Sprintf("0%015b", addr), nil
}
