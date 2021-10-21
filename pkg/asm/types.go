package asm

// base types
type Address uint16
type Code uint16
type Instruction string
type Symbol string
type SymbolTable map[Symbol]Address

// Address ("A") instructions look like `@value`, and are assembled to binary
// values that look like `0bbb bbbb bbbb bbbb`
type AInstruction Instruction

// Compute ("C") instructions look like `dest=comp;jump`, where <comp> is the
// value computed, <dest> is the storage destination, and <jump> is the jump
// condition.
type CInstruction Instruction

// Labels are placeholder instructions that make human-readable memory
// addresses. Labels don't get assembled into machine instructions.
type Label Instruction

// Assembler types need to be able to do two things:
// 1. return 0 or 1 binary codes
// 2. add an address to the symbol table
type Assembler interface {
	Assemble(SymbolTable) ([]Code, error)
	Resolve(SymbolTable, Address) Address
}

// Program represents the program to be compiled.
type Program struct {
	Filename     string
	Instructions []Assembler
	SymbolTable  SymbolTable
}
