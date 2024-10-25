package asm

import (
	"bufio"
)

// base types
type Address uint16
type MachineCode uint16
type Instruction string
type Symbol string

// Hack address ("A") instructions look like `@value`, and are assembled to
// binary values that look like `0bbb bbbb bbbb bbbb`
type AInstruction Instruction

// Hack compute ("C") instructions look like `dest=comp;jump`, where <comp>
// is the value computed, <dest> is the storage destination, and <jump> is the
// jump condition.
type CInstruction Instruction

// Hack labels are placeholder instructions that generate human-readable memory
// addresses. Labels don't get assembled into machine instructions, but they do
// create symbol table entries.
type Label Instruction

// The symbol table maps labels to addresses in both RAM and ROM
// FIXME: is there a chance of collision? should there be two tables?
type SymbolTable struct {
	Table   map[Symbol]Address
	Pointer Address
}

// Assemblable instruction objects types need to do two things:
// 1. assemble into a slice of type MachineCode
// 2. update the symbol table
type Assemblable interface {
	Assemble(SymbolTable) (SymbolTable, []MachineCode, error)
	UpdateSymbolTable(SymbolTable, Address) (SymbolTable, Address)
}

// The _sourcecode_ object represents the input code to be compiled.
type Source struct {
	FileName string
	Scanner  *bufio.Scanner
}

// The Hack _assembler_ object does the work of converting the Hack source
// file into binary. Along the way the assembler creates a symbol table.
type Assembler struct {
	Source       Source
	SymbolTable  SymbolTable
	Instructions []Assemblable
	MachineCode  []MachineCode
}
