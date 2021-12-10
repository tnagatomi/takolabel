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

type Create struct {
	Target CreateTarget
}

type CreateTargetConfig struct {
	Repositories []string
	Labels       Labels
}

type CreateTarget struct {
	Repositories Repositories
	Labels       Labels
}

func (c *Create) Gather() error {
	content, err := os.ReadFile("takolabel_create.yml")
	if err != nil {
		return fmt.Errorf("read file failed: %v", err)
	}

	target, err := ParseCreate(content)
	if err != nil {
		return fmt.Errorf("parse create failed: %v", err)
	}
	c.Target = target

	return nil
}

func (c *Create) DryRun() {
	for _, repository := range c.Target.Repositories {
		for _, label := range c.Target.Labels {
			fmt.Printf("Would create label \"%s\" for repository \"%s\"\n", label.Name, repository.Owner+"/"+repository.Repo)
		}
	}
}

func (c *Create) Execute(ctx context.Context, client *github.Client) {
	for _, repository := range c.Target.Repositories {
		for _, label := range c.Target.Labels {
			_, err := CreateLabel(ctx, client.Issues, label, repository.Owner, repository.Repo)
			if err != nil {
				fmt.Printf("error creating label \"%s\" for repository \"%s\": %s\n", label.Name, repository.Owner+"/"+repository.Repo, err)
			} else {
				fmt.Printf("created label \"%s\" for repository \"%s\"\n", label.Name, repository.Owner+"/"+repository.Repo)
			}
		}
	}
}
