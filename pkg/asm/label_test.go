package asm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsLabel(t *testing.T) {
	assert.True(t, isLabel(`(END)`))
	assert.False(t, isLabel(`// comment`))
	assert.False(t, isLabel(`@value`))
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
		i := Label(tc.inst)
		st := SymbolTable{}
		code, err := i.Assemble(st)
		assert.Nil(t, err)
		assert.Equal(t, []MachineCode{}, code)
	}
}

func TestLabelResolve(t *testing.T) {
	label := Label(`(foo)`)
	st := SymbolTable{}

	// initial address resolution: 1234
	addr := label.Resolve(st, 1234)
	assert.Equal(t, Address(1234), addr)

}
