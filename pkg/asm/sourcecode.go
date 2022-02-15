package asm

import (
	"bufio"
	"log"
	"os"
	"regexp"
)

// NewProgram is a constructor for Programs
func NewSourceCode(filename string) (sc *SourceCode) {
	sc = new(SourceCode)
	sc.FileName = filename
	fh, err := os.Open(sc.FileName)
	if err != nil {
		log.Fatal(err)
	}
	sc.Scanner = bufio.NewScanner(fh)
	return sc
}

func (sc SourceCode) GetInstruction() (inst string, more bool) {
	for sc.Scanner.Scan() {
		txt := sc.Scanner.Text()
		if IsInstruction(txt) {
			return txt, true
		}
	}
	if err := sc.Scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return ``, false
}

// Return True if _txt_ is a Hack ASM instruction
func IsInstruction(txt string) bool {
	txt = regexp.MustCompile(`//.*`).ReplaceAllString(txt, "")
	txt = regexp.MustCompile(`\s`).ReplaceAllString(txt, "")
	return txt != ``
}
