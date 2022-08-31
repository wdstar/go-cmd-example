/*
Copyright © 2019-2022 wdstar

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
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/kballard/go-shellquote"
	"github.com/spf13/cobra"
	//"gopkg.in/alessio/shellescape.v1"
)

// argsCmd represents the args command
var argsCmd = &cobra.Command{
	Use:   "args",
	Short: "Arguments processing examples",
	Long: `You can pass arbitrary arguments.
This subcommand print those command line.`,
	// Cobra counts all arguments except the first double dash `--`.
	// e.g. command line: go-cmd-example args a b -- c d (4 arguments)
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("args called")

		// e.g. command line: go-cmd-example args a b -- c d
		// args does not contain the first double dash `--`.
		// args => [a b c d]
		if i := cmd.ArgsLenAtDash(); i != -1 {
			// => i: 2
			fmt.Printf("Command has `--` argument. Command.ArgsLenAtDash: %d\n", i)
			mainArgs := args[:i]  // => [a b]
			extraArgs := args[i:] // => [c d]
			fmt.Printf("Full args: %v\n", args)
			fmt.Printf("Main args: %v\n", mainArgs)
			fmt.Printf("Extra args: %v\n", extraArgs)
		} else {
			fmt.Println("Command has no `--` argument.")
		}

		// e.g. command line: go-cmd-example args a "b c" '"d","e"' "'f', 'g'"
		fmt.Printf("Simple joined args: %s\n", strings.Join(os.Args, " "))
		// => go-cmd-example args a b c "d","e" 'f', 'g' // NG
		fmt.Printf("Quoted args: %s\n", shellquote.Join(os.Args...))
		// => go-cmd-example args a 'b c' \"d\",\"e\" \''f'\'', '\''g'\'  // OK
		var quotedArgs []string
		quoted := regexp.MustCompile(`["\s]`)
		for _, arg := range os.Args {
			//if strings.Contains(arg, " ") {
			if quoted.MatchString(arg) {
				arg = strconv.Quote(arg)
				// => go-cmd-example args a "b c" "\"d\",\"e\"" "'f', 'g'"  // OK
				//arg = shellescape.Quote(arg)
				// => go-cmd-example args a 'b c' '"d","e"' ''"'"'f'"'"', '"'"'g'"'"''  // OK but !!
			}
			quotedArgs = append(quotedArgs, arg)
		}
		fmt.Printf("Quoted args: %s\n", strings.Join(quotedArgs, " "))
	},
}

func init() {
	rootCmd.AddCommand(argsCmd)
}
