// Package cmd /*
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/rowanjacobs/bsa-reader/bsareader"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List records in a BSA file",
	Long: `List records in a BSA file.

Usage:
  bsa-reader list MAPS.BSA
`,
	Run: func(cmd *cobra.Command, args []string) {
		// slow way. will probably break:
		bsa, err := os.ReadFile(args[0])
		if err != nil {
			log.Fatal(err.Error())
		}
		records := bsareader.List(bsa)

		for i := 0; i < len(records); i++ {
			fmt.Printf("%s, %d\n", records[i].Name, records[i].Size)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
