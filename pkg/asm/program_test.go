package asm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
		assert.Equal(t, IsInstruction(tc.txt), tc.isInst)
	}
}

// // Compile Max.asm and compare output to Max.hack
// func TestAssembleMaxAsm(t *testing.T) {

// 	// open and assemble the source file
// 	p := Program{Filename: "testdata/Max.asm"}
// 	a := Assembler{Program: p}
// 	got := a.assemble()

// 	// open and read the reference file
// 	ref, err := os.Open(`testdata/Max.hack`)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer ref.Close()

// 	// compare the file contents
// 	lineno := 1
// 	scanner := bufio.NewScanner(ref)
// 	for scanner.Scan() {
// 		want := scanner.Text()
// 		if len(got) == 0 {
// 			t.Fatalf(`Max.hack line %d: ran out of instructions`, lineno)
// 		}
// 		if want != got[0].String() {
// 			t.Fatalf(`Max.hack line %d: want '%s', got '%s'`, lineno, want, got[0].String())
// 		} else {
// 			got = got[1:]
// 			lineno++
// 		}
// 	}
// 	if err := scanner.Err(); err != nil {
// 		log.Fatal(err)
// 	}
// }

// // Compile Max.asm and compare output to Max.hack
// func TestAssembleExamples(t *testing.T) {
// 	tests := []struct {
// 		asm  string
// 		hack string
// 	}{
// 		{asm: "Add.asm", hack: "Add.hack"},
// 		{asm: "Max.asm", hack: "Max.hack"},
// 		{asm: "MaxL.asm", hack: "MaxL.hack"},
// 		{asm: "Rect.asm", hack: "Rect.hack"},
// 		{asm: "RectL.asm", hack: "RectL.hack"},
// 	}

// 	for _, tc := range tests {

// 		// read and assemble the *.asm source
// 		p := NewProgram(`testdata/` + tc.asm)
// 		p.Read()
// 		p.ResolveSymbols()
// 		got := p.EmitToSlice()

// 		// Open the reference file
// 		ref, err := os.Open(`testdata/` + tc.hack)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		defer ref.Close()

// 		// load the file contents & compare
// 		lineno := 1
// 		scanner := bufio.NewScanner(ref)
// 		for scanner.Scan() {
// 			want := scanner.Text()
// 			if len(got) == 0 {
// 				t.Fatalf(`assemble failed, file %s, line %d: ran out of instructions`, tc.asm, lineno)
// 			}
// 			if want != got[0].String() {
// 				t.Fatalf(`assemble failed, file %s, line %d: want '%s', got '%s'`, tc.asm, lineno, want, got[0].String())
// 			} else {
// 				got = got[1:]
// 				lineno++
// 			}
// 		}
// 		if err := scanner.Err(); err != nil {
// 			log.Fatal(err)
// 		}
// 	}
// }
