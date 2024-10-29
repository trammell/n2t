package cmd

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var Verbose bool = false

var RootCmd = &cobra.Command{
	Use:   "n2t",
	Short: "A helper application for Nand2Tetris",
	Long:  `n2t is a command-line helper application for Nand2Tetris
providing an assembler and a compiler. See
https://github.com/trammell/nand2tetris for details.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func setUpLogging(cmd *cobra.Command, args []string) {
	output := zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: `//`}
	output.FormatLevel = func(i interface{}) string {
		return ``
	}
	log.Logger = log.Output(output)
	zerolog.SetGlobalLevel(zerolog.Disabled)
	if Verbose {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}
}

func Execute() error {
	if err := RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	RootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
}
