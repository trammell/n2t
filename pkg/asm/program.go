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
	p.SymbolTable = PredefinedSymbols.clone()
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
		// A and C instructions increment address, labels do not
		addr = inst.Resolve(p.SymbolTable, addr)
	}
}

// emit all instructions, as machine code, to STDOUT
func (p *Program) EmitToStdout() {
	gen := p.Emit()
	for s, x := gen(); !x; s, x = gen() {
		fmt.Println(s)
	}
}

// emit all instructions, as machine code, in a slice of strings
func (p *Program) EmitToSlice() []string {
	gen := p.Emit()
	var out []string
	for s, x := gen(); !x; s, x = gen() {
		out = append(out, s)
	}
	return out
}

// emit all instructions as machine code
func (p *Program) Emit() func() (string, bool) {

	// make a copy of p.Instructions, and then just keep shifting them off
	instructions := make([]Assembler, len(p.Instructions))
	copy(instructions, p.Instructions)

	// return a closure that shifts off & assembles the first slice elt
	return func() (string, bool) {
		if len(instructions) == 0 {
			return ``, true // no value, and exit
		}

		i := instructions[0]
		instructions = instructions[1:]
		code, err := i.Assemble(p.SymbolTable)
		if err != nil {
			log.Fatal(code[0], err)
		}
		if len(code) > 0 {
			return fmt.Sprintf("%016b", code[0]), false
		}
		return ``, false // no value, but continue
	}
}
