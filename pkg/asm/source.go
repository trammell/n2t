package asm

import (
	"bufio"
	"log"
	"os"
	"regexp"
)

// NewSource is a constructor for code sources
func NewSource(filename string) (src *Source) {
	src = new(Source)
	src.FileName = filename
	fh, err := os.Open(src.FileName)
	if err != nil {
		log.Fatal(err)
	}
	src.Scanner = bufio.NewScanner(fh)
	return src
}

func (src Source) GetInstruction() (inst string, more bool) {
	for src.Scanner.Scan() {
		txt := src.Scanner.Text()
		if IsInstruction(txt) {
			return txt, true
		}
	}
	if err := src.Scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return ``, false
}

// return True if _txt_ is a comment
func IsComment(txt string) bool {
	return regexp.MustCompile(`^//`).MatchString(txt)
}

// Return True if _txt_ is a Hack ASM instruction
func IsInstruction(txt string) bool {
	if IsComment(txt) {
		return false
	}
	txt = regexp.MustCompile(`//.*`).ReplaceAllString(txt, "")
	txt = regexp.MustCompile(`\s`).ReplaceAllString(txt, "")
	return IsAInstruction(Instruction(txt)) ||
		IsCInstruction(Instruction(txt)) ||
		IsLabel(Instruction(txt))
}
