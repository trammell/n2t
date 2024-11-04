package main

import (
	"maps"
)

// These are the symbols defined in the Hack spec and their values.
var PredefinedSymbols = SymbolTable{
	"R0":     0,
	"R1":     1,
	"R2":     2,
	"R3":     3,
	"R4":     4,
	"R5":     5,
	"R6":     6,
	"R7":     7,
	"R8":     8,
	"R9":     9,
	"R10":    10,
	"R11":    11,
	"R12":    12,
	"R13":    13,
	"R14":    14,
	"R15":    15,
	"SCREEN": 16384,
	"KBD":    24576,
	"SP":     0,
	"LCL":    1,
	"ARG":    2,
	"THIS":   3,
	"THAT":   4,
}

// The first open variable slot is at this RAM address.
const FirstVariableAddress = 16

// Construct a "default" symbol table containing the Hack standard symbols.
// Make it a clone for isolation during testing.
func NewSymbolTable() SymbolTable {
	return maps.Clone(PredefinedSymbols)
}

func (st SymbolTable) contains(sym string) bool {
	return false
}

func (st *SymbolTable) addEntry(sym string, address int) {


}

func (st SymbolTable) getAddress(sym string) int {
	return 1

}
