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

	types "github.com/superorbital/wordchain/types"
)

//var cfgFile string
var prefs types.Preferences

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Version: "v1.0.0",
	Use:     "wordchain",
	Short:   "Generates a word chain",
	Long: `wordchain is an application that can generate a readable chain
	of customizable words for naming things like
	containers, clusters, and other objects.
	( e.g. silly-code )`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) {},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	// We should likely init this after we have read the config flag.
	//	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	//	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $HOME/.wordchain.yaml)")
	rootCmd.PersistentFlags().StringVarP(&prefs.Divider, "divider", "d", "-", "The divider to use between words")
	rootCmd.PersistentFlags().StringVarP(&prefs.WordFile, "json", "j", "", "word list json file")
	rootCmd.PersistentFlags().Int8VarP(&prefs.Length, "length", "l", 5, "The length of words to use")
	rootCmd.PersistentFlags().StringVarP(&prefs.Postpend, "postpend", "o", "", "string to postpend to the output")
	rootCmd.PersistentFlags().StringVarP(&prefs.Prepend, "prepend", "r", "", "string to prepend to the output")
	rootCmd.PersistentFlags().StringSliceVarP(&prefs.Type, "type", "t", []string{"adjective", "noun"}, "Comma seperated list of word types")
	rootCmd.PersistentFlags().StringVarP(&prefs.Seed, "seed", "s", "", "A string seed to use when you want deterministic responses")
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
//func initConfig() {
//	if cfgFile != "" {
//		// Use config file from the flag.
//		viper.SetConfigFile(cfgFile)
//	} else {
//		// Find home directory.
//		home, err := homedir.Dir()
//		cobra.CheckErr(err)
//
//		// Search config in home directory with name ".wordchain" (without extension).
//		viper.AddConfigPath(home)
//		viper.SetConfigName(".wordchain")
//	}
//
//	viper.AutomaticEnv() // read in environment variables that match
//
//	// If a config file is found, read it in.
//	if err := viper.ReadInConfig(); err == nil {
//		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
//	}
//}
