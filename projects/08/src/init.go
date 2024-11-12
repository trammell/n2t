package main

import "fmt"

// push values from the stack to various segments
func (cw *CodeWriter) writeInit() (error) {
	asm := `// init: set stack pointer and call Sys.init
@256
D=A
@SP
M=D
@Sys.init
0; JMP`
	_, err := fmt.Fprintf(cw.Writer, asm + "\n\n")
	return err
}
