package main

import "fmt"

// push values from the stack to various segments
func (cw *CodeWriter) writeInit() error {
	asm := `// init
@256		// set stack pointer
D=A
@SP
M=D
@Sys.init	// call Sys.init
0; JMP`
	_, err := fmt.Fprintf(cw.Writer, asm+"\n\n")
	return err
}
