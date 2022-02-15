package asm

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/rs/zerolog/log"
)

// the first 16 memory addresses are reserved; global currentVariableAddress
// is used to keep track of where we are.
const FirstVariableAddress = 16

var currentVariableAddress Address = FirstVariableAddress

// Return True if this string represents an A instruction
func IsAInstruction(i Instruction) bool {
	return regexp.MustCompile(`^@.+$`).MatchString(string(i))
}

// Assemble a single A instruction into binary
func (i AInstruction) Assemble(st SymbolTable) ([]MachineCode, error) {

	inst := strings.Trim(string(i), "@")
	m := log.Info().Str("A", string(inst))

	// if it's a number then print it in binary
	if regexp.MustCompile(`^[0-9]+$`).MatchString(inst) {
		m.Send()
		num, err := strconv.Atoi(inst)
		if err != nil {
			return []MachineCode{}, fmt.Errorf("unable to assemble A instruction: %v", i)
		}
		return []MachineCode{MachineCode(num)}, nil
	}

	// If the symbol doea not resolve, then claim another variable slot.
	addr, exists := st[Symbol(inst)]
	if exists {
		m.Uint16("addr", uint16(addr)).Send()
	} else {
		st[Symbol(inst)] = currentVariableAddress
		addr = currentVariableAddress
		m.Uint16("new addr", uint16(addr)).Send()
		currentVariableAddress++
	}
	return []MachineCode{MachineCode(addr)}, nil
}

// FIXME
func (i AInstruction) Resolve(s SymbolTable, a Address) Address {
	return a + 1
}
