package main

import (
	"log"
	"os"
)

// The main VM translator function: reads source file name from os.Args.
// Read the instructions, resolve symbols, and emit the assembled code
func main() {

	// get the name(s) of the source VM file(s) to translate
	srcfiles, err := getSourceFiles(os.Args[1])
	if err != nil {
		log.Fatalf("Unable to get source files: %s", err)
	}

	// construct the output file and the codewriter object
	destfile, err := getDestFile(os.Args[1])
	if err != nil {
		log.Fatalf("Unable to get destination file: %s", err)
	}
	cw, err := NewCodeWriter(destfile)
	if err != nil {
		log.Fatalf("Unable to open code writer: %s", err)
	}
	defer cw.Close()

	// write the init block
	cw.writeInit()

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
			case C_PUSH:
				arg1, err := p.arg1()
				if err != nil {
					log.Fatalf(`Error fetching arg1: "%s"`, err)
				}
				arg2, err := p.arg2()
				if err != nil {
					log.Fatalf(`Error fetching arg2: "%s"`, err)
				}
				err = cw.writePush(arg1, arg2)
				if err != nil {
					log.Fatalf(`Error writing push: "%s"`, err)
				}
			case C_POP:
				arg1, err := p.arg1()
				if err != nil {
					log.Fatalf(`Error fetching arg1: "%s"`, err)
				}
				arg2, err := p.arg2()
				if err != nil {
					log.Fatalf(`Error fetching arg2: "%s"`, err)
				}
				err = cw.writePop(arg1, arg2)
				if err != nil {
					log.Fatalf(`Error writing pop: "%s"`, err)
				}
			case C_LABEL:
				arg1, err := p.arg1()
				if err != nil {
					log.Fatalf(`Error fetching arg1: "%s"`, err)
				}
				err = cw.writeLabel(arg1)
				if err != nil {
					log.Fatalf(`Error writing label: "%s"`, err)
				}
			case C_GOTO:
				arg1, err := p.arg1()
				if err != nil {
					log.Fatalf(`Error fetching arg1: "%s"`, err)
				}
				err = cw.writeGoto(arg1)
				if err != nil {
					log.Fatalf(`Error writing goto: "%s"`, err)
				}
			case C_IF:
				arg1, err := p.arg1()
				if err != nil {
					log.Fatalf(`Error fetching arg1: "%s"`, err)
				}
				err = cw.writeIf(arg1)
				if err != nil {
					log.Fatalf(`Error writing if-goto: "%s"`, err)
				}
			case C_FUNCTION:
				arg1, err := p.arg1()
				if err != nil {
					log.Fatalf(`Error fetching arg1: "%s"`, err)
				}
				err = cw.setFunction(arg1)
				if err != nil {
					log.Fatalf(`Error setting function name: "%s"`, err)
				}
				arg2, err := p.arg2()
				if err != nil {
					log.Fatalf(`Error fetching arg2: "%s"`, err)
				}
				err = cw.writeFunction(arg1, arg2)
				if err != nil {
					log.Fatalf(`Error writing function: "%s"`, err)
				}
			case C_RETURN:
				err := cw.writeReturn()
				if err != nil {
					log.Fatalf(`Error writing return: "%s"`, err)
				}
			case C_CALL:
				arg1, err := p.arg1()
				if err != nil {
					log.Fatalf(`Error fetching arg1: "%s"`, err)
				}
				arg2, err := p.arg2()
				if err != nil {
					log.Fatalf(`Error fetching arg2: "%s"`, err)
				}
				err = cw.writeCall(arg1, arg2)
				if err != nil {
					log.Fatalf(`Error writing call: "%s"`, err)
				}
			}
		}
	}
	log.Println("Done.")
}
