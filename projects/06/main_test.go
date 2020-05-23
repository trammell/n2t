package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestCleanUpInstruction(t *testing.T) {
	i1 := new(Instruction)
	var x string
	if x = i1.CleanUp("// this is just a comment"); x != "" {
		t.Errorf("got '%s'; want ''", x)
	}
	if x = i1.CleanUp("D = M; JMP // more comments"); x != "D=M;JMP" {
		t.Errorf("got '%s'; want 'D=M;JMP'", x)
	}
	if x = i1.CleanUp("@1234	// another // comment"); x != "@1234" {
		t.Errorf("got '%s'; want '@1234'", x)
	}
	if x = i1.CleanUp("DA = M; JNE"); x != "DA=M;JNE" {
		t.Errorf("got '%s'; want 'AD=M;JNE'", x)
	}
}

// Test our instruction type identifier
func TestInstructionType(t *testing.T) {

	// list sample instructions and their correct type
	var tests = []struct {
		in string // instruction
		ty int    // type
	}{
		{`(label)`, L},
		{`0;JMP`, C},
		{`@0`, A},
		{`@1234`, A},
		{`@FOO`, A},
		{`@bar`, A},
		{"A=M", C},
		{"D=D-1;JEQ", C},
		{"(LABEL.ONE)", L},
		{`(label.foo)`, L},
		{`D;JNE`, C},
		{`@ball.move$if_true0`, A},
	}

	for _, tc := range tests {
		i1 := NewInstruction(tc.in)
		got, _ := i1.Type()
		if !reflect.DeepEqual(tc.ty, got) {
			t.Errorf("instruction: %v, expected: %v, got: %v", tc.in, tc.ty, got)
		}
	}
}

// Test our A instruction assembler
func TestAInstructionAssembler(t *testing.T) {
	// list of instructions and their correct encoding
	var tests = []struct {
		input string // instruction
		want  string // correct encoding
	}{
		{`@0`, `0000000000000000`},
		{`@1`, `0000000000000001`},
		{`@2`, `0000000000000010`},
		{`@1234`, `0000010011010010`},
		{`@FOO`, `0000000000000111`},
		{`@ball.move$if_true0`, `0000000000111111`},
		{`@bar`, `0000000011111111`},
	}

	// the symbol table used for testing
	symbols := map[string]int{
		"FOO":                7,
		`ball.move$if_true0`: 63,
		`bar`:                255,
	}

	for _, tc := range tests {
		i1 := NewInstruction(tc.input)
		got, _ := i1.Assemble(symbols)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("instruction: %v, expected: %v, got: %v", tc.input, tc.want, got)
		}
	}

	// verify that assembly fails for missing symbols
	i2 := NewInstruction(`@missing`)
	_, err := i2.Assemble(symbols)
	if err == nil {
		t.Errorf("missing symbol should throw error")
	}

}

// Test our C instruction assembler
func TestCInstructionAssembler(t *testing.T) {
	// helper function to strip whitespace from a string
	s := func(x string) string {
		return strings.Replace(x, " ", "", -1)
	}

	// list of instructions and their correct encoding
	var tests = []struct {
		input string // instruction
		want  string // correct encoding
	}{
		{`0;JMP`, s(`111 0 000000 000 000`)},
		// {"A=M", C},
		// {"D=D-1;JEQ", C},
		// {`D;JNE`, C},
	}

	for _, tc := range tests {
		i1 := NewInstruction(tc.input)
		got, _ := i1.Assemble(nil)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("instruction: %v, expected: %v, got: %v", tc.input, tc.want, got)
		}
	}
}
