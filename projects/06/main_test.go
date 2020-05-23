package main

import (
	"reflect"
	"testing"
)

func TestCanonicalizeInstruction(t *testing.T) {
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
}

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
	}

	for _, tc := range tests {
		i1 := NewInstruction(tc.in)
		got, _ := i1.Type()
		if !reflect.DeepEqual(tc.ty, got) {
			t.Errorf("instruction: %v, expected: %v, got: %v", tc.in, tc.ty, got)
		}
	}
}

// func TestIsAInstruction(t *testing.T) {
// 	i1 := Instruction{`@0`, 0}
// 	if !i1.IsAInstruction() {
// 		t.Errorf("@0 should be an A instruction")
// 	}

// 	i2 := Instruction{`(END_LT)`, 0}
// 	if i2.IsAInstruction() {
// 		t.Errorf("(END_LT) should be a label")
// 	}

// 	i3 := Instruction{`D;JNE`, 0}
// 	if i3.IsAInstruction() {
// 		t.Errorf("D;JNE should be a C instruction")
// 	}

// 	i4 := Instruction{`@ball.move$if_true0`, 0}
// 	if !i4.IsAInstruction() {
// 		t.Errorf("@ball.move$if_true0 should be a C instruction")
// 	}
// }

// func TestIsCInstruction(t *testing.T) {

// 	type test struct {
// 		Inst string
// 		IsC  bool
// 	}

// 	tests := []test{
// 		{Inst: "0;JMP", IsC: true},
// 		{Inst: "@0", IsC: false},
// 		{Inst: "A=M", IsC: true},
// 		{Inst: "D=D-1;JEQ", IsC: true},
// 		{Inst: "(LABEL.ONE)", IsC: false},
// 	}

// 	for _, tc := range tests {
// 		i1 := Instruction{tc.Inst, 0}
// 		got := i1.IsCInstruction()
// 		if !reflect.DeepEqual(tc.IsC, got) {
// 			t.Errorf("instruction: %v, expected: %v, got: %v", tc.Inst, tc.IsC, got)
// 		}
// 	}
// }
