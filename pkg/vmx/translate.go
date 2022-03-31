package vmx

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/rs/zerolog/log"
)

func Translate(file string) {
	log.Info().Msgf(`input file/dir is "%s"`, file)

	basename := path.Base(file)
	out := strings.TrimSuffix(basename, filepath.Ext(basename)) + ".asm"
	log.Info().Msgf(`output file is "%s"`, out)

	// construct CodeWriter
	cw := NewCodeWriter(out)
	defer cw.Close()

	// construct parser(s)
	for _, fn := range getSourceFiles(file) {
		log.Info().Msgf(`Parsing source file: "%s"`, fn)
		p := NewParser(fn)
		for p.Scan() {
			//cw.writeComment(p.Text())
			fmt.Println(p.Scanner.Text())
			switch p.commandType() {
			case C_ARITHMETIC:
				cw.writeComment("arithmetic: " + p.Text())
				log.Info().Msgf(`looks like arithmetic`)
				cw.writeArithmetic(p.Text())
			case C_POP, C_PUSH:
				cw.writeComment("pushpop: " + p.Text())
				log.Info().Msgf(`looks like pushpop`)
				cw.writePushPop(p.commandType(), p.arg1(), p.arg2())
			default:
				cw.writeComment("unknown: " + p.Text())
				log.Info().Msgf(`unknown type`)
			}
		}
	}
}

// Is this a file or a directory? If it's a directory, return all the .vm
// files in the directory.
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
