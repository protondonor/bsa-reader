package cmd

import (
	"fmt"
	"github.com/rowanjacobs/bsa-reader/bsareader/maps"
	"io/ioutil"

	"github.com/spf13/cobra"
)

// namesCmd represents the names command
var namesCmd = &cobra.Command{
	Use:   "names",
	Short: "Read a .names file",
	Long: `Read a particular record from a .names file.
These files can be produced by running "bsa-reader read"
on MAPS.BSA with the argument MAPNAMES.0XY, where XY is
the region number.

See https://en.uesp.net/wiki/Daggerfall_Mod:Region_Numbers
for a list of region numbers.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		bsaPath := args[0]
		// slow way. will probably break:
		bsa, err := ioutil.ReadFile(bsaPath)
		if err != nil {
			panic(err)
		}
		names := maps.ReadNames(bsa)
		for i := 0; i < len(names.Names); i++ {
			fmt.Println(names.Names[i])
		}
	},
}

func init() {
	rootCmd.AddCommand(namesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// namesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// namesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
