// vim: set ts=4 :

package main

import (
	"fmt"
	"maps"
	"sort"
)

// These are the symbols defined in the Hack spec and their values.
var PredefinedSymbols = map[string]int{
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
func NewSymbolTable() *SymbolTable {
	st := new(SymbolTable)
	st.table = maps.Clone(PredefinedSymbols)
	return st
}

func (st SymbolTable) contains(sym string) bool {
	_, ok := st.table[sym]
	return ok
}

func (st *SymbolTable) addEntry(sym string, address int) {
	st.table[sym] = address
}

func (st SymbolTable) getAddress(sym string) int {
	return st.table[sym]
}

func (st SymbolTable) String() string {
	symbols := make([]string, 0)
	maxlen := 0
	for s, _ := range st.table {
		symbols = append(symbols, s)
		if len(s) > maxlen {
			maxlen = len(s)
		}
	}
	mysort := func(i, j int) bool {
		return st.table[symbols[i]] < st.table[symbols[j]]
	}
	sort.Slice(symbols, mysort)
	out := ""
	format := fmt.Sprintf("st[%%%ds] = %%5d\n", maxlen)
	for _, s := range symbols {
		out += fmt.Sprintf(format, s, st.table[s])
	}
	return out
}
