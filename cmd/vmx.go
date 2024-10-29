package cmd

import (
	"github.com/spf13/cobra"
	"github.com/trammell/n2t/pkg/vmx"
)

// The main VM translator ("VMX") function. All the magic happens in
// function vmx.Translate().
func vmxFn(cmd *cobra.Command, args []string) {
	vmx.Translate(args[0])
}

var vmxCmd = &cobra.Command{
	Use:    "vmx file1.vm ...",
	Short:  "Assembles a single Hack VM (.vm) file into Hack assembly code",
	Long:   `See https://github.com/trammell/n2t`,
	Args:   cobra.MinimumNArgs(1),
	Run:    vmxFn,
	PreRun: setUpLogging,
}

func init() {
	RootCmd.AddCommand(vmxCmd)
}
