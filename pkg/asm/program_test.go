package asm

import (
	"bufio"
	"log"
	"os"
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
		{line: "(symbol) // a symbol", want: "(symbol)"},
		{line: " ", want: ""},
		{line: "", want: ""},
	}
	for _, tc := range tests {
		assert.Equal(t, CleanUp(tc.line), Instruction(tc.want))
	}
}

// Compile Max.asm and compare output to Max.hack
func TestAssembleMax(t *testing.T) {
	p := NewProgram(`testdata/Max.asm`)
	p.Read()
	p.ResolveSymbols()
	got := p.EmitToSlice()

	// Open the reference file
	ref, err := os.Open(`testdata/Max.hack`)
	if err != nil {
		log.Fatal(err)
	}
	defer ref.Close()

	// load the file contents
	lineno := 1
	scanner := bufio.NewScanner(ref)
	for scanner.Scan() {
		want := scanner.Text()
		if len(got) == 0 {
			t.Fatalf(`assemble failed, line %d: ran out of instructions`, lineno)
		}
		if want != got[0] {
			t.Fatalf(`assemble failed, line %d: want %s, got %s`, lineno, want, got[0])
		} else {
			got = got[1:]
			lineno++
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
