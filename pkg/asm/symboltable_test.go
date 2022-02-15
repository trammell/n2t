package asm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSymbolTableClone(t *testing.T) {
	st1 := SymbolTable{"foo": 100, "bar": 200}
	st2 := st1.Clone()
	assert.Equal(t, st1, st2)
}
