package main

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// return a new CodeWriter object
func NewCodeWriter(outfile string) (*CodeWriter, error) {
	cw := new(CodeWriter)

	// reasonable initial values
	cw.Outfile = outfile
	cw.VMFile = "Sys"
	cw.Function = "init"

	// create and open the output file
	log.Printf("Opening output file: %s\n", outfile)
	file, err := os.Create(outfile)
	if err != nil {
		return &CodeWriter{}, err
	}
	cw.File = file
	cw.Writer = bufio.NewWriter(file)

	return cw, nil
}

func (cw CodeWriter) Close() {
	cw.Writer.Flush()
	cw.File.Close()
}

func (cw CodeWriter) Flush() {
	cw.Writer.Flush()
}

// Set the name of the current file, minus any preceding path or file
// extension. Needed for static segments.
func (cw *CodeWriter) setFileName(vmfile string) {
	suffix := filepath.Ext(vmfile)                 // find the suffix
	vmfile = filepath.Base(vmfile)                 // remove path
	vmfile = strings.TrimSuffix(vmfile, suffix)    // strip suffix
	pattern := regexp.MustCompile(`[^A-Za-z0-9_]`) // problematic chars pattern
	vmfile = pattern.ReplaceAllString(vmfile, "")  // remove problematic chars
	cw.VMFile = vmfile
}

// Set the name of the current file, minus any preceding path or file
// extension. Needed for static segments.
func (cw *CodeWriter) setFunction(fname string) error {
	// match and strip problematic characters
	pattern := regexp.MustCompile(`[^A-Za-z0-9_.]`)
	cw.Function = pattern.ReplaceAllString(fname, "")
	return nil
}
