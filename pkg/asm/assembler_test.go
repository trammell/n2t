package asm_test

import (
	"bufio"
	"os"
	"testing"

	"github.com/rs/zerolog/log"
	"github.com/trammell/nand2tetris/pkg/asm"
)

// Compile Max.asm and compare output to Max.hack
func TestAssembleMaxAsm(t *testing.T) {

	// Create the assembler object, assemble the source file, and fetch
	// the generated machine code.
	a := asm.NewAssembler("testdata/Max.asm")
	a.Assemble()
	got := a.GetMachineCode()

	// open and read the reference file
	ref, err := os.Open(`testdata/Max.hack`)
	if err != nil {
		log.Fatal()
	}
	defer ref.Close()

	// compare the file contents line by line
	lineno := 1
	scanner := bufio.NewScanner(ref)
	for scanner.Scan() {
		want := scanner.Text()
		if len(a.MachineCode) == 0 {
			t.Fatalf(`Max.hack line %d: ran out of instructions`, lineno)
		}
		if want != got[0].String() {
			t.Fatalf(`Max.hack line %d: want '%s', got '%s'`, lineno, want, got[0].String())
		} else {
			got = got[1:]
			lineno++
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal()
	}
}

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
