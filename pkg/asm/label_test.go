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
