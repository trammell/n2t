package asm

import (
	"fmt"
	"sort"
)

// These are the symbols defined in the Hack spec and their values.
var PredefinedSymbols = map[Symbol]Address{
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
func DefaultSymbolTable() SymbolTable {
	return SymbolTable{
		Pointer: FirstVariableAddress,
		Table:   PredefinedSymbols,
	}
}

// Clone a symbol table. This is handy for testing.
func (st SymbolTable) Clone() SymbolTable {
	t := make(map[Symbol]Address)
	for k, v := range st.Table {
		t[k] = v
	}
	return SymbolTable{Pointer: st.Pointer, Table: t}
}

// implement the Stringer interface
func (st SymbolTable) String() (out string) {

	// get a sorted slice of symbol table keys
	keys := make([]string, 0, len(st.Table))
	for k := range st.Table {
		keys = append(keys, string(k))
	}
	sort.Strings(keys)

	// show the pointer first
	out = fmt.Sprintf("pointer=%v\n", st.Pointer)

	// append the other symbol table members
	for _, k := range keys {
		out += fmt.Sprintf("st[%v]=%v\n", k, st.Table[Symbol(k)])
	}

	return out
}
