package main

const (
	A_COMMAND uint8 = iota
	C_COMMAND
	L_COMMAND
	X_COMMAND
)

type Parser struct {
	FileName string
	lines []string
	index int
}

type Code struct {}
