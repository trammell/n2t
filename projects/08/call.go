package main

import "fmt"

// write assembly code equivalent to a VM label command
func (cw *CodeWriter) writeCall(fname string, nargs int) error {
	cw.Counter++
	here := fmt.Sprintf("RETURN_%s.%s.%d", cw.VMFile, cw.Function, cw.Counter)
	dest := fmt.Sprintf("%s$%s", cw.VMFile, cw.Function)
	format := `// call %[1]s %[2]d
@%[3]s	// push return address
D=A
@SP
A=M
M=D
@SP
M=M+1
@LCL	// push LCL
A=M
D=A
@SP
A=M
M=D
@SP
M=M+1
@ARG	// push ARG
A=M
D=A
@SP
A=M
M=D
@SP
M=M+1
@THIS	// push THIS
A=M
D=A
@SP
A=M
M=D
@SP
M=M+1
@THAT	// push THAT
A=M
D=A
@SP
A=M
M=D
@SP
M=M+1
@SP		// ARG = SP - (n+5) ???
A=M
D=M
@%[4]d
D=D-A
@ARG
M=D
@SP		//   LCL = SP
A=M
D=M
@LCL
A=M
M=D
@%[5]s	// goto f
0; JMP
(%[3]s_)`
	_, err := fmt.Fprintf(cw.Writer, format, fname, nargs, here, nargs+5, dest)
	return err
}
