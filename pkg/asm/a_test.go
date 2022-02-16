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
		st := asm.DefaultSymbolTable()
		assert.Equal(t, asm.FirstVariableAddress, st.Pointer)
		st, code, err := i.Assemble(st)
		assert.Nil(t, err)
		assert.Equal(t, []asm.MachineCode{tc.mc}, code)
		assert.Equal(t, asm.FirstVariableAddress, st.Pointer)
	}
}

// Check compiling of symbolic A instructions, default and user-defined.
func TestCompileA_symbol(t *testing.T) {
	st := asm.DefaultSymbolTable()
	st.Table[`foo`] = 1234

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
		assert.Equal(t, asm.FirstVariableAddress, st.Pointer)
		st, code, err := i.Assemble(st)
		assert.Nil(t, err)
		assert.Equal(t, []asm.MachineCode{tc.code}, code)
		assert.Equal(t, asm.FirstVariableAddress, st.Pointer)
	}
}

func TestCompileA_new_symbol(t *testing.T) {
	// compile in the standard Hack environment
	st := asm.DefaultSymbolTable()

	// the address of the first variable (symbol) is controlled by package
	// constant `FirstVariableAddress` (equal to 16 in the Hack standard).
	base := asm.FirstVariableAddress

	// sanity check: the symbol table pointer should point to the next
	// free address
	assert.Equal(t, asm.Address(base), st.Pointer)

	// assembling an A instruction that introduces a new symbol should
	// take a new slot at the start of available memory
	i := asm.AInstruction(`@foo`)
	st, code, err := i.Assemble(st)
	assert.Nil(t, err)
	assert.Equal(t, []asm.MachineCode{asm.MachineCode(base)}, code)
	assert.Equal(t, asm.Address(base), st.Table[`foo`])

	// the symbol table pointer should now point to the next free address
	assert.Equal(t, asm.Address(base+1), st.Pointer)

	// assembling another new symbol should take another new slot at
	// the start of available memory
	i = asm.AInstruction(`@bar`)
	st, code, err = i.Assemble(st)
	assert.Nil(t, err)
	assert.Equal(t, []asm.MachineCode{asm.MachineCode(base + 1)}, code)
	assert.Equal(t, asm.Address(base+1), st.Table[`bar`])

	// the symbol table pointer should point to the next free address
	assert.Equal(t, st.Pointer, base+2)

	// reuse of first symbol should return established address, and
	// not use a new memory slot
	i = asm.AInstruction(`@foo`)
	st, code, err = i.Assemble(st)
	assert.Nil(t, err)
	assert.Equal(t, []asm.MachineCode{asm.MachineCode(base)}, code)
	assert.Equal(t, asm.Address(base), st.Table[`foo`])

	// reuse of second symbol should return established address
	i = asm.AInstruction(`@bar`)
	st, code, err = i.Assemble(st)
	assert.Nil(t, err)
	assert.Equal(t, []asm.MachineCode{asm.MachineCode(base + 1)}, code)
	assert.Equal(t, asm.Address(base+1), st.Table[`bar`])

	// the symbol table pointer should still point to the next free address
	assert.Equal(t, st.Pointer, base+2)

}
