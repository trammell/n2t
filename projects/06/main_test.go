package main

import (
	"reflect"
	"testing"
)

func TestCanonicalizeInstruction(t *testing.T) {
	i1 := new(inst)
	if x := i1.Canonicalize("// this is just a comment"); x != "" {
		t.Errorf("got '%s'; want ''", x)
	}
}

func TestIsAInstruction(t *testing.T) {
	i1 := inst{"@0", 0}
	if !i1.IsAInstruction() {
		t.Errorf("misidentified A instruction")
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
		//A=M
	}

	for _, tc := range tests {
		i1 := inst{tc.Inst, 0}
		got := i1.IsCInstruction()
		if !reflect.DeepEqual(tc.IsC, got) {
			t.Errorf("instruction: %v, expected: %v, got: %v", tc.Inst, tc.IsC, got)
		}
	}

}
