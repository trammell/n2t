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
		{inst: "@0", code: 0},
		{inst: "@1", code: 1},
		{inst: "@5", code: 5},
		{inst: "@100", code: 100},
	}
	for _, tc := range tests {
		i := AInstruction(tc.inst)
		st := SymbolTable{}
		code, err := i.Assemble(st)
		assert.Nil(t, err)
		assert.Equal(t, tc.code, code)
	}
}

func TestCompileA_symbol(t *testing.T) {
	st := DefaultSymbols.clone()
	st[`foo`] = 1234

	tests := []struct {
		inst string
		code Code
	}{
		{inst: "@R0", code: 0},
		{inst: "@R1", code: 1},
		{inst: "@SCREEN", code: 16384},
		{inst: "@foo", code: 1234},
	}
	for _, tc := range tests {
		i := AInstruction(tc.inst)
		code, err := i.Assemble(st)
		assert.Nil(t, err)
		assert.Equal(t, tc.code, code)
	}
}


