package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/trammell/n2t/pkg/asm"
)

func asmFn(cmd *cobra.Command, args []string) {
	a := asm.NewAssembler(args[0])
	a.Assemble()
	for _, mc := range a.MachineCode {
		fmt.Printf("%016b\n", mc)
	}
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
	RootCmd.AddCommand(asmCmd)
}
