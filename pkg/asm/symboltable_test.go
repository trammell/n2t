package asm_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/trammell/n2t/pkg/asm"
)

func TestSymbolTableClone(t *testing.T) {
	st1 := asm.SymbolTable{
		Pointer: 99,
		Table:   map[asm.Symbol]asm.Address{"foo": 100, "bar": 200},
	}
	st2 := st1.Clone()
	assert.Equal(t, st1, st2)
}

func TestSymbolTableStringer(t *testing.T) {
	st1 := asm.SymbolTable{
		Pointer: 1,
		Table:   map[asm.Symbol]asm.Address{"foo": 100, "bar": 200},
	}
	assert.Equal(t, "pointer=1\nst[bar]=200\nst[foo]=100\n", st1.String())
}
