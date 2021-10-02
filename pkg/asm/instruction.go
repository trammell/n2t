package asm

// NewInstruction is an Instruction constructor
func NewInstruction(txt string) *InstructionAssembler {
	if isAInstruction(txt) {
		return NewAInstruction(txt)
	} else if isC(txt) {
		return NewCInstruction(txt)
	} else if isLabel(txt) {

	} else {

	}
	i := new(Instruction)
	i.Text = CleanUp(txt)
	return i
}
