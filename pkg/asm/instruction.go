package asm

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

const (
	a = iota // address instruction
	b = iota // label instruction
	c = iota // compute instruction
	u = iota // unrecognized or invalid instruction
)

// the first 16 memory addresses are reserved
const StartingAddress = 16
var nextVariableAddress int = StartingAddress

// Instruction represents a single instruction in a program
type Instruction struct {
	Text    string
	Address int
}

// NewInstruction is an Instruction constructor
func NewInstruction(txt string) *Instruction {
	i := new(Instruction)
	i.Text = i.CleanUp(txt)
	return i
}

// CleanUp strips whitespace and comments from an instruction
func (i *Instruction) CleanUp(txt string) string {
	txt = regexp.MustCompile(`//.*`).ReplaceAllString(txt, "")
	return regexp.MustCompile(`\s`).ReplaceAllString(txt, "")
}

// GetType determines the type (A-instruction, C-instruction, Label) of an instruction
func (i *Instruction) GetType() (int, error) {
	// is this an A-instruction?
	re, err := regexp.Compile(`^@.+$`)
	if err != nil {
		log.Fatal("error compiling regexp")
	}
	if re.MatchString(i.Text) {
		return a, nil
	}

	// is this a label?
	re, err = regexp.Compile(`^\(.*\)$`)
	if err != nil {
		log.Fatal("error compiling regexp")
	}
	if re.MatchString(i.Text) {
		return b, nil
	}

	// is this a C-instruction?
	crex := makeCInstructionRegexp()
	re, err = regexp.Compile(crex)
	if err != nil {
		log.Fatal("error compiling C instruction regexp")
	}
	if re.MatchString(i.Text) {
		return c, nil
	}

	// error
	return -1, fmt.Errorf("unrecognized instruction type")
}


// Assemble assembles a single instruction from text into binary
func (i *Instruction) Assemble(symbols map[string]int) (string, error) {
	itype, err := i.GetType()
	if err != nil {
		log.Fatalf("error determining instruction type: %v", i)
	}

	switch itype {
	case a:
		return i.AssembleAInstruction(symbols)
	case c:
		return i.AssembleCInstruction()
	case b:
		/* label; do nothing */
	default:
		log.Fatalf("unable to assemble instruction: %v", i)
	}

	return "", nil
}


// AssembleAInstruction converts an A instruction to binary
func (i *Instruction) AssembleAInstruction(symbols map[string]int) (string, error) {

	inst := strings.Trim(i.Text, "@") // strip leading @ from instruction

	// if it's a number then print it in binary, otherwise look it up in the symbol table
	isNumeric := regexp.MustCompile(`^[0-9]+$`)
	if isNumeric.MatchString(inst) {
		num, err := strconv.Atoi(inst)
		if err != nil {
			return ``, fmt.Errorf("unable to assemble A instruction: %v", i)
		}
		return fmt.Sprintf("0%015b", num), nil
	}
	addr, ok := symbols[inst]
	if !ok {
		// if the symbol doesn't resolve at this point, it's a variable
		addr = nextVariableAddress
		symbols[inst] = nextVariableAddress
		nextVariableAddress++
		//return ``, fmt.Errorf("unable to resolve symbol: %v (%v)", inst, i)
	}
	return fmt.Sprintf("0%015b", addr), nil
}

// SplitCInstruction is a helper function to split up a C instructions into parts
func (i *Instruction) SplitCInstruction() (string, string, string, error) {

	// compile the regex
	crex := makeCInstructionRegexp()
	r, err := regexp.Compile(crex)
	if err != nil {
		log.Fatal("error compiling regexp")
	}

	match := r.FindStringSubmatch(i.Text)
	if len(match) > 0 {
		return match[1], match[2], match[3], nil
	}
	return ``, ``, ``, fmt.Errorf("error splitting C instruction: %v", i.Text)
}

// AssembleCInstruction converts a C instruction into binary
func (i *Instruction) AssembleCInstruction() (string, error) {
	// extract dest, comp, and jump expressions from C instruction with regexp
	dest, comp, jump, err := i.SplitCInstruction()
	if err != nil {
		log.Fatalf("error splitting instruction: %v", err)
	}

	// calculate dest bits: M=1, D=2, A=4
	destbits := 0
	if strings.Contains(dest, "M") {
		destbits = destbits | 1
	}
	if strings.Contains(dest, "D") {
		destbits = destbits | 2
	}
	if strings.Contains(dest, "A") {
		destbits = destbits | 4
	}

	// calculate comp bits
	compbits, ok := CComp[comp]
	if !ok {
		log.Fatalf("error finding comp bits for %v", comp)
	}

	// calculate jump bits
	jumpbits, ok := CJump[jump]
	if !ok {
		log.Fatalf("error finding jump bits for %v", jump)
	}

	return fmt.Sprintf("111%s%03b%s", compbits, destbits, jumpbits), nil
}
