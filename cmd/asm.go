package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/trammell/nand2tetris/pkg/asm"
)

func init() {
	rootCmd.AddCommand(asmCmd)
}

var asmCmd = &cobra.Command{
	Use:   "asm",
	Short: "Assemble a file",
	Long:  `See https://github.com/trammell/nand2tetris`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		p := asm.NewProgram(os.Args[1])
		p.Read()
		p.Resolve()
		p.Emit()
	},
}
