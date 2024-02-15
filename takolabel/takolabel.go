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
	"strings"
)

// Takolabel composites github.Client and has dry-run option
type Takolabel struct {
	client *github.Client
	dryRun bool
}

// NewTakolabel returns new Takolabel struct
func NewTakolabel(dryRun bool) (*Takolabel, error) {
	githubToken := os.Getenv("TAKOLABEL_TOKEN")
	if githubToken == "" {
		return nil, fmt.Errorf("TAKOLABEL_TOKEN environment variable is not set")
	}
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: githubToken},
	)
	tc := oauth2.NewClient(context.Background(), ts)

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

	return &Takolabel{
		client: client,
		dryRun: dryRun,
	}, nil
}

// Create GitHub labels for specified repositories with specified labels
func (takolabel *Takolabel) Create(labels []Label, repos []Repo) error {
	for _, r := range repos {
		for _, l := range labels {
			if takolabel.dryRun {
				fmt.Printf("Would create label %s for repository %q\n", l.Name, r.Owner+"/"+r.Repo)
				continue
			}

			label := &github.Label{
				Name:        github.String(l.Name),
				Description: github.String(l.Description),
				Color:       github.String(l.Color),
			}
			_, _, err := takolabel.client.Issues.CreateLabel(context.Background(), r.Owner, r.Repo, label)

			if err != nil {
				if strings.Contains(err.Error(), "already_exists") {
					fmt.Printf("Label %q already exists for repository %q\n", l.Name, r.Owner+"/"+r.Repo)
					continue
				}
				fmt.Printf("error creating label %s for repository %q: %v\n", l.Name, r.Owner+"/"+r.Repo, err)
				continue
			}
			fmt.Printf("Created label %s for repository %q\n", l.Name, r.Owner+"/"+r.Repo)
		}
	}

	return nil
}

// Delete GitHub labels for specified repositories with specified labels
func (takolabel *Takolabel) Delete(labels []string, repos []Repo) error {
	for _, r := range repos {
		for _, l := range labels {
			if takolabel.dryRun {
				fmt.Printf("Would delete label %s for repository %q\n", l, r.Owner+"/"+r.Repo)
				continue
			}
			if _, err := takolabel.client.Issues.DeleteLabel(context.Background(), r.Owner, r.Repo, l); err != nil {
				if strings.Contains(err.Error(), "404 Not Found") {
					fmt.Printf("Label %q doesn't exist for repository %q\n", l, r.Owner+"/"+r.Repo)
					continue
				}
				return fmt.Errorf("error deleting label %s for repository %q: %v", l, r.Owner+"/"+r.Repo, err)
			}
			fmt.Printf("Deleted label %s for t %s/%s\n", l, r.Owner, r.Repo)
		}
	}

	return nil
}

// Sync GitHub labels for specified repositories with specified labels
func (takolabel *Takolabel) Sync(labels []Label, repos []Repo) error {
	for _, r := range repos {
		if takolabel.dryRun {
			for _, l := range labels {
				fmt.Printf("Would set label %s for repository %q\n", l.Name, r.Owner+"/"+r.Repo)
			}
			continue
		}

		fmt.Printf("Deleting all labels for %q first\n", r.Repo)

		err := takolabel.deleteAllLabels(r)

		if err != nil {
			fmt.Printf("error deleting all labels for repository %q\n", r.Repo)
			continue
		}

		fmt.Printf("Deleted all labels for %q\n", r.Repo)

		for _, l := range labels {
			label := &github.Label{
				Name:        github.String(l.Name),
				Description: github.String(l.Description),
				Color:       github.String(l.Color),
			}
			_, _, err := takolabel.client.Issues.CreateLabel(context.Background(), r.Owner, r.Repo, label)

			if err != nil {
				fmt.Printf("error creating label %s for repository %q: %v\n", l.Name, r.Owner+"/"+r.Repo, err)
				continue
			}
			fmt.Printf("Created label %s for repository %q\n", l.Name, r.Owner+"/"+r.Repo)
		}

		fmt.Printf("Sync completed for repository %q\n", r.Repo)
	}

	return nil
}

// Empty GitHub labels for specified repositories
func (takolabel *Takolabel) Empty(repos []Repo) error {
	for _, r := range repos {
		labels, _, err := takolabel.client.Issues.ListLabels(context.Background(), r.Owner, r.Repo, nil)
		if err != nil {
			return fmt.Errorf("list labels failed: %v", err)
		}

		if takolabel.dryRun {
			for _, l := range labels {
				fmt.Printf("Would delete label %s for repository %q\n", *l.Name, r.Owner+"/"+r.Repo)
			}
			continue
		}

		for _, l := range labels {
			_, err := takolabel.client.Issues.DeleteLabel(context.Background(), r.Owner, r.Repo, *l.Name)
			if err != nil {
				return fmt.Errorf("error deleting label %s for repository %q: %v\n", l, r.Owner+"/"+r.Repo, err)
			} else {
				fmt.Printf("Deleted label %s for repository %q\n", *l.Name, r.Owner+"/"+r.Repo)
			}
		}
	}
	return nil
}

func (takolabel *Takolabel) deleteAllLabels(r Repo) error {
	existingLabels, _, err := takolabel.client.Issues.ListLabels(context.Background(), r.Owner, r.Repo, nil)

	if err != nil {
		return fmt.Errorf("error getting labels for repository %q: %v", r.Repo, err)
	}

	for _, l := range existingLabels {
		_, err := takolabel.client.Issues.DeleteLabel(context.Background(), r.Owner, r.Repo, *l.Name)
		if err != nil {
			fmt.Printf("error deleting l %v for repository: %q\n", *l.Name, r.Repo)
		}
	}

	return nil
}
