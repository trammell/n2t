package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// return a new CodeWriter object
func NewCodeWriter(outfile string) (*CodeWriter, error) {
	cw := new(CodeWriter)
	cw.Outfile = outfile

	// create and open the output file
	log.Printf("Opening file: %s\n", outfile)
	file, err := os.Create(outfile)
	if err != nil {
		return &CodeWriter{}, err
	}
	cw.File = file
	cw.Writer = bufio.NewWriter(file)
	return cw, nil
}

func (cw CodeWriter) Close() {
	cw.Writer.Flush()
	cw.File.Close()
}

func (cw *CodeWriter) setFileName(filename string) {
	cw.Infile = filename
}

// There are some opportunities in this function for code reuse, but I think
// it's clearer if all the assembly code is laid out in full. Maybe if I get
// better at reading .asm code.
func (cw *CodeWriter) writeArithmetic(cmd string) (error) {
	var asm string
	switch cmd {
	// unary operations
	case `neg`:
		asm = "// neg@SP\nA=M\nA=A-1\nM=-M\n\n"
	case `not`:
		asm = "// not@SP\nA=M\nA=A-1\nM=!M\n\n"
	// binary operations
	case `add`:
		asm = "// add\n@SP\nM=M-1\nA=M\nD=M\nA=A-1\nM=D+M\n\n"
	case `sub`:
		asm = "// sub\n@SP\nM=M-1\nA=M\nD=M\nA=A-1\nM=M-D\n\n"
	case `and`:
		asm = "// and\n@SP\nM=M-1\nA=M\nD=M\nA=A-1\nM=M&D\n\n"
	case `or`:
		asm = "// or\n@SP\nM=M-1\nA=M\nD=M\nA=A-1\nM=M|D\n\n"
	// comparisons
	case `eq`:
		cw.Counter++
		fmt = `// eq
@SP
M=M-1
A=M
D=M
A=A-1
D=D-M
@EQ_TRUE_%[1]d
D; JEQ
@SP
A=M
A=A-1
M=0
@EQ_CONT_%[1]d
0; JEQ
(EQ_TRUE_%[1]d)
@SP
A=M
A=A-1
M=-1
(EQ_CONT_%[1]d)
`
		asm = fmt.Sprintf(fmt, cw.Counter)

	case `gt`:
		cw.Counter++
		asm = fmt.Sprintf(`// gt
@SP
M=M-1
A=M
D=M
A=A-1
D=M-D
@GT_TRUE_%[1]d
D; JGT
@SP
A=M
A=A-1
M=0
@GT_CONT_%[1]d
0; JEQ
(GT_TRUE_%[1]d)
@SP
A=M
A=A-1
M=-1
(GT_CONT_%[1]d)
`, cw.Counter)

	case `lt`:
		cw.Counter++
		asm = fmt.Sprintf(`// lt
@SP
M=M-1
A=M
D=M
A=A-1
D=M-D
@LT_TRUE_%[1]d
D; JLT
@SP
A=M
A=A-1
M=0
@LT_CONT_%[1]d
0; JEQ
(LT_TRUE_%[1]d)
@SP
A=M
A=A-1
M=-1
(LT_CONT_%[1]d)
`, cw.Counter)
	default:
		return fmt.Errorf(`Unrecognized command: %s`, cmd)
	}
	_, err := fmt.Fprintln(cw.Writer, asm)
	return err
}

func (w *CodeWriter) writePushPop(cmd uint8, segment string, index int) (error) {
	var asm string
	switch cmd {
	case C_PUSH:
		switch segment {
		case `constant`:
			asm = fmt.Sprintf("// push constant %d\n", index)
			asm += fmt.Sprintf("@%d\nD=A\n@SP\nA=M\nM=D\n@SP\nM=M+1\n\n", index)
		default:
			return fmt.Errorf(`Unrecognized segment: %s`, segment)
		}
	case C_POP:
		switch segment {
		case `constant`:
			return fmt.Errorf("Can't POP to constant segment")
		default:
			return fmt.Errorf(`Unrecognized segment: %s`, segment)
		}
	}
	fmt.Fprintf(w.Writer, asm)
	return nil
}
