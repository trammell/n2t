package main

import (
	"fmt"
	"log"
	"os"
)


// main function: takes no arguments, reads source file name from os.Args
// Read the instructions, resolve symbols, and emit the assembled code
func main() {

	// construct the parser object
	p, err := NewParser(os.Args[1])
	if err != nil {
		log.Fatalf("NewParser: %s", err)
	}

	// first pass: loop through the commands printing their type
	for {
		if !p.hasMoreCommands() {
			break
		}
		p.advance()
		fmt.Println(p.index, p.commandType())
	}
}
