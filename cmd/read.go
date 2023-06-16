// Package cmd /*
package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/rowanjacobs/bsa-reader/bsareader"
	"github.com/spf13/cobra"
)

var Output string

// readCmd represents the read command
var readCmd = &cobra.Command{
	Use:   "read",
	Short: "Read a particular record from a BSA file",
	Long: `Read a particular record from a BSA file.
Optionally output to a file path using -o, otherwise will write to stdout.

Usage:
  bsa-reader read MAPS.BSA MAPPITEM.017
	bsa-reader read BLOCKS.BSA N0000019.RDB -o n019.rdb`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		bsaPath := args[0]
		name := args[1]
		// slow way. will probably break:
		bsa, err := ioutil.ReadFile(bsaPath)
		if err != nil {
			log.Fatal(err.Error())
		}
		record := bsareader.ReadRecord(bsa, name)

		if Output == "" {
			fmt.Println(string(record.Contents))
		} else {
			err := os.WriteFile(Output, record.Contents, 0644)
			if err != nil {
				log.Fatal(err.Error())
			}

			fmt.Printf("Read record %q and output %d bytes to %s\n", record.Name, record.Size, Output)
		}
	},
}

func init() {
	rootCmd.AddCommand(readCmd)
	readCmd.Flags().StringVarP(&Output, "output", "o", "", "File path to write output to (optional)")
}
