package asm

import (
	"regexp"
	"strings"

	"github.com/rs/zerolog/log"
)

// Return True if this looks like a label
func IsLabel(i Instruction) bool {
	return regexp.MustCompile(`^\(.+\)$`).MatchString(string(i))
}

// implement the stringer interface
func (x Label) String() string {
	return string(x)
}

// Labels are pseudoinstructions, so they don't really assemble.
func (i Label) Assemble(s SymbolTable) (SymbolTable, []MachineCode, error) {
	return s, []MachineCode{}, nil // return an empty array, but successful!
}

// Labels do affect the symbol table though. Save the current address
// in the symbol table under this instruction label.
func (x Label) UpdateSymbolTable(st SymbolTable, addr Address) (SymbolTable, Address) {
	sym := Symbol(strings.Trim(x.String(), "()")) // strip parens
	log.Info().Str("sym", string(sym)).Uint16("addr", uint16(addr))
	st.Table[sym] = addr
	return st, addr
}
