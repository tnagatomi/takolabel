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
	"github.com/google/go-github/v33/github"
	"gopkg.in/yaml.v3"
	"os"
	"strings"
)

type Create struct {
	Target CreateTarget
}

type CreateTargetConfig struct {
	Repositories []string
	Labels       []Label
}

type CreateTarget struct {
	Repositories []Repository
	Labels       []Label
}

func (c *Create) Parse(bytes []byte) error {
	targetConfig := CreateTargetConfig{}
	err := yaml.Unmarshal(bytes, &targetConfig)
	if err != nil {
		return err
	}

	target := CreateTarget{Labels: targetConfig.Labels}
	for _, repository := range targetConfig.Repositories {
		s := strings.Split(repository, "/")
		if len(s) != 2 {
			return fmt.Errorf("repository %s is not properly formatted in setting yaml file", repository)
		}
		target.Repositories = append(target.Repositories, Repository{s[0], s[1]})
	}

	c.Target = target

	return nil
}

func (c *Create) Gather() error {
	content, err := os.ReadFile("takolabel_create.yaml")
	if err != nil {
		return err
	}

	err = c.Parse(content)
	if err != nil {
		return err
	}

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
