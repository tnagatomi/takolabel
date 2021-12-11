// Copyright (c) 2021 Takayuki NAGATOMI
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package cmd

import (
	"context"
	"fmt"
	"github.com/tommy6073/takolabel/takolabel"
	"os"

	"github.com/spf13/cobra"
)

// emptyCmd represents the empty command
var emptyCmd = &cobra.Command{
	Use:   "empty",
	Short: "Delete labels specified in takolabel_delete.yml",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()
		client, err := takolabel.GitHubClient(ctx)
		if err != nil {
			return fmt.Errorf("couldn't get github client: %v", err)
		}

		e := takolabel.Empty{}
		if err = e.Gather(); err != nil {
			return fmt.Errorf("couldn't gather empty: %v", err)
		}
		if dryRun {
			if err = e.DryRun(ctx, client); err != nil {
				return fmt.Errorf("empty dry-run failed: %v", err)
			}
		} else {
			if takolabel.Confirm() {
				if err = e.Execute(ctx, client); err != nil {
					return fmt.Errorf("empty execution failed: %v", err)
				}
			} else {
				os.Exit(0)
			}
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(emptyCmd)
}
