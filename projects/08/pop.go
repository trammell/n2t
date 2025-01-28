package main

import "fmt"

// push values from the stack to various segments
func (cw *CodeWriter) writePop(segment string, index int) error {
	var asm string
	switch segment {
	case `constant`:
		return fmt.Errorf("can't POP to constant segment")
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
			return fmt.Errorf(`invalid command: "pop pointer %d"`, index)
		}
		format := `// pop pointer %[1]d
@SP
AM=M-1
D=M
@%[2]d
M=D`
		asm = fmt.Sprintf(format, index, index+3)

	case `temp`:
		if index < 0 || index > 7 {
			return fmt.Errorf(`invalid command: "pop temp %d"`, index)
		}
		format := `// pop temp %[1]d
@SP
AM=M-1
D=M
@%[2]d
M=D`
		asm = fmt.Sprintf(format, index, index+5)

	case `static`:
		format := `// pop static %[1]d
@SP
AM=M-1
D=M
@%[2]s.%[1]d
M=D`
		asm = fmt.Sprintf(format, index, cw.VMFile)

	default:
		return fmt.Errorf(`unrecognized segment: %s`, segment)
	}
	_, err := fmt.Fprintln(cw.Writer, asm+"\n")
	return err
}
