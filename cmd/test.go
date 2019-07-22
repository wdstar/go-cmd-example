/*
Copyright © 2019 wdstar

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
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	PreRun: func(cmd *cobra.Command, args []string) {
		// Preprocess string slice (from CSV)
		// e.g. --list='[ a, b, c ]'
		fmt.Println(strings.Join(viper.GetStringSlice("list"), "|"))
		// e.g. => '[ a| b| c ]'
		list := []string{}
		for _, elm := range viper.GetStringSlice("list") {
			// clean up each element
			list = append(list, strings.Trim(elm, " []"))
		}
		// override string slice
		viper.Set("list", list)
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Info: test called")

		// Normalized string slice
		fmt.Println(strings.Join(viper.GetStringSlice("list"), "|"))
		// e.g. => 'a|b|c'

		// Flag alias resolution
		ws := viper.GetString("workspace")
		if ws == "" {
			ws = viper.GetString("ws")
		}
		fmt.Printf("Info: --workspace: %s\n", ws)
	},
}

func init() {
	rootCmd.AddCommand(testCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// testCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// testCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	testCmd.Flags().StringSliceP("list", "l", []string{}, "CSV String")
	testCmd.Flags().StringP("workspace", "w", "", "Workspace ID.")
	testCmd.Flags().String("ws", "", "Workspace ID, alias of --workspace")
	// cmd.Flag() == Flags().Lookup()
	viper.BindPFlag("list", testCmd.Flag("list"))
	viper.BindPFlag("workspace", testCmd.Flags().Lookup("workspace"))
	viper.BindPFlag("ws", testCmd.Flags().Lookup("ws"))
	// This line is alias key setting for Viper's accessors and conf. persistence.
	//viper.RegisterAlias("ws", "workspace")
}
