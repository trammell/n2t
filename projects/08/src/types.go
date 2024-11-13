package main

import (
	"bufio"
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
	VMFile  string		// the name of the VM file currently being parsed
	Outfile string
	Function string		// the name of the function we're currently in 
	File    *os.File
	Writer  *bufio.Writer
	Counter int
}
