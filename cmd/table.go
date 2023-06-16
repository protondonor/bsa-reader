package cmd

import (
	"fmt"
	"github.com/rowanjacobs/bsa-reader/bsareader/maps"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// tableCmd represents the table command
var tableCmd = &cobra.Command{
	Use:   "table",
	Short: "Read a table file",
	Long: `Read a particular record from a table file.
These files can be produced by running "bsa-reader read"
on MAPS.BSA with the argument MAPTABLE.0XY, where XY is
the region number.

See https://en.uesp.net/wiki/Daggerfall_Mod:Region_Numbers
for a list of region numbers.

Usage:
  bsa-reader tables MAPTABLE.017`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		bsaPath := args[0]
		// slow way. will probably break:
		bsa, err := os.ReadFile(bsaPath)
		if err != nil {
			log.Fatal(err)
		}
		table := maps.ReadTable(bsa)
		for i := 0; i < len(table.Rows); i++ {
			fmt.Println(table.Rows[i])
		}
	},
}

func init() {
	rootCmd.AddCommand(tableCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// tableCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// tableCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
