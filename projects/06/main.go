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

type asm struct {
	Filename string
	Insts    []inst
	Sym      map[string]int
}

// Read the instructions, resolve symbols, and emit assembled code
func (a *asm) parse() {
	a.read()
	a.resolve()
	a.emit()
}

// read the instructions from the source file
func (a *asm) read() {
	// load all the instructions
	file, err := os.Open(a.Filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i := NewInstruction(scanner.Text())
		a.Insts = append(a.Insts, *i)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

// resolve all unresolved symbols in instructions
func (a *asm) resolve() {

}

// emit all instructions as machine code
func (a *asm) emit() {

}

/****************************************************/

type inst struct {
	Txt  string
	Addr int
}

func NewInstruction(txt string) *inst {
	i := new(inst)
	i.Txt = txt
	i.Type = InstructionType(txt)
	return i
}



// Strip whitespace and comments from a line of assembler
func (i *inst) Canonicalize(txt string) string {
	txt = regexp.MustCompile(`//.*`).ReplaceAllString(txt, "")
	return regexp.MustCompile(`\s`).ReplaceAllString(txt, "")
}

func (i *inst) IsEmpty() bool {
	return true
}

func (i *inst) IsAInstruction() bool {
	return regexp.MustCompile(`^@[0-9a-zA-Z]+$`).MatchString(i.Txt)
}

// A C-instruction looks like `destination=compute;jump`.
func (i *inst) IsCInstruction() bool {
	rex := makeCInstructionRegexp()
	//fmt.Println(rex)
	return regexp.MustCompile(rex).MatchString(i.Txt)
}

func (i *inst) IsLabel() bool {
	return regexp.MustCompile(`^\(.*\)$`).MatchString(i.Txt)
}

func (i *inst) Assemble(sym map[string]int) string {
	return "1111111111111111"
}

/********************************************************************/

// main function: takes a single filename as argument
func main() {
	a := asm{Filename: os.Args[1]}
	a.parse()
}
