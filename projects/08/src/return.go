package main

import "fmt"	

// write assembly code equivalent to a VM label command	
func (cw *CodeWriter) writeReturn() (error) {
	asm := `// return



	`
	_, err := fmt.Fprintf(cw.Writer, asm + "\n\n")
	return err
}
