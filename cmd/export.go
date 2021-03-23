/*
Copyright © 2021 SuperOrbital, LLC <info@superorbital.io>

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
	"io"
	"os"

	"github.com/spf13/cobra"
	"github.com/superorbital/wordchain/words"
)

// exportCmd represents the export command
var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Print the internal word list json file to the console",
	Long: `This command will print the internal work list json file to the console, which can then be used to create your own word list which can be passed in
	using the --json or -j options like this:

	$	wordchain random --json ./my-customer-list.json
	`,
	Run: func(cmd *cobra.Command, args []string) {
		internal, err := words.GetList("")
		if err != nil {
			fmt.Fprintln(os.Stderr, "[ERROR] Could not read internal word list")
			os.Exit(1)
		}
		_, err = io.Copy(os.Stdout, internal)
		if err != nil {
			fmt.Fprintln(os.Stderr, "[ERROR] Could not print internal word list")
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(exportCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// exportCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// exportCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
