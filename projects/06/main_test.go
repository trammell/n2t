package main

import (
	"testing"
)


func TestCanonicalizeInstruction(t *testing.T) {
	if x := canonicalizeInstruction("// this is just a comment"); x != "" {
		t.Errorf("got '%s'; want ''", x)
	}
}
