package asm_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/trammell/nand2tetris/pkg/asm"
)

func TestIsCInstuction(t *testing.T) {
	assert.True(t, asm.IsCInstruction(`M=1`))
	assert.False(t, asm.IsCInstruction(`// comment`))
	assert.False(t, asm.IsCInstruction(`@value`))
}

func TestAssembleC(t *testing.T) {
	tests := []struct {
		inst string
		code asm.MachineCode
	}{
		{inst: "0", code: 0b1110_1010_1000_0000}, // noop
		{inst: "M=1", code: 0b1110_1111_1100_1000},
		{inst: "M=0", code: 0b1110_1010_1000_1000},
		{inst: "0;JMP", code: 0b1110_1010_1000_0111},
		{inst: "MD=M+1", code: 0b1111_1101_1101_1000},
	}
	for _, tc := range tests {
		i := asm.CInstruction(tc.inst)
		st := asm.SymbolTable{}
		st, code, err := i.Assemble(st)
		assert.Nil(t, err)
		assert.Equal(t, []asm.MachineCode{tc.code}, code)
	}

}
