package main

import (
	"fmt"
)

// push values from the stack to various segments
func (cw *CodeWriter) writeLabel(label string) (error) {
	_, err := fmt.Fprintf(cw.Writer, "// label\n(%s)\n\n", label)
	return err
}
