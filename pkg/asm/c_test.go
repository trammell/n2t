package asm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsCInstuction(t *testing.T) {
	assert.True(t, isCInstruction(`M=1`))
	assert.False(t, isCInstruction(`// comment`))
	assert.False(t, isCInstruction(`@value`))
}