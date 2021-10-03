package asm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsAInstuction(t *testing.T) {
	assert.True(t, isAInstruction(`@value`))
	assert.False(t, isAInstruction(`M=1`))
	assert.False(t, isAInstruction(`// comment`))
}