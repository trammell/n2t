// vim: set ai ts=4 :

package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func NewParser(filename string) (*Parser, error) {
	p := new(Parser)
	p.FileName = filename
	lines, err := readLines(filename)
	if err != nil {
		return nil, err
	}
	p.lines = lines
	p.index = -1
	return p, nil
}

func (p Parser) hasMoreCommands() bool {
	return p.index < len(p.lines)-1
}

func (p *Parser) advance() {
	if p.hasMoreCommands() {
		p.index += 1
	}
}

func (p *Parser) reset() {
	p.index = -1
}

func (p Parser) commandType() uint8 {
	cur := p.lines[p.index]
	if regexp.MustCompile(`^@(\d+|[[:alnum:]])$`).MatchString(cur) {
		return A_COMMAND
	}
	if regexp.MustCompile(`^\(.+\)$`).MatchString(cur) {
		return L_COMMAND
	}
	_, _, _, err := ParseCInstruction(cur)
	if err == nil {
		return C_COMMAND
	}
	log.Fatalf("Unrecognized command: %v", cur)
	return X_COMMAND
}

func (p Parser) symbol() (string, error) {
	cur := p.lines[p.index]

	// A instructions
	if regexp.MustCompile(`^@(\d+|[[:alnum:]])$`).MatchString(cur) {
		return regexp.MustCompile(`^@`).ReplaceAllString(cur, ""), nil
	}

	// labels, minus the parentheses
	if regexp.MustCompile(`^\(.+\)$`).MatchString(cur) {
		return strings.Trim(cur, "()"), nil
	}

	return "", fmt.Errorf("unrecognized symbol: %s", cur)
}

// thin wrapper to ParseCInstruction()
func (p Parser) destCompJump() (string, string, string, error) {
	cur := p.lines[p.index]
	return ParseCInstruction(cur)
}

// read source file into `lines`
func readLines(src string) ([]string, error) {
	file, err := os.Open(src)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inst := stripInstruction(scanner.Text())
		if inst != "" {
			lines = append(lines, inst)
		}
	}
	return lines, scanner.Err()
}

// Remove all comments and whitespace from the instruction
func stripInstruction(inst string) string {
	inst = regexp.MustCompile(`//.*`).ReplaceAllString(inst, "")
	inst = regexp.MustCompile(`\s`).ReplaceAllString(inst, "")
	return inst
}

// attempt to parse a string as a C instruction. This is a combination of the
// dest(), comp() and jump() methods as described in the N2T book, page 114.
func ParseCInstruction(inst string) (string, string, string, error) {
	// cut out the "dest" part of the instruction
	dest, rest, found_eq := strings.Cut(inst, "=")
	comp, jump, found_semi := strings.Cut(rest, ";")
	if found_eq || found_semi {
		return dest, comp, jump, nil
	} else {
		return "", "", "", errors.New("No = or ; found")
	}
}
