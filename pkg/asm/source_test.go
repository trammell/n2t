package asm_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/trammell/n2t/pkg/asm"
)

func TestIsComment(t *testing.T) {
	tests := []struct {
		txt       string
		isComment bool
	}{
		{txt: "@256", isComment: false},
		{txt: "D = A // assign accum.", isComment: false},
		{txt: "// just a comment", isComment: true},
		{txt: " ", isComment: false},
		{txt: "", isComment: false},
	}
	for _, tc := range tests {
		assert.Equal(t, asm.IsComment(tc.txt), tc.isComment)
	}
}

func TestIsInstruction(t *testing.T) {
	tests := []struct {
		txt    string
		isInst bool
	}{
		{txt: "@256", isInst: true},
		{txt: "D = A // assign accum.", isInst: true},
		{txt: "// just a comment", isInst: false},
		{txt: "(symbol) // a symbol", isInst: true},
		{txt: " ", isInst: false},
		{txt: "", isInst: false},
	}
	for _, tc := range tests {
		assert.Equal(t, asm.IsInstruction(tc.txt), tc.isInst)
	}
}
