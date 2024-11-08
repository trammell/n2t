package main

import (
	"bufio"
	"io"
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

type Command struct {
	vmCommand string
	fields    []string
}

type Parser struct {
	FileName string
	lines    []string
	index    int
}

type CodeWriter struct {
	Writer io.Writer
}

type Translatable interface {
	GetAsm() string
}
