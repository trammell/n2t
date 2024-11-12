package main

import (
	"fmt"
)

// push values from the stack to various segments
func (cw *CodeWriter) writeIf(label string) (error) {
	asm := `// if-goto %[1]s
@SP
A=M
D=A
@%[1]s
D; JNE`
	_, err := fmt.Fprintf(cw.Writer, asm + "\n\n", label)
	return err
}
