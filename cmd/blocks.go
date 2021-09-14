package cmd

import (
	"fmt"
	"github.com/rowanjacobs/bsa-reader/bsareader"
	"github.com/spf13/cobra"
	"io/ioutil"
	"path/filepath"
)

var blocksCmd = &cobra.Command{
	Use:   "blocks",
	Short: "List blocks given a dungeon name and region number",
	Long: `
Usage:
  bsa-reader blocks "Ruins of Cosh Hall" 061`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		region := args[1]

		bsaPath := filepath.Join(bsareader.GetDaggerfallPath(), "MAPS.BSA") // brittle, shitty
		// slow way. will probably break:
		bsa, err := ioutil.ReadFile(bsaPath)
		if err != nil {
			panic(err)
		}

		blocks := bsareader.ReadBlocks(bsa, name, region)
		for i := 0; i < len(blocks); i++ {
			fmt.Println(blocks[i])
		}
	},
}

func init() {
	rootCmd.AddCommand(blocksCmd)
}
