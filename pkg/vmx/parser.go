package vmx

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/rs/zerolog/log"
)

func NewParser(file string) Parser {
	p := Parser{Source: file}
	fh, err := os.Open(p.Source)
	if err != nil {
		log.Fatal().Err(err)
	}
	p.Scanner = bufio.NewScanner(fh)
	p.Scanner.Split(bufio.ScanLines)
	return p
}

func (p *Parser) Scan() bool {
	return p.Scanner.Scan()
}

func (p *Parser) Text() string {
	return p.Scanner.Text()
}

//
func (p *Parser) commandType() Command {
	txt := p.Scanner.Text()

	if regexp.MustCompile(`^\s*(add|sub|neg|eq|gt)\s*$`).MatchString(txt) {
		return C_ARITHMETIC
	} else if regexp.MustCompile(`^\s*pop\s+`).MatchString(txt) {
		return C_POP
	} else if regexp.MustCompile(`^\s*push\s+`).MatchString(txt) {
		return C_PUSH
	}

	log.Fatal().Msg(txt)
	return C_UNDEFINED
}

func (p *Parser) arg1() string {
	return strings.Fields(p.Scanner.Text())[1]
}

func (p *Parser) arg2() int {
	txt := p.Scanner.Text()
	fields := strings.Fields(txt)
	arg2 := fields[2]
	j, _ := strconv.Atoi(arg2)
	return j
}
