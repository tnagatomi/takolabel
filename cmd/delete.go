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
	"github.com/spf13/cobra"
	"github.com/tommy6073/takolabel/takolabel"
	"os"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete labels specified in takolabel_delete.yml",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()
		client := takolabel.GetGitHubClient(ctx)

		delete_ := takolabel.Delete{}
		if err := delete_.Gather(); err != nil {
			return err
		}
		if dryRun {
			delete_.DryRun()
		} else {
			if takolabel.Confirm() {
				delete_.Execute(ctx, client)
			} else {
				os.Exit(0)
			}
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
