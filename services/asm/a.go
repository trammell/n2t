package asm

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/rs/zerolog/log"
)

// Return True if this string represents an A instruction
func IsAInstruction(i Instruction) bool {
	return regexp.MustCompile(`^@.+$`).MatchString(string(i))
}

// Assemble a single A instruction into binary. A instructions can be either
// explicit or implicit (symbolic):
//   * Explicit (e.g. `@100`) instructions have a constant value
//   * Implicit (e.g. `@abc`) instructions needs to be managed in the symbol table.
func (i AInstruction) Assemble(st SymbolTable) (SymbolTable, []MachineCode, error) {

	inst := strings.Trim(string(i), "@")
	m := log.Debug().Str("A", string(inst))

	// if the instruction is a number then it assembles trivially to an integer
	if regexp.MustCompile(`^[0-9]+$`).MatchString(inst) {
		m.Send()
		num, err := strconv.Atoi(inst)
		if err != nil {
			return st, []MachineCode{}, fmt.Errorf("unable to assemble A instruction: %v", i)
		}
		return st, []MachineCode{MachineCode(num)}, nil
	}

	// If it's a symbol that already exists in the symbol table, return it.
	addr, exists := st.Table[Symbol(inst)]
	if exists {
		m.Uint16("addr", uint16(addr)).Send()
		return st, []MachineCode{MachineCode(addr)}, nil
	}

	// If we get to this point we have a symbolic A-instruction that does
	// not exist (yet) in the symbol table. Add it, and increment the
	// next-free address pointer.
	addr = st.Pointer
	st.Table[Symbol(inst)] = st.Pointer
	st.Pointer += 1

	m.Uint16("new addr", uint16(addr)).Send()
	return st, []MachineCode{MachineCode(addr)}, nil
}

// A instructions don't update the symbol table, but they do...
func (i AInstruction) UpdateSymbolTable(st SymbolTable, a Address) (SymbolTable, Address) {
	return st, a + 1
}
