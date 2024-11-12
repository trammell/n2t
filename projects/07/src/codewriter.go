package main

import (
	"bufio"
	"fmt"
	"log"
	"path/filepath"
	"os"
	"regexp"
	"strings"
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

// Set the name of the current file, minus any preceding path or file
// extension. Needed for static segments.
func (cw *CodeWriter) setFileName(vmfile string) {
	suffix := filepath.Ext(vmfile)					// find the suffix
	vmfile = filepath.Base(vmfile)					// remove path
	vmfile = strings.TrimSuffix(vmfile, suffix)	// strip suffix
	pattern := regexp.MustCompile(`[^A-Za-z0-9_]`)	// problematic chars pattern
	vmfile = pattern.ReplaceAllString(vmfile, "")	// remove problematic chars
	cw.VMFile = vmfile
}

// There are some opportunities in this function for code reuse, but I think
// it's clearer if all the assembly code is laid out in full. Maybe if I get
// better at reading .asm code.
func (cw *CodeWriter) writeArithmetic(cmd string) (error) {
	var asm string
	switch cmd {
	// unary operations
	case `neg`:
		asm = `// neg
@SP
A=M
A=A-1
M=-M`
	case `not`:
		asm = `// not
@SP
A=M
A=A-1
M=!M`
	// binary operations
	case `add`:
		asm = `// add
@SP
M=M-1
A=M
D=M
A=A-1
M=D+M`
	case `sub`:
		asm = `// sub
@SP
M=M-1
A=M
D=M
A=A-1
M=M-D`
	case `and`:
		asm = `// and
@SP
M=M-1
A=M
D=M
A=A-1
M=M&D`
	case `or`:
		asm = `// or
@SP
M=M-1
A=M
D=M
A=A-1
M=M|D`
	// comparisons
	case `eq`:
		cw.Counter++
		format := `// eq
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
(EQ_CONT_%[1]d)`
		asm = fmt.Sprintf(format, cw.Counter)

	case `gt`:
		cw.Counter++
		format := `// gt
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
(GT_CONT_%[1]d)`
		asm = fmt.Sprintf(format, cw.Counter)

	case `lt`:
		cw.Counter++
		format := `// lt
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
(LT_CONT_%[1]d)`
		asm = fmt.Sprintf(format, cw.Counter)

	default:
		return fmt.Errorf(`Unrecognized command: %s`, cmd)
	}
	_, err := fmt.Fprintf(cw.Writer, asm + "\n\n")
	return err
}

// push various values on to the stack from various segments
func (cw *CodeWriter) writePush(segment string, index int) (error) {
	var asm string
	switch segment {
	case `constant`:
		format := `// push constant %[1]d
@%[1]d
D=A
@SP
A=M
M=D
@SP
M=M+1`
		asm = fmt.Sprintf(format, index)

	case `local`:
		format := `// push local %[1]d
@%[1]d
D=A
@LCL
A=M
A=A+D
D=M
@SP
A=M
M=D
@SP
M=M+1`
		asm = fmt.Sprintf(format, index)

	case `argument`:
		format := `// push argument %[1]d
@%[1]d
D=A
@ARG
A=M
A=A+D
D=M
@SP
A=M
M=D
@SP
M=M+1`
		asm = fmt.Sprintf(format, index)
	case `this`:
		format := `// push this %[1]d
@%[1]d
D=A
@THIS
A=M
A=A+D
D=M
@SP
A=M
M=D
@SP
M=M+1`
		asm = fmt.Sprintf(format, index)
	case `that`:
		format := `// push that %[1]d
@%[1]d
D=A
@THAT
A=M
A=A+D
D=M
@SP
A=M
M=D
@SP
M=M+1`
		asm = fmt.Sprintf(format, index)

	case `pointer`:
		if index < 0 || index > 1 {
			return fmt.Errorf(`Invalid command: "push pointer %d"`, index)
		}
		format := `// push pointer %[1]d
@%[2]d
D=M
@SP
A=M
M=D
@SP
M=M+1`
		asm = fmt.Sprintf(format, index, index + 3)

	case `temp`:
		if index < 0 || index > 7 {
			return fmt.Errorf(`Invalid command: "push temp %d"`, index)
		}
		format := `// push temp %[1]d
@%[2]d
D=M
@SP
A=M
M=D
@SP
M=M+1`
		asm = fmt.Sprintf(format, index, index + 5)

	case `static`:
		format := `// push static %[1]d
@%[2]s.%[1]d
D=M
@SP
A=M
M=D
@SP
M=M+1`
		asm = fmt.Sprintf(format, index, cw.VMFile)
	default:
		return fmt.Errorf(`Unrecognized segment: %s`, segment)
	}
	_, err := fmt.Fprintf(cw.Writer, asm + "\n\n")
	return err
}

// push values from the stack to various segments
func (cw *CodeWriter) writePop(segment string, index int) (error) {
	var asm string
	switch segment {
	case `constant`:
		return fmt.Errorf("Can't POP to constant segment")
	case `local`:
		format := `// pop local %[1]d
@%[1]d
D=A
@LCL
D=D+M
@R15
M=D
@SP
AM=M-1
D=M
@R15
A=M
M=D`
		asm = fmt.Sprintf(format, index)

	case `argument`:
		format := `// pop argument %[1]d
@%[1]d
D=A
@ARG
D=D+M
@R15
M=D
@SP
AM=M-1
D=M
@R15
A=M
M=D`
		asm = fmt.Sprintf(format, index)

	case `this`:
		format := `// pop this %[1]d
@%[1]d
D=A
@THIS
D=D+M
@R15
M=D
@SP
AM=M-1
D=M
@R15
A=M
M=D`
		asm = fmt.Sprintf(format, index)

	case `that`:
		format := `// pop that %[1]d
@%[1]d
D=A
@THAT
D=D+M
@R15
M=D
@SP
AM=M-1
D=M
@R15
A=M
M=D`
		asm = fmt.Sprintf(format, index)

	case `pointer`:
		// This is simple as there's no real math to do. It might be easier to
		// just have separate ASM statements for cases 0, 1
		if index < 0 || index > 1 {
			return fmt.Errorf(`Invalid command: "pop pointer %d"`, index)
		}
		format := `// pop pointer %[1]d
@SP
AM=M-1
D=M
@%[2]d
M=D`
		asm = fmt.Sprintf(format, index, index + 3)

	case `temp`:
		if index < 0 || index > 7 {
			return fmt.Errorf(`Invalid command: "pop temp %d"`, index)
		}
		format := `// pop temp %[1]d
@SP
AM=M-1
D=M
@%[2]d
M=D`
		asm = fmt.Sprintf(format, index, index + 5)

	case `static`:
		format := `// pop static %[1]d
@SP
AM=M-1
D=M
@%[2]s.%[1]d
M=D`
		asm = fmt.Sprintf(format, index, cw.VMFile)

	default:
		return fmt.Errorf(`Unrecognized segment: %s`, segment)
	}
	_, err := fmt.Fprintf(cw.Writer, asm + "\n\n")
	return err
}
