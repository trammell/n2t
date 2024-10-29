package asm_test

import (
	"bufio"
	"os"
	"testing"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
	"github.com/trammell/n2t/pkg/asm"
)

// Compile Max.asm and compare output to Max.hack
func TestAssembleMaxAsm(t *testing.T) {
	zerolog.SetGlobalLevel(zerolog.Disabled)

	// Create the assembler object, assemble the source file, and fetch
	// the generated machine code.
	a := asm.NewAssembler("testdata/Max.asm")

	// there should be 19 instructions (3 labels and 16 assembly)
	a.ReadInstructions()
	assert.Len(t, a.Instructions, 19)

	// make sure labels are added to the symbol table
	a.FillSymbolTable()
	assert.Equal(t, asm.Address(0xe), a.SymbolTable.Table[`INFINITE_LOOP`])
	assert.Equal(t, asm.Address(0xc), a.SymbolTable.Table[`OUTPUT_D`])
	assert.Equal(t, asm.Address(0xa), a.SymbolTable.Table[`OUTPUT_FIRST`])
	assert.Contains(t, a.String(), `pointer=16`)

	// Check the assembled instructions
	a.AssembleInstructions()
	got := a.GetMachineCode()

	assert.Equal(t, 16, len(got))
	assert.Equal(t, `0000000000000000`, got[0].String())
	assert.Equal(t, `1111110000010000`, got[1].String())
	assert.Equal(t, `0000000000000001`, got[2].String())

	// 1111010011010000
	// 0000000000001010
	// 1110001100000001
	// 0000000000000001
	// 1111110000010000
	// 0000000000001100
	// 1110101010000111
	// 0000000000000000
	// 1111110000010000
	// 0000000000000010
	// 1110001100001000
	// 0000000000001110
	// 1110101010000111

}

// Compile all test .asm files and compare to reference
func TestAssembleAllExamples(t *testing.T) {
	tests := []struct {
		asm  string
		hack string
	}{
		{asm: "Add.asm", hack: "Add.hack"},
		{asm: "Max.asm", hack: "Max.hack"},
		{asm: "MaxL.asm", hack: "MaxL.hack"},
		{asm: "Rect.asm", hack: "Rect.hack"},
		{asm: "RectL.asm", hack: "RectL.hack"},
	}

	for _, tc := range tests {

		// read and assemble the *.asm source
		a := asm.NewAssembler(`testdata/` + tc.asm)
		a.Assemble()
		got := a.GetMachineCode()

		// Open the reference file
		ref, err := os.Open(`testdata/` + tc.hack)
		if err != nil {
			log.Fatal().Err(err)
		}
		defer ref.Close()

		// load the file contents & compare
		lineno := 1
		scanner := bufio.NewScanner(ref)
		for scanner.Scan() {
			want := scanner.Text()
			if len(got) == 0 {
				t.Fatalf(`assemble failed, file %s, line %d: ran out of instructions`, tc.asm, lineno)
			}
			if want != got[0].String() {
				t.Fatalf(`assemble failed, file %s, line %d: want '%s', got '%s'`, tc.asm, lineno, want, got[0].String())
			} else {
				got = got[1:]
				lineno++
			}
		}
		if err := scanner.Err(); err != nil {
			log.Fatal().Err(err)
		}
	}
}
