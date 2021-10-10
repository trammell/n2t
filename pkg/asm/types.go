package asm

// base types
type Address uint
type Code uint16
type Instruction string
type Symbol string
type SymbolTable map[Symbol]Address

// Address ("A") instructions look like `@value`, and are assembled to binary
// values that look like `0bbb bbbb bbbb bbbb`
type AInstruction Instruction

// C instructions look like `dest=comp;jump`, where <dest> is the storage destination
// of the computed value, <comp>
type CInstruction Instruction

// Labels are not really full-fledged instructions as they don't get assembled to a code;
// they're more of a helper to building sane jump commands.
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
