package asm

import "regexp"

// Instruction represents a single instruction in a program
type InstructionAssembler interface {
	Assemble()
}
type Instruction string
type AInstruction Instruction
type CInstruction Instruction

// Program represents the program to be compiled.
type Program struct {
	Filename     string
	Instructions []InstructionAssembler
	Symbols      map[string]int
}

// CleanUp strips whitespace and comments from an instruction
func CleanUp(txt string) string {
	txt = regexp.MustCompile(`//.*`).ReplaceAllString(txt, "")
	return regexp.MustCompile(`\s`).ReplaceAllString(txt, "")
}
