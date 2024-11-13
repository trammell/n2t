package main

import "fmt"

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
			return fmt.Errorf(`invalid command: "push pointer %d"`, index)
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
			return fmt.Errorf(`invalid command: "push temp %d"`, index)
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
		return fmt.Errorf(`unrecognized segment: %s`, segment)
	}
	_, err := fmt.Fprintf(cw.Writer, asm + "\n\n")
	return err
}
