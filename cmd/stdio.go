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
	"io/ioutil"
	"os"
	"syscall"

	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
)

// stdioCmd represents the stdio command
var stdioCmd = &cobra.Command{
	Use:   "stdio",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("stdio called")
		// no pipe: fd == 0 (stdin)
		if !terminal.IsTerminal(int(syscall.Stdin)) {
			data, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				cmd.Println(err.Error())
				return err
			}

			fmt.Printf("Info: stdin data: %v\n", string(data))
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(stdioCmd)
}
