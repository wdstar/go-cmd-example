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
	"encoding/csv"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

// serdeCmd represents the serde command
var serdeCmd = &cobra.Command{
	Use:   "serde",
	Short: "Test command for SerDe",
	Long:  `This subcommand tests SerDe processing.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Info: serde called")

		var list []string

		// e.g. input data:
		//   - CSV: 'a,b,c'
		csvStr := viper.GetString("csv")
		if csvStr != "" {
			stringReader := strings.NewReader(csvStr)
			csvReader := csv.NewReader(stringReader)
			list, err := csvReader.Read()
			if err != nil {
				return err
			}
			fmt.Println("from CSV: " + strings.Join(list, "|"))
			// => "a|b|c"
		}

		// e.g. input data:
		//   - YAML flow style:  '[a, b, "c"]'
		//   - YAML block style: "$(echo -e '- a\n- b\n- c\n')"
		//   - JSON: '["a", "b", "c"]'
		yamlStr := viper.GetString("yaml")
		if yamlStr != "" {
			err := yaml.Unmarshal([]byte(yamlStr), &list)
			if err != nil {
				return err
			}
			fmt.Println("from YAML: " + strings.Join(list, "|"))
			// => "a|b|c"
		}

		// e.g. input data:
		//   - JSON: '["a", "b", "c"]'
		jsonStr := viper.GetString("json")
		if jsonStr != "" {
			err := json.Unmarshal([]byte(jsonStr), &list)
			if err != nil {
				return err
			}
			fmt.Println("from JSON: " + strings.Join(list, "|"))
			// => "a|b|c"
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(serdeCmd)

	serdeCmd.Flags().StringP("csv", "c", "", "CSV data")
	serdeCmd.Flags().StringP("json", "j", "", "JSON array data")
	serdeCmd.Flags().StringP("yaml", "y", "", "YAML array data")
	viper.BindPFlags(serdeCmd.Flags())
}
