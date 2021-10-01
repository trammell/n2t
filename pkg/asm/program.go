package asm

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
)

// Consruct a regular expression to parse a C-Instruction
func makeCInstructionRegexp() string {

	// "pipe" is a helper function to build regexp components.
	// One tricky thing here is that keys need to be
	pipe := func(m map[string]string, cap string) string {
		keys := make([]string, 0, len(m))
		for k := range m {
			keys = append(keys, regexp.QuoteMeta(k))
		}
		sort.Strings(keys)
		return "(" + strings.Join(keys, "|") + ")"
	}

	destRe := "(?:([ADM]+)=)?"    // the destination part of the regexp
	compRe := pipe(CComp, `comp`) // the compute part of the regexp
	jumpRe := pipe(CJump, `jump`) // the jump part of the regexp
	regexp := fmt.Sprintf("^%s%s(?:;%s)?$", destRe, compRe, jumpRe)
	return regexp
}

/***********************************************************************/

// Program represents the program to be compiled.
type Program struct {
	Filename     string
	Instructions []Instruction
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
		if i.Text != "" {
			p.Instructions = append(p.Instructions, *i)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
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
