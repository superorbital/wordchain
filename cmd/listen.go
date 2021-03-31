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
	"github.com/spf13/cobra"
	"github.com/superorbital/wordchain/listen"
	types "github.com/superorbital/wordchain/types"
)

var settings types.Listener

// listenCmd represents the listen command
var listenCmd = &cobra.Command{
	Use:   "listen",
	Short: "Listen on a port for requests to generate wordchains",
	Long:  `This launches the application as a long running service that will respond to a JSON request with a wordchain`,
	Run: func(cmd *cobra.Command, args []string) {
		listen.Listen(prefs, settings)
	},
}

func init() {
	rootCmd.AddCommand(listenCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listenCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listenCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	listenCmd.Flags().IntVar(&settings.Port, "port", 8080, "Port to bind to")
}
