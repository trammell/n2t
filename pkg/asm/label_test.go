package asm_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/trammell/nand2tetris/pkg/asm"
)

func TestIsLabel(t *testing.T) {
	assert.True(t, asm.IsLabel(`(END)`))
	assert.False(t, asm.IsLabel(`// comment`))
	assert.False(t, asm.IsLabel(`@value`))
	assert.False(t, asm.IsLabel(`MD=M+1`))
}

func TestAssembleLabel(t *testing.T) {
	tests := []struct {
		inst string
		addr int
		addx int
	}{
		{inst: "(foo)", addr: 0, addx: 1}, // noop
		{inst: "(bar)", addr: 123, addx: 124},
		{inst: "(x)", addr: 9999, addx: 1000},
	}
	for _, tc := range tests {
		i := asm.Label(tc.inst)
		st := asm.SymbolTable{}
		st, code, err := i.Assemble(st)
		assert.Nil(t, err)
		assert.Equal(t, []asm.MachineCode{}, code)
	}
}

func TestLabelResolve(t *testing.T) {
	label := asm.Label(`(foo)`)
	st := asm.SymbolTable{Table: map[asm.Symbol]asm.Address{}}

	// initial address resolution: 1234
	st, addr := label.UpdateSymbolTable(st, 1234)
	assert.Equal(t, asm.Address(1234), addr)
}
