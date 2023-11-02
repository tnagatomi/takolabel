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
	"fmt"
	"github.com/spf13/cobra"
	"github.com/tnagatomi/takolabel/takolabel"
)

// emptyCmd represents the empty command
var emptyCmd = &cobra.Command{
	Use:   "empty",
	Short: "Delete labels specified in takolabel_empty.yml",
	RunE: func(cmd *cobra.Command, args []string) error {
		t, err := takolabel.NewTakolabel(dryRun)
		if err != nil {
			return fmt.Errorf("failed initialization: %v", err)
		}
		c := takolabel.ConfigEmpty{}
		if err := c.Parse("takolabel_empty.yml"); err != nil {
			return fmt.Errorf("failed parsing create config: %v", err)
		}

		if !dryRun && !confirm() {
			fmt.Printf("Canceled execution\n")
			return nil
		}

		if err := t.Empty(c.Repos); err != nil {
			return fmt.Errorf("failed emptying labels: %v", err)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(emptyCmd)
}
