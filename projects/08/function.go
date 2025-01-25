package main

import "fmt"

// write assembly code equivalent to a VM "function" command
func (cw *CodeWriter) writeFunction(fname string, nlocals int) error {
	asm := ""
	format := "// function %[2]s %[3]d\n(%[1]s$%[2]s)\n"
	asm += fmt.Sprintf(format, cw.VMFile, fname, nlocals)
	for i := 0; i < nlocals; i++ {
		format := `@0    // push constant 0 (%d of %d)
D=A
@SP
A=M
M=D
@SP
M=M+1`
		asm += fmt.Sprintf(format, i+1, nlocals)
	}
	_, err := fmt.Fprintf(cw.Writer, cw.VMFile, fname)
	return err
}
