package vmx

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/rs/zerolog/log"
)

func NewCommand(src string) Command {
	return Command{vmCommand: src, fields: strings.Fields(src)}
}

func (c *Command) GetAsm() string {

	// weed out comments & blank lines
	if regexp.MustCompile(`^\s*(//.*)?$`).MatchString(c.vmCommand) {
		return ``
	}

	// arithmetic commands
	if regexp.MustCompile(`^\s*(add|sub|neg|eq|gt)\s*$`).MatchString(c.vmCommand) {
		return c.Math()
	}

	// push
	if c.fields[0] == "push" {
		return c.Push()
	}

	// pop
	if c.vmCommand == "pop" {
		return c.Pop()
	}

	// catchall
	return fmt.Sprintf("ERROR: I don't recognize VM command '%s'\n", c.vmCommand)
}

func (c *Command) Math() string {
	var asm string
	switch c.fields[0] {
	case "add":
		asm = Trim(`
			// add
			@SP     // A = 0
			M=M-1   // SP--
			A=M     // A = SP
			D=M     // D = RAM[SP] = *SP
			A=A-1   // A-- 		
			M=D+M   // RAM[SP-1] += D`)
	case "sub":
		return "// sub\n@SP\nD=M\nM=M-1\nD=D+M\nM=D\n\n"
	default:
		asm = "ERROR"
	}
	return Trim(asm)
}

func (c *Command) Push() string {
	segment := c.fields[1]
	switch segment {
	case `constant`:
		return fmt.Sprintf("// %s\n", c.vmCommand) +
			fmt.Sprintf("@%s\n", c.fields[2]) +
			"M=A\n@SP\nM=D\n@SP\nM=M+1\n\n"
	default:
		log.Fatal().Msgf(`Unrecognized segment: %s`, segment)
	}
	return `ERROR\n`
}

func (c *Command) Pop() string {
	return `ERROR\n`
}

// trim leading whitespace on each line in a multiline string
func Trim(str string) string {
	var out string
	for _, element := range strings.Split(str, "\n") {
		out += strings.TrimSpace(element) + "\n"
	}
	return out
}
