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
func CleanUp(txt string) string {
	txt = regexp.MustCompile(`//.*`).ReplaceAllString(txt, "")
	return regexp.MustCompile(`\s`).ReplaceAllString(txt, "")
}

func (p *Program) AppendInstruction(txt string) {
	if isAInstruction(txt) {
		append(p.Instructions, AInstruction(txt))
	} else if isCInstruction(txt) {
		append(p.Instructions, CInstruction(txt))
	} else if isLabel(txt) {
		append(p.Instructions, Label(txt))
	} else {
		log.Fatal(``, txt)
	}
}

// resolve all unresolved symbols in instructions
func (p *Program) ResolveSymbols() {
	addr := Address(0)
	for _, inst := range p.Instructions {
		addr = inst.ResolveSymbol(p.Symbols, addr)
	}
}

// emit all instructions as machine code
func (p *Program) Emit() {
	for _, i := range p.Instructions {
		out, err := i.Assemble(Symbols)
		if err != nil {
			log.Fatal(err)
		}
		if out != "" {
			fmt.Println(out)
		}
	}
}
