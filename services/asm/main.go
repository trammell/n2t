package main

import (
	"log"
	"os"
)

// main function: takes no arguments, reads source file name from os.Args
// Read the instructions, resolve symbols, and emit the assembled code
func main() {

// 1. read source file into array
// 2. compile instructions one at a time
// 3. ???
// 4. profit

	f := NewSourceFile(os.Args[1])
	a := NewAssembler()
	a.Assemble(f)
	a.Emit()
}
