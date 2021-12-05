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
	"github.com/spf13/cobra"
	"github.com/tommy6073/takolabel/takolabel"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create labels specified in takolabel_create.yml",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()
		client, err := takolabel.GetGitHubClient(ctx)
		if err != nil {
			return fmt.Errorf("couldn't get github client: %v", err)
		}

		c := takolabel.Create{}
		if err = c.Gather(); err != nil {
			return fmt.Errorf("couldn't gather create: %v", err)
		}
		if dryRun {
			c.DryRun()
		} else {
			c.Execute(ctx, client)
		}
		return err
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}
