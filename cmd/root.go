// Package cmd /*
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "bsa-reader",
	Short: "Reads Daggerfall BSA files",
	Long: `bsa-reader is a CLI tool that reads Daggerfall BSA files.
It can be used to find the number of records in a BSA file
and their size and format, and to read any individual record.

Functions to read particular record formats (such as MapPItem)
are forthcoming.
`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
}
