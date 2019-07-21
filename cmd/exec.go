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
	"io"
	"os"
	"os/exec"
	"strings"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"

	"github.com/spf13/cobra"
)

// execCmd represents the exec command
var execCmd = &cobra.Command{
	Use:                   "exec command_path [-- command_flags...] [command_args...]",
	DisableFlagsInUseLine: true,
	Short:                 "command wrapper",
	Long: `exec subcommand executes an external command.
This has pipeline feature too.`,
	SilenceErrors: true,
	SilenceUsage:  true,
	Args:          cobra.MinimumNArgs(1),
	//DisableFlagParsing: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Info: exec called")
		path := args[0]
		args = args[1:]
		fmt.Printf("Info: exec command: %s %s\n", path, strings.Join(args[:], " "))
		targetCmd := exec.Command(path, args...)
		errChan := make(chan error, 1)

		// no pipe: fd == 0 (stdin)
		if !terminal.IsTerminal(int(syscall.Stdin)) {
			stdin, err := targetCmd.StdinPipe()
			if err != nil {
				// cobra.Command default output is `stderr`.
				cmd.Println("Error: failed to open StdinPipe.")
				return err
			}
			go func() {
				defer stdin.Close()
				defer close(errChan)
				// Note: we must use goroutine,
				// because when writing data exceeding pipe capacity this line is blocked until reading it.
				_, err = io.Copy(stdin, os.Stdin)
				errChan <- err
				if err != nil {
					fmt.Printf("failed to copy piped command stdin from os.Stdin: %v\n", err)
				}
			}()
		} else {
			// not used.
			close(errChan)
		}

		targetCmd.Stdout = os.Stdout
		targetCmd.Stderr = os.Stderr
		err := targetCmd.Run()

		// Note: closed channel returns buffered message or a zero value if it is empty.
		stdinErr := <-errChan
		if stdinErr != nil {
			return stdinErr
		}
		if err != nil {
			cmd.Println("Error: failed to execute command")
			return err
		}

		state := targetCmd.ProcessState
		fmt.Printf("Info: System Time: %v, User Time: %v\n", state.SystemTime(), state.UserTime())

		return nil
	},
}

func init() {
	rootCmd.AddCommand(execCmd)
}
