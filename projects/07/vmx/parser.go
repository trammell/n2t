package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func NewParser(filename string) (*Parser, error) {
	p := new(Parser)
	p.FileName = filename
	lines, err := readLines(filename)
	if err != nil {
		return nil, err
	}
	p.Lines = lines
	p.Index = -1
	return p, nil
}

func (p Parser) currentCommand() string {
	return p.Lines[p.Index]
}

func (p Parser) hasMoreCommands() bool {
	return p.Index < len(p.Lines) - 1
}

func (p *Parser) advance() {
	if p.hasMoreCommands() {
		p.Index += 1
	}
}

func (p *Parser) reset() {
	p.Index = -1
}

func (p Parser) commandType() (uint8, error) {
	cur := p.Lines[p.Index]
	if regexp.MustCompile(`^(add|sub|neg|eq|gt)\s+$`).MatchString(cur) {
		return C_ARITHMETIC, nil
	} else if regexp.MustCompile(`^pop\s+`).MatchString(cur) {
		return C_POP, nil
	} else if regexp.MustCompile(`^push\s+`).MatchString(cur) {
		return C_PUSH, nil
	}
	return C_UNDEFINED, fmt.Errorf("Unrecognized command type: %s", cur)
}

func (p *Parser) arg1() string {
	cur := p.Lines[p.Index]
	return strings.Fields(cur)[1]
}

func (p *Parser) arg2() int {
	cur := p.Lines[p.Index]
	fields := strings.Fields(cur)
	arg2 := fields[2]
	j, _ := strconv.Atoi(arg2)
	return j
}

// read source file into a slice of strings
func readLines(src string) ([]string, error) {
	file, err := os.Open(src)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		txt := strings.TrimSpace(scanner.Text())
		if txt != "" {
			lines = append(lines, txt)
		}
	}
	return lines, scanner.Err()
}
