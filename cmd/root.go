package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var Verbose bool = false

var rootCmd = &cobra.Command{
	Use:   "n2t",
	Short: "n2t is a helper application for Nand2Tetris",
	Long:  `See https://github.com/trammell/nand2tetris`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}


func Execute() {
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
