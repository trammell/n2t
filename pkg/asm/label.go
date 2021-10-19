package asm

import (
	"regexp"
	"strings"

	"github.com/rs/zerolog/log"
)

// Return True if this looks like a label
func isLabel(i Instruction) bool {
	return regexp.MustCompile(`^\(.+\)$`).MatchString(string(i))
}

// Labels are pseudoinstructions, so they don't really assemble.
func (i Label) Assemble(s SymbolTable) ([]Code, error) {
	return []Code{}, nil // return an empty array, but successful!
}

// Labels do affect the symbol table though. Save the current address
// in the symbol table under this instruction label.
func (i Label) Resolve(st SymbolTable, addr Address) Address {
	sym := Symbol(strings.Trim(string(i), "()")) // strip parens
	log.Info().Str("sym", string(sym)).Uint16("addr", uint16(addr))
	st[sym] = addr
	return addr
}
