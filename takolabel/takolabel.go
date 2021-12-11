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
	"golang.org/x/oauth2"
	"os"
)

type Repository struct {
	Owner string
	Repo  string
}

type Repositories []Repository

type Label struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Color       string `yaml:"color"`
}

type Labels []Label

func GetGitHubClient(ctx context.Context) (*github.Client, error) {
	githubToken := os.Getenv("TAKOLABEL_TOKEN")
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: githubToken},
	)
	tc := oauth2.NewClient(ctx, ts)

	var client *github.Client
	enterpriseURL := os.Getenv("TAKOLABEL_HOST")
	if enterpriseURL == "" {
		client = github.NewClient(tc)
	} else {
		var err error
		client, err = github.NewEnterpriseClient(enterpriseURL, enterpriseURL, tc)
		if err != nil {
			return nil, fmt.Errorf("error setting ghe client: %s\n", err)
		}
	}
	return client, nil
}

func CreateLabel(ctx context.Context, issuesService *github.IssuesService, label Label, owner string, repo string) (*github.Label, error) {
	githubLabel := &github.Label{
		Name:        github.String(label.Name),
		Description: github.String(label.Description),
		Color:       github.String(label.Color),
	}
	createdLabel, _, err := issuesService.CreateLabel(ctx, owner, repo, githubLabel)
	return createdLabel, err
}

func DeleteLabel(ctx context.Context, issuesService *github.IssuesService, label string, owner string, repo string) error {
	_, err := issuesService.DeleteLabel(ctx, owner, repo, label)
	return err
}

func ListLabels(ctx context.Context, issuesService *github.IssuesService, owner string, repo string, opt *github.ListOptions) ([]*github.Label, error) {
	labels, _, err := issuesService.ListLabels(ctx, owner, repo, opt)
	if err != nil {
		return nil, fmt.Errorf("list labels failed: %v", err)
	}
	return labels, nil
}
