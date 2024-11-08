package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
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
		log.Printf(`Translating source file "%s"\n`, srcfile)
		cw.setFileName(srcfile)
		p, err := NewParser(srcfile)
		if err != nil {
			log.Fatalf("Unable to construct parser: %s", err)
		}
		for p.hasMoreCommands() {
			p.advance()
			ct, err := p.commandType()
			if err != nil {
				log.Fatalf("Unknown command type: %s", err)
			}
			switch ct {
			case C_ARITHMETIC:
				cw.writeArithmetic(p.arg1())
			case C_PUSH, C_POP:
				cw.writePushPop(ct, p.arg1(), p.arg2())
			}
		}
	}
	log.Println("Done.")
}

// If `src` is a directory, return all the .vm files in the directory,
// otherwise return a single .vm filename.
func getSourceFiles(path string) ([]string, error) {
	info, err := os.Stat(path)
	if err != nil {
		return []string{}, err
	}

	// if it's a directory, match out the .vm files
	if info.IsDir() {
		log.Printf(`source file "%s" is a directory`, path)
		f, err := os.Open(path)
		if err != nil {
			return []string{}, err
		}
		files, err := f.Readdir(0)
		if err != nil {
			return []string{}, err
		}
		var vmfiles []string
		for _, v := range files {
			if v.Mode().IsRegular() && strings.HasSuffix(v.Name(), ".vm") {
				vmfiles = append(vmfiles, v.Name())
			}
		}
		if len(vmfiles) == 0 {
			err := fmt.Errorf(`No .vm files found in directory "%s"`, path)
			return []string{}, err
		}
		return vmfiles, nil
	} else {
		if strings.HasSuffix(path, ".vm") {
			err := fmt.Errorf(`"%s" is the wrong type, should be .vm`, path)
			return []string{}, err
		} else {
			return []string{path}, nil
		}
	}

	err = fmt.Errorf(`File "%s" is not a file or a directory.`, path)
	return []string{}, err
}

// If `src` is a directory, return all the .vm files in the directory,
// otherwise return a single .vm filename.
func getDestFile(path string) (string, error) {
	stat, err := os.Stat(path)
	if err != nil {
		return "", err
	}
	if stat.IsDir() {
		log.Printf(`"%s" is a directory`, path)
		dest := path + ".vm"
		return dest, nil
	}
	// it must be a file
	log.Printf(`"%s" is a regular file`, path)
	if strings.HasSuffix(path, ".vm") {
		dest := regexp.MustCompile(`.vm$`).ReplaceAllString(path, ".asm")
		return dest, nil
	}
	return "", fmt.Errorf(`"%s" is the wrong type, should be .vm`, path)
}
