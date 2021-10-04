package asm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsAInstuction(t *testing.T) {
	assert.True(t, isAInstruction(`@value`))
	assert.True(t, isAInstruction(`@99`))
	assert.True(t, isAInstruction(`@value`))
	assert.False(t, isAInstruction(`M=1`))
	assert.False(t, isAInstruction(`// comment`))
}

func TestCompileA_const(t *testing.T) {
	tests := []struct {
		inst string
		code Code
	}{
		{inst: "@0", code: 0b0000000000000000},
		{inst: "@1", code: 0b0000000000000001},
		{inst: "@5", code: 0b0000000000000101},
		{inst: "@100", code: 0b0000000001100100},
	}
	for _, tc := range tests {
		i := AInstruction(tc.inst)
		st := SymbolTable{}
		code, err := i.Assemble(st)
		assert.Nil(t, err)
		assert.Equal(t, tc.code, code)
	}
}
