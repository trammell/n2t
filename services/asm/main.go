package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

// Symbols maps text symbols to memory addresses.
var Symbols = map[string]int{
	"R0":     0,
	"R1":     1,
	"R2":     2,
	"R3":     3,
	"R4":     4,
	"R5":     5,
	"R6":     6,
	"R7":     7,
	"R8":     8,
	"R9":     9,
	"R10":    10,
	"R11":    11,
	"R12":    12,
	"R13":    13,
	"R14":    14,
	"R15":    15,
	"SCREEN": 16384,
	"KBD":    24576,
	"SP":     0,
	"LCL":    1,
	"ARG":    2,
	"THIS":   3,
	"THAT":   4,
}

// CComp lists the possible C-instruction computations.
var CComp = map[string]string{
	"0":   "0101010",
	"1":   "0111111",
	"-1":  "0111010",
	"D":   "0001100",
	"A":   "0110000",
	"!D":  "0001101",
	"!A":  "0110001",
	"-D":  "0001111",
	"-A":  "0110011",
	"D+1": "0011111",
	"A+1": "0110111",
	"D-1": "0001110",
	"A-1": "0110010",
	"D+A": "0000010",
	"D-A": "0010011",
	"A-D": "0000111",
	"D&A": "0000000",
	"D|A": "0010101",
	"M":   "1110000",
	"!M":  "1110001",
	"-M":  "1110011",
	"M+1": "1110111",
	"M-1": "1110010",
	"D+M": "1000010",
	"D-M": "1010011",
	"M-D": "1000111",
	"D&M": "1000000",
	"D|M": "1010101",
}

// CJump lists the C-instruction jump encodings
var CJump = map[string]string{
	"":    "000",
	"JGT": "001",
	"JEQ": "010",
	"JGE": "011",
	"JLT": "100",
	"JNE": "101",
	"JLE": "110",
	"JMP": "111",
}

// main function: takes a single filename as argument
// Read the instructions, resolve symbols, and emit assembled code
func main() {
	f := NewAssemblyFile(os.Args[1])
	a := NewAssembler()
	a.Assemble(f)
	a.Emit()
}

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
func (p *Program) read() {
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
func (p *Program) resolve() {
	address := 0
	for _, inst := range p.Instructions {
		itype, err := inst.GetType()
		if err != nil {
			log.Fatal(err)
		}
		if itype == L {
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
func (p *Program) emit() {
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

/****************************************************/

// instruction types
const (
	A = iota // address instruction
	C = iota // compute instruction
	L = iota // label instruction
)

// Instruction represents a single instruction in a program
type Instruction struct {
	Text    string
	Address int
}

// GetType determines the type (A-instruction, C-instruction, Label) of an instruction
func (i *Instruction) GetType() (int, error) {
	// is this an A-instruction?
	re, err := regexp.Compile(`^@.+$`)
	if err != nil {
		log.Fatal("error compiling regexp")
	}
	if re.MatchString(i.Text) {
		return A, nil
	}

	// is this a label?
	re, err = regexp.Compile(`^\(.*\)$`)
	if err != nil {
		log.Fatal("error compiling regexp")
	}
	if re.MatchString(i.Text) {
		return L, nil
	}

	// is this a C-instruction?
	crex := makeCInstructionRegexp()
	re, err = regexp.Compile(crex)
	if err != nil {
		log.Fatal("error compiling C instruction regexp")
	}
	if re.MatchString(i.Text) {
		return C, nil
	}

	// error
	return -1, fmt.Errorf("unrecognized instruction type")
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

// Assemble assembles a single instruction from text into binary
func (i *Instruction) Assemble(symbols map[string]int) (string, error) {
	itype, err := i.GetType()
	if err != nil {
		log.Fatalf("error determining instruction type: %v", i)
	}

	switch itype {
	case A:
		return i.AssembleAInstruction(symbols)
	case C:
		return i.AssembleCInstruction()
	case L:
		/* do nothing */
	default:
		log.Fatalf("unable to assemble instruction: %v", i)
	}

	return "", nil
}

var nextVariableAddress int = 16

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
	return ``, ``, ``, fmt.Errorf("Error splitting C instruction: %v", i.Text)
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

/********************************************************************/

