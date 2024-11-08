package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// return a new CodeWriter object
func NewCodeWriter(filename string) (CodeWriter, error) {
	log.Printf("Writing code to file: %s", filename)

	// don't clobber output file
	if _, err := os.Stat(filename); err == nil {
		return nil, fmt.Errorf(`%s exists already`, filename)
	}
}

func (cw CodeWriter) Close() {
	cw.file.Close()
	cw.writer.Flush()
}

func (cw *CodeWriter) setFileName(filename string) {
	cw.fileName = filename
}

func (cw *CodeWriter) writeArithmetic(cmd string) {
	switch cmd {
	// unary operations
	case `neg`:
		return `// neg@SP\nA=M\nA=A-1\nM=-M`, nil
	case `not`:
		return `// neg@SP\nA=M\nA=A-1\nM=!M`, nil
	// binary operations
	case `add`:
		return `// add\n@SP\nM=M-1\nA=M\nD=M\nA=A-1\nM=D+M`, nil
	case `sub`:
		return `// sub\n@SP\nM=M-1\nA=M\nD=M\nA=A-1\nM=M-D`, nil
	case `and`:
		return `// and\n@SP\nM=M-1\nA=M\nD=M\nA=A-1\nM=M&D`, nil
	case `or`:
		return `// or\n@SP\nM=M-1\nA=M\nD=M\nA=A-1\nM=M|D`, nil
	// comparisons
	case `eq`:
		return `// FIXME\n@SP\nM=M-1\nA=M\nD=M\nA=A-1\nM=M-D`, nil
	case `gt`:
		return `// FIXME\n@SP\nM=M-1\nA=M\nD=M\nA=A-1\nM=M-D`, nil
	case `lt`:
		return `// FIXME\n@SP\nM=M-1\nA=M\nD=M\nA=A-1\nM=M-D`, nil
	default:
		return "", fmt.Errorf(`Unrecognized command: %s`, cmd)
	}
}

func (w *CodeWriter) writePushPop(cmd Command, segment string, index int) {
	switch cmd {
	case C_PUSH:
		switch segment {
		case `constant`:
			fmt.Fprintf(w.Writer, "@%d\nD=A\n@SP\nM=D\n", index)	// FIXME is this right?
		default:
			log.Fatal().Msgf(`Unrecognized segment: %s`, segment)
		}

	case C_POP:
		switch segment {
		case `constant`:
			log.Fatal().Msgf("Can't POP to constant segment")
		default:
			log.Fatal().Msgf(`Unrecognized segment: %s`, segment)
		}

	}
}
