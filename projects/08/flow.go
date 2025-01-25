package main

import "fmt"

// write assembly code equivalent to a VM label command
func (cw *CodeWriter) writeLabel(label string) error {
	_, err := fmt.Fprintf(cw.Writer, "// label\n(%s)\n\n", label)
	return err
}

// write assembly code equivalent to a VM goto command
func (cw *CodeWriter) writeGoto(label string) error {
	format := `// goto %[1]s
@%[1]s
0; JMP`
	_, err := fmt.Fprintf(cw.Writer, format+"\n\n", label)
	return err
}

// write assembly code equivalent to a VM if-goto command
// NOTE: this modifies the stack (pops top value)
func (cw *CodeWriter) writeIf(label string) error {
	asm := `// if-goto %[1]s
@SP
AM=M-1
D=M
@%[1]s
D; JNE`
	_, err := fmt.Fprintf(cw.Writer, asm+"\n\n", label)
	return err
}
