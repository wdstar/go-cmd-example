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
	"gopkg.in/yaml.v2"
)

// yamlCmd represents the yaml command
var yamlCmd = &cobra.Command{
	Use:   "yaml",
	Short: "Test command for YAML",
	Long:  `This subcommand tests YAML processing.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Info: yaml called")

		// e.g. input data:
		//   - YAML flow style:  '[a, b, "c"]'
		//   - YAML block style: "$(echo -e '- a\n- b\n- c\n')"
		//   - JSON:             '["a", "b", "c"]'
		var list []string
		err := yaml.Unmarshal([]byte(viper.GetString("list")), &list)
		if err != nil {
			return err
		}
		fmt.Println(strings.Join(list, "|"))
		// => "a|b|c"

		return nil
	},
}

func init() {
	rootCmd.AddCommand(yamlCmd)

	yamlCmd.Flags().StringP("list", "l", "", "YAML array data")
	viper.BindPFlags(yamlCmd.Flags())
}
