package cmd

import (
	"github.com/spf13/cobra"
)

// The main VM translator ("VMX") function
func vmxFn(cmd *cobra.Command, args []string) {
	//a := vmx.NewTranslator(args[0])
	//a.Assemble()
	//for _, mc := range a.MachineCode {
	//	fmt.Printf("%016b\n", mc)
	//}
}

var vmxCmd = &cobra.Command{
	Use:    "vmx <file.asm>",
	Short:  "Assembles a single Hack assembly (.asm) file into Hack machine code",
	Long:   `See https://github.com/trammell/n2t`,
	Args:   cobra.ExactArgs(1),
	Run:    asmFn,
	PreRun: setUpLogging,
}

func init() {
	rootCmd.AddCommand(vmxCmd)
}
