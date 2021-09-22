package cmd

import (
	"fmt"
	"github.com/rowanjacobs/bsa-reader/bsareader"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"

	"github.com/spf13/cobra"
)

// flatCmd represents the flats command
var flatCmd = &cobra.Command{
	Use:   "flat",
	Short: "Read flats files",
	Long: `Usage:
bsa-reader flat $index`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		index := args[0]
		i, err := strconv.Atoi(index)
		if err != nil {
			fmt.Printf("Error: %q is not an integer", i)
			os.Exit(1)
		}

		bsaPath := filepath.Join(bsareader.GetDaggerfallPath(), "FLATS.CFG") // brittle, shitty
		// slow way. will probably break:
		bsa, err := ioutil.ReadFile(bsaPath)
		if err != nil {
			panic(err)
		}

		flats := bsareader.ReadFlats(bsa)
		fmt.Printf("Read %d flats\n from FLATS.CFG", len(flats))
		fmt.Printf("Flat %d: %s\n", i, flats[i])
	},
}

func init() {
	rootCmd.AddCommand(flatCmd)
}
