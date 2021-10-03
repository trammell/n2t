package asm

import (
	"regexp"
)

// Return True if this looks like a label
func isLabel(txt string) bool {
	return regexp.MustCompile(`^\(.+\)$`).MatchString(txt)
}

func (i Label) Assemble(s SymbolTable) (Code, error) {
	return 0, nil
}

func (i Label) ResolveSymbol(s SymbolTable, a Address) Address {

	return 0
}
