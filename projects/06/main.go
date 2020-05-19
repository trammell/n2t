package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
)

var SYMBOLS = map[string]int{
	"R0": 0,
	"R1": 1,
	"R2": 2,
	"R3": 3,
}

type asm struct {
	filename string
	lines    []string
	sym      map[string]int
}

func (a *asm) parse() {
	file, err := os.Open(a.filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		a.appendLine(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func (a *asm) appendLine(txt string) {
	a.lines = append(a.lines, txt)
}

/****************************************************/

type inst struct {
	Txt  string
	Addr int
}

// Strip whitespace and comments from a line of assembler
func (i *inst) Canonicalize(txt string) string {
	txt = regexp.MustCompile(`//.*`).ReplaceAllString(txt, "")
	return regexp.MustCompile(`\s`).ReplaceAllString(txt, "")
}

func (i *inst) IsEmpty() bool {
	return true
}

func (i *inst) IsAInstruction() bool {
	return regexp.MustCompile(`^@.*`).MatchString(i.Txt)
}

func (i *inst) IsCInstruction() bool {
	return regexp.MustCompile(`^[ADM]*=?([^;]*)(;.*)?$`).MatchString(i.Txt)
}

func (i *inst) IsLabel() bool {
	return regexp.MustCompile(`^(.*)`).MatchString(i.Txt)
}

func (i *inst) Assemble(sym map[string]int) string {
	return "1111111111111111"
}

/********************************************************************/

func main() {
	a := asm{filename: os.Args[1]}
	a.parse()
}
