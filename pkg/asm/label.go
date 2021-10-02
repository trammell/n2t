package asm

import (
	"regexp"
)

// Return True if this is a label
func isLabel(txt string) bool {
	return regexp.MustCompile(`^\(.+\)$`).MatchString(txt)
}
