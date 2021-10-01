package cmd

import (
	"asm"
	"os"
)

// main function: takes a single filename as argument
// Read the instructions, resolve symbols, and emit assembled code
func main() {
	p := asm.NewProgram(os.Args[1])
	p.read()
	p.resolve()
	p.emit()
}
