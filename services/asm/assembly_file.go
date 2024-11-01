package main

import (
	"fmt"

	"github.com/rs/zerolog/log"
)

// Construct a new Assembler object. The first 16 RAM addresses are reserved.
func NewAssemblyFile(filename string) Assembler {
	return AssemblyFile{
		Source:      *NewSource(filename),
		SymbolTable: DefaultSymbolTable(),
	}
}

// Implement the Stringer interface
func (a Assembler) String() (out string) {
	out = a.SymbolTable.String()
	for _, inst := range a.Instructions {
		out += fmt.Sprintln(inst)
	}
	return
}

// Assemble the instructions into machine code
func (a *Assembler) Assemble() {
	a.ReadInstructions()
	a.FillSymbolTable()
	a.AssembleInstructions()
}

// Read all the instructions, stash them in the Assembler object.
func (a *Assembler) ReadInstructions() {
	for i, more := a.Source.GetInstruction(); more; i, more = a.Source.GetInstruction() {
		if IsAInstruction(Instruction(i)) {
			a.Instructions = append(a.Instructions, AInstruction(i))
		} else if IsCInstruction(Instruction(i)) {
			a.Instructions = append(a.Instructions, CInstruction(i))
		} else if IsLabel(Instruction(i)) {
			a.Instructions = append(a.Instructions, Label(i))
		} else {
			log.Warn().Msgf(`unrecognized instruction: %v`, i)
		}
	}
	return
}

// Fill the symbol table with all the Label instructions.
func (a *Assembler) FillSymbolTable() {
	addr := Address(0)
	for _, inst := range a.Instructions {
		// A and C instructions increment address, labels do not
		a.SymbolTable, addr = inst.UpdateSymbolTable(a.SymbolTable, addr)
	}
}

// Assemble all A and C instructions into machine code
func (a *Assembler) AssembleInstructions() {
	for _, inst := range a.Instructions {
		st, mc, err := inst.Assemble(a.SymbolTable)
		a.SymbolTable = st
		if err != nil {
			log.Fatal().Err(err)
		}
		a.MachineCode = append(a.MachineCode, mc...)
	}
}

// Return a copy of the generated machine code. Returning a copy makes
// testing easier.
func (a *Assembler) GetMachineCode() (mc []MachineCode) {
	mc = make([]MachineCode, len(a.MachineCode))
	copy(mc, a.MachineCode)
	return
}
