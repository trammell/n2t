package main

import "fmt"

// write assembly code equivalent to a VM "return" command
func (cw *CodeWriter) writeReturn() error {
	asm := `// return



	`
	_, err := fmt.Fprintln(cw.Writer, asm+"\n")
	return err
}
