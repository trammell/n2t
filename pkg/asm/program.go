package asm

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Program represents the program to be compiled.
type Program struct {
	Filename     string
	Instructions []InstructionAssembler
	Symbols      map[string]int
}

// NewProgram is a constructor for Programs
func NewProgram(filename string) (p *Program) {
	p = new(Program)
	p.Filename = filename
	return
}

// read the instructions from the source file
func (p *Program) Read() {
	// load all the instructions
	file, err := os.Open(p.Filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i := NewInstruction(scanner.Text())
		p.AppendInstruction(i)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func (p *Program) AppendInstruction(i InstructionAssembler) {
	if string(i) != "" {
		append(p.Instructions,i)
	}
}

// resolve all unresolved symbols in instructions
func (p *Program) Resolve() {
	address := 0
	for _, inst := range p.Instructions {
		itype, err := inst.GetType()
		if err != nil {
			log.Fatal(err)
		}
		if itype == b {
			// extrsact the label from the instruction and assign the next address
			label := strings.Trim(inst.Text, "()")
			//fmt.Fprintln(os.Stderr, "Symbol[", label, "] = ", address)
			Symbols[label] = address
		} else {
			address++
		}
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
