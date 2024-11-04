// vim: set ai ts=4 :
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
	p, err := NewParser(os.Args[1])
	if err != nil {
		log.Fatalf("NewParser: %s", err)
	}

	// create and open the output file
	outfile := outputFileName(os.Args[1])
	file, err := os.Create(outfile)
	if err != nil {
		log.Fatalf("Unable to create %v: %s", outfile, err)
	}
	defer file.Close()
	w := bufio.NewWriter(file)
	defer w.Flush()

	// first pass: loop through the commands printing their type
	for {
		if !p.hasMoreCommands() {
			break
		}
		p.advance()

		// print A instructions
		if p.commandType() == A_COMMAND {
			sym, err := p.symbol()
			if err != nil {
				log.Fatalf("Error extracting symbol: %s", err)
			}
			num, err := strconv.Atoi(sym)
			if err != nil {
				log.Fatalf("Invalid symbol: %s", err)
			}
			fmt.Printf("%016b\n", num)
		}

		// print C instructions
		if p.commandType() == C_COMMAND {
			dest, comp, jump, err := p.destCompJump()
			if err != nil {
				log.Fatalf("Error parsing C instruction: %s", err)
			}
			var code = Code{}
			fmt.Printf("111%s%s%s\n", code.comp(comp), code.dest(dest), code.jump(jump))
		}

		//		fmt.Println(p.index, p.commandType())
		//    fmt.Fprintln(w, line)
	}
}

func outputFileName(infile string) string {
	return regexp.MustCompile(`.asm$`).ReplaceAllString(infile, ".hack")
}
