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
// explicit (e.g. `@100`), which is simply a constant value, or implicit
// (e.g. `@abc`), in which the value needs to be found in the symbol table.
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

	// If the symbol does not resolve, then claim another variable slot.
	addr, exists := st.Table[Symbol(inst)]
	if exists {
		m.Uint16("addr", uint16(addr)).Send()
	} else {
		st.Table[Symbol(inst)] = st.Pointer
		addr = st.Pointer
		m.Uint16("new addr", uint16(addr)).Send()
		st.Pointer++
	}
	return []MachineCode{MachineCode(addr)}, nil
}

// A instructions don't update the symbol table, but they do...
func (i AInstruction) UpdateSymbolTable(s SymbolTable, a Address) (nextaddr Address) {
	nextaddr = a + 1
	return
}
