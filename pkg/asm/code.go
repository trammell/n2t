package asm

import "fmt"

// implement the stringer interface
func (c Code) String() string {
	return fmt.Sprintf("%016b", uint16(c))
}
