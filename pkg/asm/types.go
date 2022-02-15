package asm

import (
	"bufio"
)

// base types
type Address uint16
type MachineCode uint16
type Instruction string
type Symbol string
type SymbolTable map[Symbol]Address

// Hack _address_ ("A") instructions look like `@value`, and are assembled to
// binary values that look like `0bbb bbbb bbbb bbbb`
type AInstruction Instruction

// Hack _compute_ ("C") instructions look like `dest=comp;jump`, where <comp>
// is the value computed, <dest> is the storage destination, and <jump> is the
// jump condition.
type CInstruction Instruction

// Hack _labels_ are placeholder instructions that make human-readable memory
// addresses. Labels don't get assembled into machine instructions.
type Label Instruction

// Assemblable instruction objects types need to do two things:
// 1. assemble into a slice of type MachineCode
// 2. update the symbol table
type Assemblable interface {
	Assemble(SymbolTable) ([]MachineCode, error)
	UpdateSymbolTable(SymbolTable, Address) Address
}

// The _sourcecode_ object represents the input code to be compiled.
type SourceCode struct {
	FileName string
	Scanner  *bufio.Scanner
}

// The Hack _assembler_ object does the work of converting the Hack source
// file into binary. Along the way the assembler creates a symbol table.
type Assembler struct {
	SourceCode   SourceCode
	SymbolTable  SymbolTable
	Instructions []Assemblable
	MachineCode  chan MachineCode
}
