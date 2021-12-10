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

package takolabel

import (
	"context"
	"fmt"
	"github.com/google/go-github/v41/github"
	"os"
)

type Delete struct {
	Target DeleteTarget
}

type DeleteTargetConfig struct {
	Repositories []string
	Labels       []string
}

type DeleteTarget struct {
	Repositories Repositories
	Labels       []string
}

func (d *Delete) Gather() error {
	content, err := os.ReadFile("takolabel_delete.yml")
	if err != nil {
		return fmt.Errorf("read file failed: %v", err)
	}

	target, err := ParseDelete(content)
	if err != nil {
		return fmt.Errorf("parse delete failed: %v", err)
	}
	d.Target = target

	return nil
}

func (d *Delete) DryRun() {
	for _, repository := range d.Target.Repositories {
		for _, label := range d.Target.Labels {
			fmt.Printf("would delete label \"%s\" for repository \"%s\"\n", label, repository.Owner+"/"+repository.Repo)
		}
	}
}

func (d *Delete) Execute(ctx context.Context, client *github.Client) {
	for _, repository := range d.Target.Repositories {
		for _, label := range d.Target.Labels {
			err := DeleteLabel(ctx, client.Issues, label, repository.Owner, repository.Repo)
			if err != nil {
				fmt.Printf("error deleting label \"%s\" for repository \"%s\": %s\n", label, repository.Owner+"/"+repository.Repo, err)
			} else {
				fmt.Printf("deleted label \"%s\" for repository \"%s\"\n", label, repository.Owner+"/"+repository.Repo)
			}
		}
	}
}
