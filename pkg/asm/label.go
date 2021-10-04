package asm

import (
	"regexp"
)

// Return True if this looks like a label
func isLabel(i Instruction) bool {
	return regexp.MustCompile(`^\(.+\)$`).MatchString(string(i))
}

func (i Label) Assemble(s SymbolTable) (Code, error) {
	return 0, nil
}

func (i Label) Resolve(s SymbolTable, a Address) Address {

	return 0
}
