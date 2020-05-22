package main

import (
	"reflect"
	"testing"
)

func TestCanonicalizeInstruction(t *testing.T) {
	i1 := new(inst)
	var x string
	if x = i1.Canonicalize("// this is just a comment"); x != "" {
		t.Errorf("got '%s'; want ''", x)
	}
	if x = i1.Canonicalize("D = M; JMP // more comments"); x != "D=M;JMP" {
		t.Errorf("got '%s'; want 'D=M;JMP'", x)
	}
	if x = i1.Canonicalize("@1234	// another // comment"); x != "@1234" {
		t.Errorf("got '%s'; want '@1234'", x)
	}
}

func TestIsLabel(t *testing.T) {
	i1 := inst{`(label)`, 0}
	if !i1.IsLabel() {
		t.Errorf("(label) should be a label")
	}

	i2 := inst{`(label.foo)`, 0}
	if !i2.IsLabel() {
		t.Errorf("(label.foo) should be a label")
	}

	i3 := inst{`(LABEL.TWO)`, 0}
	if !i3.IsLabel() {
		t.Errorf("(LABEL.TWO) should be a label")
	}

	i4 := inst{`@FOO`, 0}
	if i4.IsLabel() {
		t.Errorf("@FOO should not be a label")
	}
}

func TestIsAInstruction(t *testing.T) {
	i1 := inst{`@0`, 0}
	if !i1.IsAInstruction() {
		t.Errorf("@0 should be an A instruction")
	}

	i2 := inst{`(END_LT)`, 0}
	if i2.IsAInstruction() {
		t.Errorf("(END_LT) should be a label")
	}

	i3 := inst{`D;JNE`, 0}
	if i3.IsAInstruction() {
		t.Errorf("D;JNE should be a C instruction")
	}
}

func TestIsCInstruction(t *testing.T) {

	type test struct {
		Inst string
		IsC  bool
	}

	tests := []test{
		{Inst: "0;JMP", IsC: true},
		{Inst: "@0", IsC: false},
		{Inst: "A=M", IsC: true},
		{Inst: "D=D-1;JEQ", IsC: true},
		{Inst: "(LABEL.ONE)", IsC: false},
	}

	for _, tc := range tests {
		i1 := inst{tc.Inst, 0}
		got := i1.IsCInstruction()
		if !reflect.DeepEqual(tc.IsC, got) {
			t.Errorf("instruction: %v, expected: %v, got: %v", tc.Inst, tc.IsC, got)
		}
	}
}
