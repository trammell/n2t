package asm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsCInstuction(t *testing.T) {
	assert.True(t, isCInstruction(`M=1`))
	assert.False(t, isCInstruction(`// comment`))
	assert.False(t, isCInstruction(`@value`))
}

func TestAssembleC(t *testing.T) {
	tests := []struct {
		inst string
		code Code
	}{
		{inst: "0", code: 0b1110_1010_1000_0000}, // noop
		{inst: "M=1", code: 0b1110_1111_1100_1000},
		{inst: "M=0", code: 0b1110_1010_1000_1000},
		{inst: "0;JMP", code: 0b1110_1010_1000_0111},
		{inst: "MD=M+1", code: 0b1111_1101_1101_1000},
	}
	for _, tc := range tests {
		i := CInstruction(tc.inst)
		st := SymbolTable{}
		code, err := i.Assemble(st)
		assert.Nil(t, err)
		assert.Equal(t, []Code{tc.code}, code)
	}

}
