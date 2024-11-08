// vim: set ts=4 :
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

// main function: takes no arguments, reads source file name from os.Args
// Read the instructions, resolve symbols, and emit the assembled code
func main() {

	// construct the parser object
	asmfile := os.Args[1]
	p, err := NewParser(asmfile)
	if err != nil {
		log.Fatalf("NewParser: %s", err)
	}

	// create and open the output file (foo.asm ==> foo.hack)
	hackfile := regexp.MustCompile(`.asm$`).ReplaceAllString(asmfile, ".hack")
	log.Printf("Opening file: %s\n", hackfile)
	file, err := os.Create(hackfile)
	if err != nil {
		log.Fatalf("Unable to create %v: %s", hackfile, err)
	}
	defer file.Close()
	w := bufio.NewWriter(file)
	defer w.Flush()

	// construct symbol table
	st := NewSymbolTable()

	// first pass: loop through the commands, filling the symbol table
	log.Println("Building symbol table.")
	p.reset()
	romAddr := 0
	for p.hasMoreCommands() {
		p.advance()
		switch ct := p.commandType(); ct {
		case A_COMMAND:
			romAddr++
		case C_COMMAND:
			romAddr++
		case L_COMMAND:
			sym, err := p.symbol()
			if err != nil {
				log.Fatalf("Error extracting symbol: %s", err)
			}
			log.Printf("+ Adding ROM symbol %s=%d\n", sym, romAddr)
			st.addEntry(sym, romAddr)
		}
	}
	log.Println("Symbol table with label addresses:")
	fmt.Println(st)

	// second pass: loop through the commands printing their type
	log.Println("Assembling.")
	p.reset()
	ramAddr := 16 // start address for new variables
	for p.hasMoreCommands() {
		fmt.Printf(".")
		p.advance()

		// print A instructions
		if p.commandType() == A_COMMAND {
			sym, err := p.symbol()
			if err != nil {
				log.Fatalf("Error extracting symbol: %s", err)
			}

			if regexp.MustCompile(`^\d+$`).MatchString(sym) {
				num, err := strconv.Atoi(sym)
				if err != nil {
					log.Fatalf("Invalid symbol: %s", err)
				}
				fmt.Fprintf(w, "%016b\n", num)
			} else if st.contains(sym) {
				num := st.getAddress(sym)
				fmt.Fprintf(w, "%016b\n", num)
			} else if regexp.MustCompile(`^[A-Za-z0-9_.$]+$`).MatchString(sym) {
				log.Printf("+ Adding variable %s=%d\n", sym, ramAddr)
				st.addEntry(sym, ramAddr)
				fmt.Fprintf(w, "%016b\n", ramAddr)
				ramAddr++
			} else {
				log.Fatalf("Error assembling A instruction: %s", sym)
			}

			continue
		}

		// print C instructions
		if p.commandType() == C_COMMAND {
			destbits, compbits, jumpbits, err := p.destCompJump()
			if err != nil {
				log.Fatalf("Error parsing C instruction: %s", err)
			}
			fmt.Fprintf(w, "111%s%s%s\n", compbits, destbits, jumpbits)
			continue
		}

		// ignore L instructions
		if p.commandType() == L_COMMAND {
			continue
		}

		log.Fatalf("Error, found instruction that is neither A nor C.")
	}
	fmt.Printf("\n")
	log.Println("Done.")
}
