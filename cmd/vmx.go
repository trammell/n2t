package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// The main VM translator ("VMX") function
func vmxFn(cmd *cobra.Command, args []string) {
	fmt.Println(`Ultimately call out to vmx.Translate() here`)
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
