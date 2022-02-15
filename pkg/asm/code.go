package asm

import "fmt"

// implement the stringer interface
func (mc MachineCode) String() string {
	return fmt.Sprintf("%016b", uint16(mc))
}
