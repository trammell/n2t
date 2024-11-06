package main

import (
	"bufio"
	"io"
)

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

type Command struct {
	vmCommand string
	fields []string
}

type Parser struct {
	Source  string
	Scanner *bufio.Scanner
}

type CodeWriter struct {
	Writer io.Writer
}

type Translatable interface {
	GetAsm() string
}
