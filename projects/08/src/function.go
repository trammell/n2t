package main

import "fmt"

// write assembly code equivalent to a VM label command
func (cw *CodeWriter) writeFunction(fname string, nargs int) error {
	format := `// call %[1]s %[2]d
@RETURN_%[1]s

push LCL 
push ARG
push THIS
push THAT
ARG = SP - n - 5 ???
LCL = SP
goto f

(RETURN_%[1]s_)
`
	_, err := fmt.Fprintf(cw.Writer, format, fname, nargs)
	return err
}
