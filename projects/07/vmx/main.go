package main

import (
	"log"
	"os"
)

// The main VM translator function: reads source file name from os.Args.
// Read the instructions, resolve symbols, and emit the assembled code
func main() {

	srcfiles, err := getSourceFiles(os.Args[1])
	if err != nil {
		log.Fatalf("Unable to get source files: %s", err)
	}

	// construct the codewriter object
	destfile, err := getDestFile(os.Args[1])
	if err != nil {
		log.Fatalf("Unable to get destination file: %s", err)
	}
	cw, err := NewCodeWriter(destfile)
	if err != nil {
		log.Fatalf("Unable to open code writer: %s", err)
	}
	defer cw.Close()

	// translate all lines in all source files
	for _, srcfile := range srcfiles {
		log.Printf(`Translating source file "%s"`, srcfile)
		cw.setFileName(srcfile)
		p, err := NewParser(srcfile)
		if err != nil {
			log.Fatalf(`Unable to construct parser: "%s"`, err)
		}
		for p.hasMoreCommands() {
			p.advance()
			ct, err := p.commandType()
			if err != nil {
				log.Fatalf(`Error finding command type: "%s"`, err)
			}
			switch ct {
			case C_ARITHMETIC:
				err := cw.writeArithmetic(p.currentCommand())
				if err != nil {
					log.Fatalf(`Error writing arithmetic: "%s"`, err)
				}
			case C_PUSH, C_POP:
				arg1, err := p.arg1()
				if err != nil {
					log.Fatalf(`Error fetching arg1: "%s"`, err)
				}
				arg2, err := p.arg2()
				if err != nil {
					log.Fatalf(`Error fetching arg2: "%s"`, err)
				}
				err = cw.writePushPop(ct, arg1, arg2)
				if err != nil {
					log.Fatalf(`Error writing arithmetic: "%s"`, err)
				}
			}
		}
	}
	log.Println("Done.")
}
