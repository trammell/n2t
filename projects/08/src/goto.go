package main

import (
	"fmt"
)

// push values from the stack to various segments
func (cw *CodeWriter) writeGoto(label string) (error) {
	format := `// goto %[1]s
@%[1]s
0; JMP`
	_, err := fmt.Fprintf(cw.Writer, format + "\n\n", label)
	return err
}
