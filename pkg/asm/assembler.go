package asm

import (
	"fmt"

	"github.com/rs/zerolog/log"
)

func NewAssembler(filename string, cchan chan MachineCode) Assembler {
	sc := SourceCode{FileName: filename}
	return Assembler{SourceCode: sc, MachineCode: cchan,
		SymbolTable: PredefinedSymbols.Clone()}
}

func (a Assembler) String() string {
	return `foo` // FIXME, this should return a string representation of the assembler
}

func (a *Assembler) Assemble() []MachineCode {
	a.SourceCode.GetInstruction()

	return []MachineCode{}
}

// resolve all unresolved symbols in instructions
func (a *Assembler) ResolveSymbols() {
	addr := Address(0)
	for _, inst := range a.Instructions {
		// A and C instructions increment address, labels do not
		addr = inst.UpdateSymbolTable(a.SymbolTable, addr)
	}
}

// emit all instructions, as machine code, to STDOUT
func (a *Assembler) EmitToStdout() {
	gen := a.AssemblyGenerator()
	for codes, cont := gen(); cont; codes, cont = gen() {
		for _, x := range codes {
			fmt.Println(x.String())
		}
	}
}

// Emit all instructions, compiled to machine code, in a slice of MachineCodes.
// This is useful in testing.
func (a *Assembler) EmitToSlice() []MachineCode {
	gen := a.AssemblyGenerator()
	var out []MachineCode
	for s, x := gen(); x; s, x = gen() {
		out = append(out, s...)
	}
	return out
}

// This function returns a generator object
func (a *Assembler) AssemblyGenerator() func() (codes []MachineCode, more bool) {

	// make a copy of p.Instructions, and then just keep shifting them off
	instructions := make([]Assemblable, len(a.Instructions))
	copy(instructions, a.Instructions)

	// return a closure that shifts off & assembles the first slice elt
	return func() ([]MachineCode, bool) {
		if len(instructions) == 0 {
			return []MachineCode{}, false // no value & no more instructions
		}

		// shift off the first instruction & assemble it
		i := instructions[0]
		instructions = instructions[1:]
		codes, err := i.Assemble(a.SymbolTable)
		if err != nil {
			log.Fatal()
		}
		return codes, true // continue
	}
}
