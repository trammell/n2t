// vim: set ts=4 :

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
	if regexp.MustCompile(`^@\d+$`).MatchString(cur) {
		return A_COMMAND
	}
	if regexp.MustCompile(`^@[A-Za-z0-9_.$]+$`).MatchString(cur) {
		return A_COMMAND
	}
	if regexp.MustCompile(`^\(.+\)$`).MatchString(cur) {
		return L_COMMAND
	}
	_, _, _, err := ParseCInstruction(cur)
	if err == nil {
		return C_COMMAND
	}
	log.Fatalf(`Unrecognized instruction: "%v"`, cur)
	return X_COMMAND
}

func (p Parser) symbol() (string, error) {
	cur := p.lines[p.index]

	// A instructions
	if regexp.MustCompile(`^@\d+$`).MatchString(cur) {
		return regexp.MustCompile(`^@`).ReplaceAllString(cur, ""), nil
	}
	if regexp.MustCompile(`^@[A-Za-z0-9_.$]+$`).MatchString(cur) {
		return regexp.MustCompile(`^@`).ReplaceAllString(cur, ""), nil
	}

	// labels, minus the parentheses
	if regexp.MustCompile(`^\(.+\)$`).MatchString(cur) {
		return strings.Trim(cur, "()"), nil
	}

	return "", fmt.Errorf("unrecognized symbol: %s", cur)
}

// Method destCompJump() is a thin wrapper to ParseCInstruction(). It returns a
// string containing a bitwise representation of the C instruction.
func (p Parser) destCompJump() (
	destbits string,
	compbits string,
	jumpbits string,
	err error) {
	cur := p.lines[p.index]

	// parse the C instructions into components
	dest, comp, jump, err := ParseCInstruction(cur)
	if err != nil {
		return "", "", "", fmt.Errorf(`Error parsing C instruction "%s": %s`, cur, err)
	}

	// translate the components into their bit representations
	var code = Code{}
	destbits, err = code.dest(dest)
	if err != nil {
		err = fmt.Errorf(`Error calculating dest bits, C instruction "%s": %s`, cur, err)
		return "", "", "", err
	}
	compbits, err = code.comp(comp)
	if err != nil {
		err = fmt.Errorf(`Error calculating comp bits, C instruction "%s": %s`, cur, err)
		return "", "", "", err
	}
	jumpbits, err = code.jump(jump)
	if err != nil {
		err = fmt.Errorf(`Error calculating jump bits, C instruction "%s": %s`, cur, err)
		return "", "", "", err
	}

	return destbits, compbits, jumpbits, nil
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

/*
Attempt to parse a string as a C instruction. C instructions look like
"dest=comp;jump", but either the "dest" or the "jump" parts can be empty.

In this function I assume that one of the "dest" or "jump" parts must be
present, since anything else is a NOOP and is probably an error. This is a
combination of the dest(), comp() and jump() methods as described in the N2T
book, page 114.

I am not in love with this logic, but it's good enough.
*/
func ParseCInstruction(inst string) (dest string, comp string, jump string, err error) {
	dest, comp, jump = "", "", ""
	err = nil
	if strings.Contains(inst, "=") {
		var tmp string
		dest, tmp, _ = strings.Cut(inst, "=")
		if strings.Contains(tmp, ";") {
			comp, jump, _ = strings.Cut(tmp, ";")
		} else {
			comp = tmp
		}
	} else if strings.Contains(inst, ";") {
		comp, jump, _ = strings.Cut(inst, ";")
	} else {
		err = errors.New("No = or ; found")
	}
	return
}
