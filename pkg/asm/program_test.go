package asm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCleanUp(t *testing.T) {
	tests := []struct {
		line string
		want string
	}{
		{line: "@256", want: "@256"},
		{line: "D = A // assign accum.", want: "D=A"},
		{line: "// just a comment", want: ""},
		{line: " ", want: ""},
		{line: "", want: ""},
	}
	for _, tc := range tests {
		assert.Equal(t, CleanUp(tc.line), Instruction(tc.want))
	}
}
