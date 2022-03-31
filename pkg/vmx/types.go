package vmx

import (
	"bufio"
	"os"
)

type Command int

// These are all our command types
const (
	C_ARITHMETIC = iota
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
	Source  string
	Scanner *bufio.Scanner
}

type CodeWriter struct {
	FileName   string
	Writer     *bufio.Writer
	FileHandle *os.File
}
