package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
)

// Initialize the address map for symbols
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

var CComp = map[string]string{
	"0":   "0101010",
	"1":   "0111111",
	"-1":  "1111010",
	"D":   "1001100",
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

// Consruct a regular expression to parse a C-Instruction
func makeCInstructionRegexp() string {

	// helper function to build regexp components
	pipe := func(m map[string]string) string {
		keys := make([]string, 0, len(m))
		for k := range m {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		return "(" + strings.Join(keys, "|") + ")"
	}

	destRe := "(?:([ADM]+)=)?" // the destination part of the regexp
	compRe := pipe(CComp)      // the compute part of the regexp
	jumpRe := pipe(CJump)      // the jump part of the regexp

	return fmt.Sprintf("^%s%s(?:;(%s))?$", destRe, compRe, jumpRe)
}

/***********************************************************************/

type Assembler struct {
	Filename     string
	Instructions []Instruction
	Symbols      map[string]int
}

func NewAssembler(filename string) *Assembler {
	a := new(Assembler)
	a.Filename = filename
	return a
}

// Read the instructions, resolve symbols, and emit assembled code
func (a *Assembler) parse() {
	a.read()
	//a.resolve()
	a.emit()
}

// read the instructions from the source file
func (a *Assembler) read() {
	// load all the instructions
	file, err := os.Open(a.Filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i := NewInstruction(scanner.Text())
		a.Instructions = append(a.Instructions, *i)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

// resolve all unresolved symbols in instructions
func (a *Assembler) resolve() {

}

// emit all instructions as machine code
func (a *Assembler) emit() {
	//for i := range a.Instructions {
	//}
}

/****************************************************/

// instruction types
const (
	A = iota // address instruction
	C = iota // compute instruction
	L = iota // label instruction
)

type Instruction struct {
	Text    string
	Address int
}

func (i *Instruction) Type() (int, error) {
	// is this an A-instruction?
	re, err := regexp.Compile(`^@[0-9a-zA-Z]+$`)
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
		log.Fatal("error compiling regexp")
	}
	if re.MatchString(i.Text) {
		return C, nil
	}

	// error
	return 0, fmt.Errorf("unrecognized instruction type")
}

func NewInstruction(txt string) *Instruction {
	i := new(Instruction)
	i.Text = i.CleanUp(txt)
	return i
}

// Strip whitespace and comments from a line of assembler
func (i *Instruction) CleanUp(txt string) string {
	txt = regexp.MustCompile(`//.*`).ReplaceAllString(txt, "")
	return regexp.MustCompile(`\s`).ReplaceAllString(txt, "")
}

func (i *Instruction) IsEmpty() bool {
	return true
}

func (i *Instruction) Assemble(symbols map[string]int) string {
	ty, err := i.Type()
	if err != nil {
		log.Fatal("error determining instruction type")
	}

	switch ty {
	case A:
	case C:
	case L:
	default:
	}

	return "1111111111111111"
}

/********************************************************************/

// main function: takes a single filename as argument
func main() {
	a := NewAssembler(os.Args[1])
	a.parse()
}
