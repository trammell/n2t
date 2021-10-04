package asm

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

// NewProgram is a constructor for Programs
func NewProgram(filename string) (p *Program) {
	p = new(Program)
	p.Filename = filename
	return
}

// read the instructions from the source file
func (p *Program) Read() {
	// Open the source file
	file, err := os.Open(p.Filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// load the file contents
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		txt := CleanUp(scanner.Text())
		if txt != "" {
			p.AppendInstruction(txt)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

// CleanUp strips whitespace and comments from an instruction
func CleanUp(txt string) Instruction {
	txt = regexp.MustCompile(`//.*`).ReplaceAllString(txt, "")
	txt = regexp.MustCompile(`\s`).ReplaceAllString(txt, "")
	return Instruction(txt)
}

func (p *Program) AppendInstruction(i Instruction) {
	if isAInstruction(i) {
		p.Instructions = append(p.Instructions, AInstruction(i))
	} else if isCInstruction(i) {
		p.Instructions = append(p.Instructions, CInstruction(i))
	} else if isLabel(i) {
		p.Instructions = append(p.Instructions, Label(i))
	} else {
		log.Fatal(``, i)
	}
}

// resolve all unresolved symbols in instructions
func (p *Program) ResolveSymbols() {
	addr := Address(0)
	for _, inst := range p.Instructions {
		addr = inst.Resolve(p.SymbolTable, addr)
	}
}

// emit all instructions as machine code
func (p *Program) Emit() {
	for _, i := range p.Instructions {
		code, err := i.Assemble(p.SymbolTable)
		if err != nil {
			log.Fatal(code, err)
		} else {
			fmt.Printf("%016b\n", code)
		}
	}
}
