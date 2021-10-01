package asm

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestNewInstruction(t *testing.T) {
	src := "@256"
	assert.Equal(t, NewInstruction(src).Text, src, `Instruction cleaner`)
}
