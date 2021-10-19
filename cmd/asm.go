package cmd

import (
	"github.com/spf13/cobra"
	"github.com/trammell/nand2tetris/pkg/asm"
)

func asmFn(cmd *cobra.Command, args []string) {
	p := asm.NewProgram(args[0])
	p.Read()
	p.ResolveSymbols()
	p.EmitToStdout()
}

var asmCmd = &cobra.Command{
	Use:   "asm <file.asm>",
	Short: "Assembles a single Hack assembly (.asm) file into Hack machine code",
	Long:  `See https://github.com/trammell/nand2tetris`,
	Args:  cobra.ExactArgs(1),
	Run:   asmFn,
	PreRun: setLogLevel,
}

func init() {
	rootCmd.AddCommand(asmCmd)
}
