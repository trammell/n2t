package asm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCodeString(t *testing.T) {
	var mc MachineCode = 100
	assert.Equal(t, `0000000001100100`, mc.String())
	assert.Equal(t, int(0x64), int(mc))
}
