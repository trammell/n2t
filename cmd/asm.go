package cmd

import (
	"github.com/spf13/cobra"
	"github.com/trammell/nand2tetris/pkg/asm"
)

func init() {
	rootCmd.AddCommand(asmCmd)
}

var asmCmd = &cobra.Command{
	Use:   "asm <file.asm>",
	Short: "Assembles a single file into Hack binary code",
	Long:  `See https://github.com/trammell/nand2tetris`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		p := asm.NewProgram(args[0])
		p.Read()
		p.Resolve()
		p.Emit()
	},
}
