package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/trammell/nand2tetris/pkg/asm"
)

func asmFn(cmd *cobra.Command, args []string) {
	code := make(chan asm.MachineCode)
	a := asm.NewAssembler(args[0], code)
	fmt.Println(a)
	// p := asm.NewProgram(args[0])
	// p.Read()
	// p.ResolveSymbols()
	// p.EmitToStdout()
}

var asmCmd = &cobra.Command{
	Use:    "asm <file.asm>",
	Short:  "Assembles a single Hack assembly (.asm) file into Hack machine code",
	Long:   `See https://github.com/trammell/n2t`,
	Args:   cobra.ExactArgs(1),
	Run:    asmFn,
	PreRun: setUpLogging,
}

func init() {
	rootCmd.AddCommand(asmCmd)
}
