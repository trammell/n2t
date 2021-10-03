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

// pipe() is a helper function for building the C instruction regular expression.
// We're looking to build a big "OR" regular expression like "(a|b|c|d|...)". Some
// of the keys contain metacharacters like "|", "+" and "-" so we quotemeta() them.
func pipe(m map[string]string) string {
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
func isCInstruction(txt string) bool {
	re := compileCInstructionRegexp()
	return re.MatchString(txt)
}

// Assemble a single C instruction from text into binary
func (i CInstruction) Assemble(symbols map[string]int) (string, error) {

	// extract dest, comp, and jump expressions from C instruction with regexp
	dest, comp, jump, err := splitCInstruction(i)
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

// Split up a C instructions into parts
func splitCInstruction(i CInstruction) (string, string, string, error) {
	re := compileCInstructionRegexp()
	match := re.FindStringSubmatch(string(i))
	if len(match) > 0 {
		return match[1], match[2], match[3], nil
	}
	return ``, ``, ``, fmt.Errorf("error splitting C instruction: %v", string(i))
}

func (i CInstruction) ResolveSymbol (s SymbolTable, a Address) Address {

	return a + 1
}