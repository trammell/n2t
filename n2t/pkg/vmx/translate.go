package vmx

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/rs/zerolog/log"
)

// Translate .vm source code into .asm code. Create one CodeWriter object,
// and one Parser object for each source file.
func Translate(file string) {
	log.Info().Msgf(`input file/dir is "%s"`, file)

	basename := path.Base(file) // get filename minus path
	outfname := strings.TrimSuffix(basename, filepath.Ext(basename)) + ".asm"
	log.Info().Msgf(`output file is "%s"`, outfname)

	// construct output filehandle
	out, err := os.Create(outfname)
	if err != nil {
		fmt.Println(err)
	}
	// close the file with defer
	defer out.Close()

	// translate all lines in all source files
	for _, srcFileName := range getSourceFiles(file) {
		log.Info().Msgf(`Translating source file: "%s"`, file)
		srcFile, err := os.Open(srcFileName)
		if err != nil {
			fmt.Println(err)
		}

		scanner := bufio.NewScanner(srcFile)
		for scanner.Scan() {
			//fmt.Println(`// ` + scanner.Text())
			cmd := NewCommand(scanner.Text())
			out.WriteString(cmd.GetAsm())
		}

		if err := scanner.Err(); err != nil {
			fmt.Println(err)
		}

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
}

// If `src` is a directory, return all the .vm files in the directory,
// otherwise return a single .vm filename.
func getSourceFiles(src string) []string {
	stat, err := os.Stat(src)
	if err != nil {
		log.Fatal().Err(err)
	}

	// glob out the .vm files if it's a directory
	if stat.IsDir() {
		log.Info().Msgf(`"%s" is a directory`, src)
		files := []string{}
		// need to list all files matching *.vm in dir
		return files
	}

	// it's not a directory, so return a slice with a single filename
	log.Info().Msgf(`"%s" is a file`, src)
	return []string{src}
}
