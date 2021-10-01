package cmd

import (
	"asm"
	"os"

	"github.com/spf13/cobra"
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
		p.read()
		p.resolve()
		p.emit()
	},
}
