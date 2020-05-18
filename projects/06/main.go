package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
)

type asm struct {
	filename string
	lines    []string
}

type inst struct {
	txt string
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

// Strip whitespace and comments from a line of assembler
func canonicalizeInstruction(txt string) string {
	txt = regexp.MustCompile(`//.*`).ReplaceAllString(txt, "")
	txt = regexp.MustCompile(`\s`).ReplaceAllString(txt, "")
	return txt
}

func isLabel(inst string) bool {
	return regexp.MustCompile(`^(.*)`).MatchString(inst)
}

func isAInstruction(inst string) bool {
	return regexp.MustCompile(`^@.*`).MatchString(inst)
}

func isCInstruction(inst string) bool {
	return regexp.MustCompile(`[ADM]*=?([^;]*);())`).MatchString(inst)
}

func main() {
	a := asm{filename: os.Args[1]}
	a.parse()
}
