package main

import "fmt"

// write assembly code equivalent to a VM "return" command
func (cw *CodeWriter) writeReturn() error {
	asm := `// return



	`
	_, err := fmt.Fprintf(cw.Writer, asm+"\n\n")
	return err
}
