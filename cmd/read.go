/*
Copyright © 2020 Rowan Jacobs <rojacobs@vmware.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"io/ioutil"

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
			panic(err)
		}
		record := bsareader.ReadRecord(bsa, name)

		if Output == "" {
			fmt.Println(string(record.Contents))
		} else {
			ioutil.WriteFile(Output, record.Contents, 0644)

			fmt.Printf("Read record %q and output %d bytes to %s\n", record.Name, record.Size, Output)
		}
	},
}

func init() {
	rootCmd.AddCommand(readCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// readCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	readCmd.Flags().StringVarP(&Output, "output", "o", "", "File path to write output to (optional)")
}
