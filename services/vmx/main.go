package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

// The main VM translator function: reads source file name from os.Args.
// Read the instructions, resolve symbols, and emit the assembled code
func main() {

	srcfiles, err := getSourceFiles(os.Args[1])
	if err != nil {
		log.Fatalf("Unable to get source files: %s", err)
	}

	destfile, err := getDestFile(os.Args[1])
	if err != nil {
		log.Fatalf("Unable to get destination file: %s", err)
	}
	log.Printf(`Opening output file "%s"\n`, destfile)
	file, err := os.Create(destfile)
	if err != nil {
		log.Fatalf(`Unable to create "%v": %s`, destfile, err)
	}
	defer file.Close()
	w := bufio.NewWriter(file)
	defer w.Flush()

	// translate all lines in all source files
	for _, srcfile := range srcfiles {
		log.Printf(`Translating source file "%s"\n`, srcfile)

//srcFile, err := os.Open(srcFileName)
//		if err != nil {
//			fmt.Println(err)
//		}
//
//		scanner := bufio.NewScanner(srcFile)
//		for scanner.Scan() {
//			//fmt.Println(`// ` + scanner.Text())
//			cmd := NewCommand(scanner.Text())
//			out.WriteString(cmd.GetAsm())
//		}
//
//		if err := scanner.Err(); err != nil {
//			fmt.Println(err)
//		}

		// p := NewParser(fn)
		// for p.Scan() {
		// 	//cw.writeComment(p.Text())
		// 	fmt.Println(p.Scanner.Text())
		// 	switch p.commandType() {
		// 	case C_ARITHMETIC:
		// 		cw.writeComment("arithmetic: " + p.Text())
		// 		log.Info().Msgf(`looks like arithmetic`)
		// 		cw.writeArithmetic(p.Text())
		// 	case C_POP, C_PUSH:
		// 		cw.writeComment("pushpop: " + p.Text())
		// 		log.Info().Msgf(`looks like pushpop`)
		// 		cw.writePushPop(p.commandType(), p.arg1(), p.arg2())
		// 	default:
		// 		cw.writeComment("unknown: " + p.Text())
		// 		log.Info().Msgf(`unknown type`)
		// 	}
		// }
	}

	log.Println("Done.")
}



// If `src` is a directory, return all the .vm files in the directory,
// otherwise return a single .vm filename.
func getSourceFiles(path string) ([]string, error) {
	stat, err := os.Stat(path)
	if err != nil {
		return []string{}, err
	}

	// if it's a regular file, just return it back
	if stat.IsRegular() {
		if strings.HasSuffix(path, ".vm") {
			err := fmt.Errorf(`"%s" is the wrong type, should be .vm`, path)
			return []string{}, err
		} else {
			return []string{path}, nil
		}
	}

	// if it's a directory, match out the .vm files
	if stat.IsDir() {
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
			if v.IsRegular() && strings.HasSuffix(v.Name(), ".vm") {
				vmfiles = append(vmfiles, v.Name())
			}
		}
		if len(vmfiles) == 0 {
			err := fmt.Errorf(`No .vm files found in directory "%s"`, path)
			return []string{}, err
		}
		return vmfiles, nil
	}

	log.Fatalf(`File "%s" is not a file or a directory.`, path)
	return []string{}
}

// If `src` is a directory, return all the .vm files in the directory,
// otherwise return a single .vm filename.
func getDestFile(path string) (string, error) {
	stat, err := os.Stat(path)
	if err != nil {
		return "", err
	}

	// input path might be a regular file
	if stat.IsRegular() {
		log.Printf(`"%s" is a regular file`, path)
		if strings.HasSuffix(path, ".vm") {
			log.Fatalf(`"%s" is the wrong type, should be .vm`, path)
		} else {
			return regexp.MustCompile(`.vm$`).ReplaceAllString(path, ".asm")
		}
	}

	if stat.IsDir() {
		log.Printf(`"%s" is a directory`, path)
		return files, nil
	}

}
