package main

import (
	"bufio"
	"fmt"
	"os"
)

// return a new CodeWriter object
func NewCodeWriter(filename string) CodeWriter {
	log.Info().Msgf("codewriter file: %s", filename)

	// don't clobber output file
	if _, err := os.Stat(filename); err == nil {
		log.Fatal().Msgf(`%s exists already`, filename)
	}
}

// 	fh, _ := os.Create(filename)
// 	return CodeWriter{bufio.NewWriter(fh)}
// }

func (cw *CodeWriter) setFileName(filename string) {
	cw.fileName = filename
}

func (cw *CodeWriter) writeArithmetic(cmd string) {
	switch cmd {
	case `add`:
		asm := `// add\n@SP\nM=M-1\nA=M\nD=M\nA=A-1\nM=D+M`
	case `sub`:




	}
}

func (cw *CodeWriter) writeComment(comment string) {
	fmt.Fprintf(w.Writer, "// %s\n", comment)
}

// func (w *CodeWriter) writeIncrementSP() {
// 	fmt.Fprintf(w.Writer, "@SP\nM=M+1\n")
// }

// func (w *CodeWriter) writeDecrementSP() {
// 	fmt.Fprintf(w.Writer, "@SP\nM=M-1\n")
// }


// func (w *CodeWriter) writePushPop(cmd Command, segment string, index int) {
// 	switch cmd {
// 	case C_PUSH:
// 		switch segment {
// 		case `constant`:
// 			fmt.Fprintf(w.Writer, "@%d\nD=A\n@SP\nM=D\n", index)
// 		default:
// 			log.Fatal().Msgf(`Unrecognized segment: %s`, segment)
// 		}

// 	case C_POP:
// 		switch segment {
// 		case `constant`:
// 			log.Fatal().Msgf("Can't POP to constant segment")
// 		default:
// 			log.Fatal().Msgf(`Unrecognized segment: %s`, segment)
// 		}

// 	}
// }
