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
	"gopkg.in/yaml.v3"
	"os"
	"strings"
)

type Empty struct {
	Target EmptyTarget
}

type EmptyTargetConfig struct {
	Repositories []string
}

type EmptyTarget struct {
	Repositories Repositories
}

func (e *Empty) Gather() error {
	content, err := os.ReadFile("takolabel_empty.yml")
	if err != nil {
		return fmt.Errorf("read file failed: %v", err)
	}

	if err := e.Parse(content); err != nil {
		return fmt.Errorf("parse empty failed: %v", err)
	}

	return nil
}

func (e *Empty) Parse(bytes []byte) error {
	targetConfig := EmptyTargetConfig{}
	if err := yaml.Unmarshal(bytes, &targetConfig); err != nil {
		return fmt.Errorf("yaml unmarshal failed: %v", err)
	}

	target := EmptyTarget{}
	for _, repository := range targetConfig.Repositories {
		s := strings.Split(repository, "/")
		if len(s) != 2 {
			return fmt.Errorf("repository %s is not properly formatted in setting yaml file", repository)
		}
		target.Repositories = append(target.Repositories, Repository{s[0], s[1]})
	}

	e.Target = target
	return nil
}

func (e *Empty) DryRun(ctx context.Context, client *github.Client) error {
	opt := &github.ListOptions{}
	for _, repository := range e.Target.Repositories {
		labels, err := ListLabels(ctx, client.Issues, repository.Owner, repository.Repo, opt)
		if err != nil {
			return err
		}
		for _, label := range labels {
			fmt.Printf("would delete label \"%s\" for repository \"%s\"\n", *label.Name, repository.Owner+"/"+repository.Repo)
		}
	}
	return nil
}

func (e *Empty) Execute(ctx context.Context, client *github.Client) error {
	opt := &github.ListOptions{}
	for _, repository := range e.Target.Repositories {
		labels, err := ListLabels(ctx, client.Issues, repository.Owner, repository.Repo, opt)
		if err != nil {
			return err
		}
		for _, label := range labels {
			err := DeleteLabel(ctx, client.Issues, *label.Name, repository.Owner, repository.Repo)
			if err != nil {
				fmt.Printf("error deleting label \"%s\" for repository \"%s\": %s\n", label, repository.Owner+"/"+repository.Repo, err)
			} else {
				fmt.Printf("deleted label \"%s\" for repository \"%s\"\n", *label.Name, repository.Owner+"/"+repository.Repo)
			}
		}
	}
	return nil
}
