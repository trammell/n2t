package asm

import (
	"fmt"
	"log"
	"regexp"
	"sort"
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
var CComp = map[string]int{
	"0":   0b0101010,
	"1":   0b0111111,
	"-1":  0b0111010,
	"D":   0b0001100,
	"A":   0b0110000,
	"!D":  0b0001101,
	"!A":  0b0110001,
	"-D":  0b0001111,
	"-A":  0b0110011,
	"D+1": 0b0011111,
	"A+1": 0b0110111,
	"D-1": 0b0001110,
	"A-1": 0b0110010,
	"D+A": 0b0000010,
	"D-A": 0b0010011,
	"A-D": 0b0000111,
	"D&A": 0b0000000,
	"D|A": 0b0010101,
	"M":   0b1110000,
	"!M":  0b1110001,
	"-M":  0b1110011,
	"M+1": 0b1110111,
	"M-1": 0b1110010,
	"D+M": 0b1000010,
	"D-M": 0b1010011,
	"M-D": 0b1000111,
	"D&M": 0b1000000,
	"D|M": 0b1010101,
}

// CJump lists the C-instruction jump encodings
var CJump = map[string]int{
	"":    0, // 000
	"JGT": 1, // 001
	"JEQ": 2, // 010
	"JGE": 3, // 011
	"JLT": 4, // 100
	"JNE": 5, // 101
	"JLE": 6, // 110
	"JMP": 7, // 111
}

// pipe() is a helper function for building the C instruction regular expression.
// We're looking to build a big "OR" regular expression like "(a|b|c|d|...)". Some
// of the keys contain metacharacters like "|", "+" and "-" so we quotemeta() them.
func pipe(m map[string]int) string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, regexp.QuoteMeta(k))
	}
	sort.Strings(keys)
	return "(" + strings.Join(keys, "|") + ")"
}

// Consruct a regular expression usable to match and/or parse a C-Instruction
func compileCInstructionRegexp() *regexp.Regexp {
	destRe := "(?:([ADM]+)=)?" // the destination part of the regexp
	compRe := pipe(CComp)      // the compute part of the regexp
	jumpRe := pipe(CJump)      // the jump part of the regexp
	tmp := fmt.Sprintf("^%s%s(?:;%s)?$", destRe, compRe, jumpRe)
	return regexp.MustCompile(tmp)
}

// Return True if this looks like a C instruction
func isCInstruction(i Instruction) bool {
	return compileCInstructionRegexp().MatchString(string(i))
}

// Assemble a single C instruction from text into binary
func (i CInstruction) Assemble(st SymbolTable) (Code, error) {

	// extract dest, comp, and jump expressions from C instruction with regexp
	dest, comp, jump, err := splitCInstruction(i)
	if err != nil {
		log.Fatalf("error splitting instruction: %v", err)
	}

	// calculate destination bits: M=1, D=2, A=4
	dbits := 0
	if strings.Contains(dest, "M") {
		dbits |= 1
	}
	if strings.Contains(dest, "D") {
		dbits |= 2
	}
	if strings.Contains(dest, "A") {
		dbits |= 4
	}

	// calculate comp bits
	cbits, ok := CComp[comp]
	if !ok {
		log.Fatalf("error finding comp bits for %v", comp)
	}

	// calculate jump bits
	jbits, ok := CJump[jump]
	if !ok {
		log.Fatalf("error finding jump bits for %v", jump)
	}

	return Code((0b111 << 13) | (cbits << 8) | (dbits << 6) | jbits), nil
}

// Split up a C instructions into parts
func splitCInstruction(i CInstruction) (string, string, string, error) {
	re := compileCInstructionRegexp()
	match := re.FindStringSubmatch(string(i))
	if len(match) > 0 {
		return match[1], match[2], match[3], nil
	}
	return ``, ``, ``, fmt.Errorf("error splitting C instruction: %v", string(i))
}

func (i CInstruction) Resolve(s SymbolTable, a Address) Address {

	return a + 1
}
