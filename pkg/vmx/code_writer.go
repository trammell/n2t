package vmx

import (
	"bufio"
	"fmt"
	"os"

	"github.com/rs/zerolog/log"
)

func NewCodeWriter(filename string) CodeWriter {
	log.Info().Msgf("codewriter file: %s", filename)

	// don't clobber output file
	if _, err := os.Stat(filename); err == nil {
		log.Fatal().Msgf(`%s exists already`, filename)
	}

	fh, _ := os.Create(filename)
	writer := bufio.NewWriter(fh)
	return CodeWriter{
		FileName:   filename,
		FileHandle: fh,
		Writer:     writer,
	}
}

func (w *CodeWriter) Close() {
	w.Writer.Flush()
	w.FileHandle.Close()
}

func (w *CodeWriter) writeComment(comment string) {
	fmt.Fprintf(w.Writer, "// %s\n", comment)
}

func (w *CodeWriter) writeArithmetic(cmd string) {
	switch cmd {
	case `add`:
		fmt.Fprintf(w.Writer, "foo\n")
	}
}

func (w *CodeWriter) writePushPop(cmd Command, segment string, index int) {
	fmt.Fprintf(w.Writer, "@%d\n", index)
}
