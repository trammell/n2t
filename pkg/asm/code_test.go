package asm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCodeString(t *testing.T) {
	var c Code = 100
	assert.Equal(t, `0000000001100100`, c.String())
	assert.Equal(t, int(0x64), int(c))
}
