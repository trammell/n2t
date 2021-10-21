package cmd

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
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

func setLogLevel(cmd *cobra.Command, args []string) {
	output := zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: `//`}
	output.FormatLevel = func(i interface{}) string {
		return ``
	}
	log.Logger = log.Output(output)
	zerolog.TimeFieldFormat = ``
	if Verbose {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.Disabled)
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
}
