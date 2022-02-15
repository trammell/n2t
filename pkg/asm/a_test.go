package asm_test

import (
	"os"
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"github.com/trammell/nand2tetris/pkg/asm"
)

func TestMain(m *testing.M) {
	zerolog.SetGlobalLevel(zerolog.Disabled)

	os.Exit(m.Run())
}

func TestIsAInstuction(t *testing.T) {
	assert.True(t, asm.IsAInstruction(`@value`))
	assert.True(t, asm.IsAInstruction(`@99`))
	assert.True(t, asm.IsAInstruction(`@value`))
	assert.False(t, asm.IsAInstruction(`M=1`))
	assert.False(t, asm.IsAInstruction(`// comment`))
}

func TestCompileA_const(t *testing.T) {

	tests := []struct {
		inst string
		mc   asm.MachineCode
	}{
		{inst: "@0", mc: 0},
		{inst: "@1", mc: 1},
		{inst: "@5", mc: 5},
		{inst: "@100", mc: 100},
	}
	for _, tc := range tests {
		i := asm.AInstruction(tc.inst)
		st := asm.SymbolTable{}
		code, err := i.Assemble(st)
		assert.Nil(t, err)
		assert.Equal(t, []asm.MachineCode{tc.mc}, code)
	}
}

func TestCompileA_symbol(t *testing.T) {
	st := asm.PredefinedSymbols.Clone()
	st[`foo`] = 1234

	tests := []struct {
		inst string
		code asm.MachineCode
	}{
		{inst: "@R0", code: 0},
		{inst: "@R1", code: 1},
		{inst: "@SCREEN", code: 16384},
		{inst: "@foo", code: 1234},
	}
	for _, tc := range tests {
		i := asm.AInstruction(tc.inst)
		code, err := i.Assemble(st)
		assert.Nil(t, err)
		assert.Equal(t, []asm.MachineCode{tc.code}, code)
	}
}

func TestCompileA_newsym(t *testing.T) {
	st := asm.PredefinedSymbols.Clone()

	// the address of the first variable (symbol) is controlled by constant
	// `FirstVariableAddress` (no magic numbers here!)
	first := asm.FirstVariableAddress
	second := first + 1

	// first new symbol should take a new slot at the start of available memory
	i := asm.AInstruction(`@foo`)
	code, err := i.Assemble(st)
	assert.Nil(t, err)
	assert.Equal(t, []asm.MachineCode{asm.MachineCode(first)}, code)
	assert.Equal(t, asm.Address(first), st[`foo`])

	// second new symbol should take a new slot at the start of available memory
	i = asm.AInstruction(`@bar`)
	code, err = i.Assemble(st)
	assert.Nil(t, err)
	assert.Equal(t, []asm.MachineCode{asm.MachineCode(second)}, code)
	assert.Equal(t, asm.Address(second), st[`bar`])

	// reuse of first symbol should return established address
	i = asm.AInstruction(`@foo`)
	code, err = i.Assemble(st)
	assert.Nil(t, err)
	assert.Equal(t, []asm.MachineCode{asm.MachineCode(first)}, code)
	assert.Equal(t, asm.Address(first), st[`foo`])

	// reuse of second symbol should return established address
	i = asm.AInstruction(`@bar`)
	code, err = i.Assemble(st)
	assert.Nil(t, err)
	assert.Equal(t, []asm.MachineCode{asm.MachineCode(second)}, code)
	assert.Equal(t, asm.Address(second), st[`bar`])
}
