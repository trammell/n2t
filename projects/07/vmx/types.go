package main

import (
	"io"
	"os"
)

// These are all our command types
const (
	C_ARITHMETIC uint8 = iota
	C_PUSH
	C_POP
	C_LABEL
	C_GOTO
	C_IF
	C_FUNCTION
	C_RETURN
	C_CALL
	C_UNDEFINED
)

type Parser struct {
	FileName string
	Lines    []string
	Index    int
}

type CodeWriter struct {
	Infile  string     // the name of the file currently being parsed
	Outfile string
	File    *os.File
	Writer  io.Writer
	Counter int
}
